package main

import "net/http"

func main() {
	http.HandleFunc("/library/api/book", BookHandler)
	http.HandleFunc("/library/api/author", AuthorHandler)
	http.HandleFunc("/library/users/login", LoginHandler)
	http.HandleFunc("/library/users/logout", BookHandler)
	http.HandleFunc("/library/users/books", BookHandler)
	http.HandleFunc("/library/users/authors", BookHandler)
	http.ListenAndServe(":5050", nil)
}
