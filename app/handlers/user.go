package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go-forum-api/app/models"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"net/http"
)

type UserHandler struct {
	UserUseCase usecases.IUserUseCase
}

func CreateUserHandler(url string,
	userUseCase usecases.IUserUseCase,
	router *gin.RouterGroup) *UserHandler {
	handler := &UserHandler{
		UserUseCase: userUseCase,
	}

	urlGroup := router.Group(url)
	urlGroup.GET("/:nickname/profile", handler.Get)
	urlGroup.POST("/:nickname/create", handler.Create)

	return handler
}

func (handler *UserHandler) Get(c *gin.Context) {
	nickname := c.Param("nickname")
	model, err := handler.UserUseCase.Get(&nickname)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err.(errors.IAPIErrors).ToMessage())
		return
	}
	c.JSON(http.StatusOK, model)
}

func (handler *UserHandler) Create(c *gin.Context) {
	model := &models.User{}
	model.NickName = c.Param("nickname")
	err := easyjson.UnmarshalFromReader(c.Request.Body, model)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest.ToMessage())
		return
	}

	err = handler.UserUseCase.Create(model)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err.(errors.IAPIErrors).ToMessage())
		return
	}

	c.JSON(http.StatusOK, model)
}
