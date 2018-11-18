package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	//"reflect"
)

func SetUpDB() *APIMongoDB {
	db := APIMongoDB{
		"mongodb://127.0.0.1:27017",
		"libraryApp",
		"books",
		"users",
	}
	return &db
}

func APIBookURL(isbn int) string {
	return "http://openlibrary.org/api/books?bibkeys=ISBN:" + strconv.Itoa(isbn) + "&format=json&jscmd=data"
}

func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	var isbn ISBN
	err := json.NewDecoder(r.Body).Decode(&isbn)
	if err != nil || isbn.Isbn == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		jsonResp, err := RetrieveBookInfo(isbn.Isbn)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			book, check := ParseBookInfo(jsonResp, isbn.Isbn)
			if !check {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(book)
			}
		}

	}
}

func RetrieveBookInfo(isbn int) (jsonResp map[string]interface{}, err error) {
	resp, err := http.Get(APIBookURL(isbn))
	if err != nil {
		return
	} else {
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		return
	}
}

func ParseBookInfo(jsonResp map[string]interface{}, isbn int) (book Book, allIsWell bool) {
	allIsWell = true
	bookInfo, ok := jsonResp["ISBN:"+strconv.Itoa(isbn)].(map[string]interface{})
	if ok {

		var title, subtitle, publishDate, author, subject bool
		book.ISBN = isbn
		book.Title, title = bookInfo["title"].(string)
		book.Subtitle, subtitle = bookInfo["subtitle"].(string)
		book.Pages = int(bookInfo["number_of_pages"].(float64))
		book.PublishDate, publishDate = bookInfo["publish_date"].(string)
		authors, author := bookInfo["authors"].([]interface{})
		for i := 0; i < len(authors); i++ {
			tempAuthor := authors[i].(map[string]interface{})
			book.Authors = append(book.Authors, tempAuthor["name"].(string))
		}
		subjects, subject := bookInfo["subjects"].([]interface{})
		for i := 0; i < len(subjects); i++ {
			tempSubject := subjects[i].(map[string]interface{})
			book.Subjects = append(book.Subjects, tempSubject["name"].(string))
		}

		if !title || !subtitle || !publishDate || !author || !subject {
			allIsWell = false
			fmt.Printf("this should not happen, things have broken")
		}

	} else {
		allIsWell = false
		return
	}
	return
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	var registerStruct RegisterStruct
	err := json.NewDecoder(r.Body).Decode(&registerStruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		db := SetUpDB()
		_, found := db.GetUserByUsername(registerStruct.Username)
		if found {
			http.Error(w, "Username taken", http.StatusUnauthorized)
		} else {
			var user User
			user.Username = registerStruct.Username
			user.Email = registerStruct.Email
			user.Password = registerStruct.Password
			user.Status.Online = true
			user.Status.TimeStamp = time.Now().Unix()
			err := db.AddUser(user)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			} else {
				TokenReply(w, user)
			}
		}
	}
}

func TokenReply(w http.ResponseWriter, user User) {
	w.Header().Add("Token", user.Username+"isauthorized")
}

func LoginUser(w http.ResponseWriter, r *http.Request) {
	var loginStruct LoginStruct
	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		db := SetUpDB()
		user, found := db.GetUserByUsername(loginStruct.Username)
		if !found {
			FailedLoginAttempt(w)
		} else {
			if user.Password != loginStruct.Password {
				FailedLoginAttempt(w)
			} else {
				OnlineStatusChage(user, db, true)
				TokenReply(w, user)
			}
		}
	}
}

func LogoutUser(w http.ResponseWriter, r *http.Request) {
	var loginStruct LoginStruct
	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	token, header := r.Header["Token"]

	if !header || !ValidateToken(token, loginStruct.Username) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		} else {
			db := SetUpDB()
			user, found := db.GetUserByUsername(loginStruct.Username)
			if !found {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			} else {
				OnlineStatusChage(user, db, false)
			}
		}
	}
}

func OnlineStatusChage(user User, db *APIMongoDB, status bool) error {
	user.Status.Online = status
	user.Status.TimeStamp = time.Now().Unix()
	err := db.UpdateUserStatus(user)
	return err
}

func FailedLoginAttempt(w http.ResponseWriter) {
	http.Error(w, "Username and password does not match database entry", http.StatusUnauthorized)
}

func ValidateToken(token []string, username string) bool {
	return token[0] == username+"isauthorized"
}
