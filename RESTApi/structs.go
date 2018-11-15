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

type Author struct {
}

type User struct {
}

type ISBN struct {
	Isbn int `json:"isbn"`
}
