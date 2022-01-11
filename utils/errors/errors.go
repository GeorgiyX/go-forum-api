package errors

import (
	"go-forum-api/app/models"
	"net/http"
)

type APIErrors struct {
	ErrCode     int
	Description string
}

type IAPIErrors interface {
	Error() string
	ToMessage() *models.Message
	Code() int
}

func (apiError *APIErrors) Error() string {
	return apiError.Description
}

func (apiError *APIErrors) ToMessage() *models.Message {
	return &models.Message{MessageText: apiError.Error()}
}

func (apiError *APIErrors) Code() int {
	return apiError.ErrCode
}

var (
	ErrUserNotFound IAPIErrors = &APIErrors{ErrCode: http.StatusNotFound, Description: "не найден юзер"}
	ErrUserUpdate   IAPIErrors = &APIErrors{ErrCode: http.StatusConflict, Description: "новые данные профиля пользователя конфликтуют с имеющимися пользователями"}
	ErrBadRequest   IAPIErrors = &APIErrors{ErrCode: http.StatusBadRequest, Description: "bad request"}
)
