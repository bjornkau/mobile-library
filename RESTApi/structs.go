package main

type Book struct {
	Title       string   `json:"title"`
	Subtitle    string   `json:"subtitle"`
	Pages       int      `json:"pages"`
	Authors     []string `json:"authors"`
	Subjects    []string `json:"subjects"`
	ISBN        int      `json:"isbn"`
	PublishDate string   `json:"publish_date"`
}

//not implmented in the base api
type Author struct {
}

type User struct {
	Username  string `json:"username"` //Has to be unique, used as ID
	Password  string `json:"password"`
	Email     string `json:"email"`
	BooksISBN []int  `json:"books"`
	Status    StatusStruct
}

type ISBN struct {
	Isbn int `json:"isbn"`
}

type StatusStruct struct {
	Online    bool  `json:"online"`
	TimeStamp int64 `json:"time_stamp"`
}

type LoginStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterStruct struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LogOutStruct struct {
	Username string `json:"username"`
}

type RegisterBookStruct struct {
	Username string `json:"username"`
	ISBN     int    `json:"isbn"`
}
