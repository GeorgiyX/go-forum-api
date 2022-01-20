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
	urlGroup.POST("/:nickname/profile", handler.Update)
	urlGroup.POST("/:nickname/create", handler.Create)

	return handler
}

func (handler *UserHandler) Get(c *gin.Context) {
	nickname := c.Param("nickname")
	model, err := handler.UserUseCase.Get(&nickname)
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}
	c.JSON(http.StatusOK, model)
}

func (handler *UserHandler) Create(c *gin.Context) {
	model := &models.User{}
	model.NickName = c.Param("nickname")
	err := easyjson.UnmarshalFromReader(c.Request.Body, model)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	users, err := handler.UserUseCase.Create(model)
	if users != nil {
		c.JSON(err.(errors.IAPIErrors).Code(), users)
		return
	}

	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusCreated, model)
}

func (handler *UserHandler) Update(c *gin.Context) {
	model := &models.User{}
	model.NickName = c.Param("nickname")
	err := easyjson.UnmarshalFromReader(c.Request.Body, model)
	if err != nil {
		c.AbortWithStatusJSON(errors.ErrBadRequest.Code(), errors.ErrBadRequest)
		return
	}

	user, err := handler.UserUseCase.Update(model)

	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, user)
}
