package main

import (
	"fmt"
	"net/http"
)

//POST
func BookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		GetBookInfo(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//POST
func AuthorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting Author")
}

//POST
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Loging in user")
}

//POST
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Loging out")
}

//GET
func UserBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting books for user")
}

//GET
func UserAuthorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting authors for user")
}
