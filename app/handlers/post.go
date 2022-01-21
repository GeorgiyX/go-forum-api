package handlers

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/usecases"
)

type PostHandler struct {
	postUseCase usecases.IPostUseCase
}

func CreatePostHandler(url string,
	postUseCase usecases.IPostUseCase,
	router *gin.RouterGroup) *PostHandler {
	handler := &PostHandler{
		postUseCase: postUseCase,
	}

	urlGroup := router.Group(url)
	urlGroup.POST("/:id/details", handler.Get)
	urlGroup.GET("/:id/details", handler.Update)

	return handler
}

func (handler *PostHandler) Get(c *gin.Context) {
	return
}

func (handler *PostHandler) Update(c *gin.Context) {
	return
}
