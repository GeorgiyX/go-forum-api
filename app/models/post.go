package models

import "time"

type Post struct {
	ID       int64     `json:"id"`
	Parent   int64     `json:"parent"`
	Author   string    `json:"author"`
	Forum    string    `json:"forum"`
	Thread   int32     `json:"thread"`
	Created  time.Time `json:"created"`
	IsEdited bool      `json:"isEdited"`
	Message  string    `json:"message"`
}
