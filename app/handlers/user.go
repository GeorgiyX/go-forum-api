package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mailru/easyjson"
	"go-forum-api/app/models"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
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
	urlGroup.GET("/:nickname/create", handler.Create)

	return handler
}

func (handler *UserHandler) Get(c *gin.Context) {

}

func (handler *UserHandler) Create(c *gin.Context) {
	model := &models.User{}
	model.NickName = c.Param("nickname")
	err := easyjson.UnmarshalFromReader(c.Request.Body, model)
	if err != nil {
		message := &models.Message{
			Message: errors.ErrBadRequest.Error(),
		}
		c.
	}
}
