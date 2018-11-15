package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	//"reflect"
)

func APIBookURL(isbn int) string {
	return "http://openlibrary.org/api/books?bibkeys=ISBN:" + strconv.Itoa(isbn) + "&format=json&jscmd=data"
}

func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	var isbn ISBN
	err := json.NewDecoder(r.Body).Decode(&isbn)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	} else {
		jsonResp, err := RetrieveBookInfo(isbn.Isbn)
		if err != nil {
			http.Error(w, http.StatusText(418), 418)
		} else {
			book, check := ParseBookInfo(jsonResp, isbn.Isbn)
			if !check {
				http.Error(w, http.StatusText(500), 500)
			} else {
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

/*type Book struct {
	Authors     []string `json:"authors"`
	Subjects    []string          `json:"subjects"`
}*/

func ParseBookInfo(jsonResp map[string]interface{}, isbn int) (book Book, allIsWell bool) {
	allIsWell = true
	bookInfo, ok := jsonResp["ISBN:"+strconv.Itoa(isbn)].(map[string]interface{})
	if ok {
		//fmt.Fprint(w,info)
		var title bool
		var subtitle bool
		var publishDate bool
		var author bool
		var subject bool
		book.ISBN = isbn
		book.Title, title = bookInfo["title"].(string)
		book.Subtitle, subtitle = bookInfo["subtitle"].(string)
		book.Pages = int(bookInfo["number_of_pages"].(float64))
		book.PublishDate, publishDate = bookInfo["publish_date"].(string)
		//
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
