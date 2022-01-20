package handlers

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/usecases"
)

type ForumHandler struct {
	ForumUseCase usecases.IForumUseCase
}

func CreateForumHandler(url string,
	forumUseCase usecases.IForumUseCase,
	router *gin.RouterGroup) *ForumHandler {
	handler := &ForumHandler{
		ForumUseCase: forumUseCase,
	}

	urlGroup := router.Group(url)
	urlGroup.GET("/:slug/details", handler.Get)
	urlGroup.POST("/create", handler.Create)
	urlGroup.POST("/:slug/users", handler.GetUsers)

	return handler
}

func (handler *ForumHandler) Get(c *gin.Context) {
}

func (handler *ForumHandler) Create(c *gin.Context) {
}

func (handler *ForumHandler) GetUsers(c *gin.Context) {
}
