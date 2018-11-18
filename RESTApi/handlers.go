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
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Open Library does not have a good api for author searches")
}

//POST
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LoginUser(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//POST
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LogoutUser(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//GET
func UserBookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting books for user")
}

//GET
func UserAuthorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Getting authors for user")
}

//POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		RegisterUser(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
