package handlers

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/usecases"
	"go-forum-api/utils/errors"
	"net/http"
)

type ServiceHandler struct {
	serviceUseCase usecases.IServiceUseCase
}

func CreateServiceHandler(url string,
	serviceUseCase usecases.IServiceUseCase,
	router *gin.RouterGroup) *ServiceHandler {
	handler := &ServiceHandler{
		serviceUseCase: serviceUseCase,
	}

	urlGroup := router.Group(url)
	urlGroup.POST("/clear", handler.Clear)
	urlGroup.GET("/status", handler.Status)

	return handler
}

func (handler *ServiceHandler) Clear(c *gin.Context) {
	err := handler.serviceUseCase.Clear()
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.Status(http.StatusOK)
}

func (handler *ServiceHandler) Status(c *gin.Context) {
	status, err := handler.serviceUseCase.Status()
	if err != nil {
		c.AbortWithStatusJSON(err.(errors.IAPIErrors).Code(), err)
		return
	}

	c.JSON(http.StatusOK, status)
}
