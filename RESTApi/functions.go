package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
	//"reflect"
)

//SetUpDB takes returns pointer to APIMongoDB struct
func SetUpDB() *APIMongoDB {
	db := APIMongoDB{
		"mongodb://127.0.0.1:27017",
		"libraryApp",
		"books",
		"users",
	}
	return &db
}

//APIBookURL takes param isbn. Returns URL string to openlibrary.org/api
func APIBookURL(isbn int) string {
	return "http://openlibrary.org/api/books?bibkeys=ISBN:" + strconv.Itoa(isbn) + "&format=json&jscmd=data"
}

//GetBookInfo parses isbn from request and returns the parsed Bookinfo json
func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	var isbn ISBN
	err := json.NewDecoder(r.Body).Decode(&isbn)
	if err != nil || isbn.Isbn == 0 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		db := SetUpDB()
		book, inDB := db.GetBookByISBN(isbn.Isbn)
		if !inDB {
			jsonResp, err := RetrieveBookInfo(isbn.Isbn)
			if err != nil {
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			} else {
				var check bool
				book, check = ParseBookInfo(jsonResp, isbn.Isbn)

				if !check {
					http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				} else {
					db.AddBook(book)
					fmt.Print("Got book from API")
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(book)
				}
			}
		} else {
			fmt.Print("Got book from DB")
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
		}

	}
}

//RetrieveBookInfo takes param isbn. Returns map[string]interface{} with info from bookAPI
func RetrieveBookInfo(isbn int) (jsonResp map[string]interface{}, err error) {
	resp, err := http.Get(APIBookURL(isbn))
	if err == nil {
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
	}
	return
}

//ParseBookInfo takes param map[string]interface{} and isbn. returns Book struct and status bool
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

//RegisterUser registers new user to db, if username is new and unique.
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
				tokenReply(w, user)
			}
		}
	}
}

//LoginUser checks calidates username and password. If it matches, returns token to log in.
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
				tokenReply(w, user)
			}
		}
	}
}

//LogoutUser checks if user exists and is authorized. Logs out unauthorized user
func LogoutUser(w http.ResponseWriter, r *http.Request) {
	var loginStruct LoginStruct
	err := json.NewDecoder(r.Body).Decode(&loginStruct)
	token, header := r.Header["Token"]

	if !header || !validateToken(token, loginStruct.Username) {
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

//OnlineStatusChage takes param user, db and status. Updates user in db with new changes. Is used for more than just OnlineStauts changes, should be renamed
func OnlineStatusChage(user User, db *APIMongoDB, status bool) error {
	user.Status.Online = status
	user.Status.TimeStamp = time.Now().Unix()
	err := db.UpdateUserStatus(user)
	return err
}

//FailedLoginAttempt writes errormessage back to requester
func FailedLoginAttempt(w http.ResponseWriter) {
	http.Error(w, "Username and password does not match database entry", http.StatusUnauthorized)
}

//UserBooks writes a json containing the slice of all ISBN values of a online authorized user
func UserBooks(w http.ResponseWriter, r *http.Request) {
	username, uHeader := r.Header["Username"]
	token, tHeader := r.Header["Token"]
	if !uHeader || !tHeader || !validateToken(token, username[0]) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		db := SetUpDB()
		user, found := db.GetUserByUsername(username[0])
		if !found || !user.Status.Online {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		} else {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(user.BooksISBN)
		}
	}
}

//RegisterUserBook registers posted book to user if online, authorized and the isbn is correct
func RegisterUserBook(w http.ResponseWriter, r *http.Request) {
	var register RegisterBookStruct
	err := json.NewDecoder(r.Body).Decode(&register)
	token, header := r.Header["Token"]
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	} else {
		if !header || !validateToken(token, register.Username) {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		} else {
			db := SetUpDB()
			user, foundUser := db.GetUserByUsername(register.Username)
			if !foundUser || !user.Status.Online {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			} else {
				book, foundBook := db.GetBookByISBN(register.ISBN)
				if !foundBook {
					jsonResp, err := RetrieveBookInfo(register.ISBN)
					if err != nil {
						http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
					} else {
						var check bool
						book, check = ParseBookInfo(jsonResp, register.ISBN)
						if !check {
							http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
						} else {
							db.AddBook(book)
							fmt.Print("Got book from API")
							user.BooksISBN = append(user.BooksISBN, register.ISBN)
							OnlineStatusChage(user, db, true)
						}
					}
				} else {
					user.BooksISBN = append(user.BooksISBN, register.ISBN)
					OnlineStatusChage(user, db, true)
				}
			}
		}
	}
}

//unexported Functions

func tokenReply(w http.ResponseWriter, user User) {
	w.Header().Add("Token", user.Username+"isauthorized")
}

func validateToken(token []string, username string) bool {
	return token[0] == username+"isauthorized"
}
