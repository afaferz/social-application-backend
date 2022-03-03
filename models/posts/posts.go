package posts

import "time"

type Comment struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Post     string    `json:"post"`
	Date     time.Time `json:"date"`
}

type Post struct {
	Id       int       `json:"id"`
	Username string    `json:"username"`
	Post     string    `json:"post"`
	Date     time.Time `json:"date"`
	Comments []Comment `json:"comments"`
}
