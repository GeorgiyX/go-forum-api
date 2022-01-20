package handlers

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/usecases"
)

type ThreadHandler struct {
	ThreadUseCase usecases.IThreadUseCase
}

func CreateThreadHandler(url string,
	threadUseCase usecases.IThreadUseCase,
	router *gin.RouterGroup) *ThreadHandler {
	handler := &ThreadHandler{
		ThreadUseCase: threadUseCase,
	}

	urlGroup := router.Group(url)
	urlGroup.GET("/:slug_or_id/details", handler.Get)
	urlGroup.POST("/:slug_or_id/details", handler.Update)

	return handler
}

func (handler *ThreadHandler) Get(c *gin.Context) {
	return
}

func (handler *ThreadHandler) Update(c *gin.Context) {
	return
}
