package main

//Book contains information on books. Parsed from API
type Book struct {
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Pages       int      `json:"pages"`
	Authors     []string `json:"authors"`
	Subjects    []string `json:"subjects"`
	ISBN        int      `json:"isbn"`
	PublishDate string   `json:"publish_date"`
}

//Author contains information about authors. NOT IMPLEMENTED
type Author struct {
}

//User contains information about users
type User struct {
	Username  string `json:"username"` //Has to be unique, used as ID
	Password  string `json:"password"`
	Email     string `json:"email"`
	BooksISBN []int  `json:"books"`
	Status    StatusStruct
}

//ISBN is used to decode request json
type ISBN struct {
	Isbn int `json:"isbn"`
}

//StatusStruct contains information about online status
type StatusStruct struct {
	Online    bool  `json:"online"`
	TimeStamp int64 `json:"time_stamp"`
}

//LoginStruct is used to decode request json
type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//RegisterStruct is used to decode request json
type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

//LogOutStruct is used to decode request json
type LogOutStruct struct {
	Username string `json:"username"`
}

//RegisterBookStruct is used to decode request json
type RegisterBookStruct struct {
	Username string `json:"username"`
	ISBN     int    `json:"isbn"`
}
