package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go-forum-api/app/models"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"net/http"
	"strconv"
	"strings"
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
	urlGroup.GET("/:id/details", handler.Get)
	urlGroup.POST("/:id/details", handler.Update)

	return handler
}

func (handler *PostHandler) Get(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	detailsRaw := c.Query("related")
	var details []string
	if detailsRaw != "" {
		details = strings.Split(detailsRaw, ",")
	}

	post, err := handler.postUseCase.Get(int(id), details)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, post)
	return
}

func (handler *PostHandler) Update(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	post := &models.Post{}
	err = easyjson.UnmarshalFromReader(c.Request.Body, post)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	post.ID = int(id)

	forum, err := handler.postUseCase.Update(post)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, forum)
	return
}
