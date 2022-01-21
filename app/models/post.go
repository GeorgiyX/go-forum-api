package models

import "time"

type Post struct {
	ID       int       `json:"id"`
	Parent   int       `json:"parent"`
	Author   string    `json:"author"`
	Forum    string    `json:"forum"`
	Thread   int       `json:"thread"`
	Created  time.Time `json:"created"`
	IsEdited bool      `json:"isEdited"`
	Message  string    `json:"message"`
}

//easyjson:json
type Posts []*Post
