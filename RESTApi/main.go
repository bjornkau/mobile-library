package main

import "net/http"
import "os"

func main() {
	http.HandleFunc("/library/api/book", BookHandler)
	http.HandleFunc("/library/api/author", AuthorHandler)
	http.HandleFunc("/library/users/login", LoginHandler)
	http.HandleFunc("/library/users/register", RegisterHandler)
	http.HandleFunc("/library/users/logout", LogoutHandler)
	http.HandleFunc("/library/users/books", UserBookHandler)
	http.HandleFunc("/library/users/authors", UserAuthorHandler)
	http.HandleFunc("/library/users/registerbook", UserRegisterBookHandler)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
