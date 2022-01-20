package handlers

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/models"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"go-forum-api/utils/validator"
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
	forum := &models.Forum{}
	forum.Slug = c.Param("slug")
	if v, _ := validator.GetInstance(); !v.ValidateSlug(forum.Slug) {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.SetDetails("Не корректный slug"))
		return
	}

	params := &models.ForumGetUsersQueryParams{}
	err := c.Bind(params)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

}

func (handler *ForumHandler) Create(c *gin.Context) {
}

func (handler *ForumHandler) GetUsers(c *gin.Context) {
}
