package handlers

import (
	"github.com/gin-gonic/gin"
	"go-forum-api/app/usecases"
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
	return
}

func (handler *ServiceHandler) Status(c *gin.Context) {
	return
}
