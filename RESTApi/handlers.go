package main

import (
	"fmt"
	"net/http"
)

//BookHandler handles request for book information. Requires method POST
func BookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		GetBookInfo(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//AuthorHandler handles request for author information. Requires method POST. NOT IMPLEMENTED
func AuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Open Library does not have a good api for author searches")
}

//LoginHandler handles login request. Requires method POST
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LoginUser(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//LogoutHandler handles logout request. Requires method POST
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		LogoutUser(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//UserBookHandler handles requests for books read by user. Req uses method GET
func UserBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		UserBooks(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//UserRegisterBookHandler handles request for book information. Requires method POST
func UserRegisterBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		RegisterUserBook(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

//UserAuthorHandler handles requests for author information. NOT IMPLEMENTED
func UserAuthorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintf(w, "Open Library does not have a good api for author searches")
}

//RegisterHandler handles register user requests. Requires method POST
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		RegisterUser(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}
