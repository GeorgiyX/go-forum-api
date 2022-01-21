package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go-forum-api/app/models"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"net/http"
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
	urlGroup.POST("/:slug_or_id/vote", handler.Vote)
	urlGroup.POST("/:slug_or_id/create", handler.CreatePosts)

	return handler
}

func (handler *ThreadHandler) Get(c *gin.Context) {
	slugOrId := c.Param("slug_or_id")

	forum, err := handler.ThreadUseCase.Get(slugOrId)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, forum)
	return
}

func (handler *ThreadHandler) Update(c *gin.Context) {
	slugOrId := c.Param("slug_or_id")

	thread := &models.Thread{}
	err := easyjson.UnmarshalFromReader(c.Request.Body, thread)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	forum, err := handler.ThreadUseCase.Update(slugOrId, thread)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, forum)
	return
}

func (handler *ThreadHandler) Vote(c *gin.Context) {
	slugOrId := c.Param("slug_or_id")

	vote := &models.Vote{}
	err := easyjson.UnmarshalFromReader(c.Request.Body, vote)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	forum, err := handler.ThreadUseCase.Vote(slugOrId, vote)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, forum)
	return
}

func (handler *ThreadHandler) CreatePosts(c *gin.Context) {
	slugOrId := c.Param("slug_or_id")

	var posts models.Posts
	err := easyjson.UnmarshalFromReader(c.Request.Body, &posts)

	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	createdPosts, err := handler.ThreadUseCase.CreatePosts(slugOrId, posts)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusCreated, createdPosts)
	return
}
