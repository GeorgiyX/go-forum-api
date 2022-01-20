package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go-forum-api/app/models"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"go-forum-api/utils/validator"
	"net/http"
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
	urlGroup.GET("/:slug/threads", handler.GetThreads)
	urlGroup.POST("/:slug/create", handler.CreateThread)

	return handler
}

func (handler *ForumHandler) Get(c *gin.Context) {
	slug := c.Param("slug")
	if v, _ := validator.GetInstance(); !v.ValidateSlug(slug) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug"))
		return
	}

	forum, err := handler.ForumUseCase.Get(slug)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, forum)
}

func (handler *ForumHandler) Create(c *gin.Context) {
	forum := &models.Forum{}

	err := easyjson.UnmarshalFromReader(c.Request.Body, forum)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	if v, _ := validator.GetInstance(); !v.ValidateSlug(forum.Slug) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug"))
		return
	}

	createdForum, err := handler.ForumUseCase.Create(forum)

	if err != nil {
		if err.(errors.IAPIErrors).Code() == errors.ErrForumAlreadyExists.Code() {
			c.JSON(errors.ErrForumAlreadyExists.Code(), createdForum)
			return
		}
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusCreated, createdForum)
}

func (handler *ForumHandler) GetUsers(c *gin.Context) {
	forum := &models.Forum{}
	forum.Slug = c.Param("slug")
	if v, _ := validator.GetInstance(); !v.ValidateSlug(forum.Slug) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug"))
		return
	}

	params := &models.ForumQueryParams{}
	err := c.Bind(params)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	fmt.Printf("params: %v", params)
}

func (handler *ForumHandler) CreateThread(c *gin.Context) {
	thread := &models.Thread{}

	thread.Forum = c.Param("slug")
	if v, _ := validator.GetInstance(); !v.ValidateSlug(thread.Forum) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug forum"))
		return
	}

	err := easyjson.UnmarshalFromReader(c.Request.Body, thread)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	if v, _ := validator.GetInstance(); thread.Slug != "" && !v.ValidateSlug(thread.Slug) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug thread"))
		return
	}

	createdThread, err := handler.ForumUseCase.CreateThread(thread)
	if err != nil {
		if err.(errors.IAPIErrors).Code() == errors.ErrThreadAlreadyExists.Code() {
			c.JSON(errors.ErrThreadAlreadyExists.Code(), createdThread)
			return
		}
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusCreated, createdThread)
	return
}

func (handler *ForumHandler) GetThreads(c *gin.Context) {
	slug := c.Param("slug")
	if v, _ := validator.GetInstance(); !v.ValidateSlug(slug) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug"))
		return
	}

	params := &models.ForumQueryParams{}
	err := c.ShouldBindQuery(params)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректные query params"))
	}

	v, _ := validator.GetInstance()
	v.ValidateForumQuery(params)

	threads, err := handler.ForumUseCase.GetThreads(slug, params)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, threads)
	return
}
