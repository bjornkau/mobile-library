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
	UserName  string         `json:"user_name"`
	Password  string         `json:"password"`
	Email     string         `json:"email"`
	BooksISBN map[int]string `json:"books"`
	Status    Status
}

type ISBN struct {
	Isbn int `json:"isbn"`
}

type Status struct {
	Online    bool  `json:"online"`
	TimeStamp int64 `json:"time_stamp"`
}
