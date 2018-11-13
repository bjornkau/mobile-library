package main

import (
	"net/http"
	"strconv"
	"encoding/json"
	"fmt"
)

func APIBookURL(isbn int) string {
	return "http://openlibrary.org/api/books?bibkeys=ISBN:" + strconv.Itoa(isbn) + "&format=json&jscmd=data"
}

func GetBookInfo(w http.ResponseWriter, r *http.Request) {
	var isbn ISBN
	err := json.NewDecoder(r.Body).Decode(&isbn)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest),http.StatusBadRequest)
		} else {
			jsonResp, err := RetrieveBookInfo(isbn.Isbn)
			if err != nil {
				http.Error(w, http.StatusText(418),418)
			} else {
				//fmt.Fprint(w,jsonResp)
				ParseBookInfo(jsonResp,isbn.Isbn,w)
			}

		}
}

func RetrieveBookInfo(isbn int) (jsonResp map[string]interface{}, err error){
	resp, err := http.Get(APIBookURL(isbn))
	
	if err != nil {
		return
	} else {
		err = json.NewDecoder(resp.Body).Decode(&jsonResp)
		return
	}
}

/*type Book struct {
	Title       string            `json:"title"`
	Subtitle    string            `json:"subtitle"`
	Publishers  []string          `json:"publishers"`
	Pages       int               `json:"pages"`
	CoverURL    map[string]string `json:"cover"`
	Authors     map[string]string `json:"authors"`
	Subjects    []string          `json:"subjects"`
	ISBN        int               `json:"isbn"`
	PublishDate string            `json:"publish_date"`
}*/

func ParseBookInfo(jsonResp map[string]interface{}, isbn int, w http.ResponseWriter)(book Book, allIsWell bool){
	var info map[string]interface{}
	allIsWell = true
	bookInfo, ok := jsonResp["ISBN:"+strconv.Itoa(isbn)].(map[string]interface{})
	if ok {
		fmt.Fprint(w,info)
		var title bool
		book.Title, title = bookInfo["title"].(string)

		} else {
			allIsWell = false
			return
		}
	return
}

