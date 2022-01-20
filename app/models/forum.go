package models

type Forum struct {
	ID      int    `json:"-"`
	Slug    string `json:"slug"`
	Title   string `json:"title"`
	User    string `json:"user"`
	Posts   int    `json:"posts,omitempty"`
	Threads int    `json:"threads,omitempty"`
}

type ForumGetUsersQueryParams struct {
	Limit int    `form:"limit"`
	Since string `form:"since"`
	Desc  bool   `form:"desc"`
}
