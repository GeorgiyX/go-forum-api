package repositories

import (
	"go-forum-api/app/models"
)

type IThreadRepository interface {
	GetBySlug(slug string) (thread *models.Thread, err error)
	GetByID(id int) (thread *models.Thread, err error)
	UpdateBySlug(thread *models.Thread) (updatedThread *models.Thread, err error)
	UpdateByID(thread *models.Thread) (updatedThread *models.Thread, err error)
	VoteBySlug(slug string, vote *models.Vote) (err error)
	VoteByID(threadId int, vote *models.Vote) (err error)
	CreatePosts(threadId int, forumSlug string, post []*models.Post) (createdPosts []*models.Post, err error)
	//GetPost(post *models.Post, threadId int /*todo sort param*/) (createdPost *models.Post, err error)

}
