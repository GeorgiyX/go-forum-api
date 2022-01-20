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
	ErrUserNotFound       IAPIErrors = &APIErrors{ErrCode: http.StatusNotFound, Description: "не найден юзер"}
	ErrUserUpdateNotFound IAPIErrors = &APIErrors{ErrCode: http.StatusNotFound, Description: "не найден пользователь для обновления"}
	ErrUserUpdateConflict IAPIErrors = &APIErrors{ErrCode: http.StatusConflict, Description: "новые данные профиля пользователя конфликтуют с имеющимися пользователями"}
	ErrUserCreateConflict IAPIErrors = &APIErrors{ErrCode: http.StatusConflict, Description: "пользователь c таким email или nickname уже существует"}
	ErrBadRequest         IAPIErrors = &APIErrors{ErrCode: http.StatusBadRequest, Description: "bad request"}
	ErrInternalServer     IAPIErrors = &APIErrors{ErrCode: http.StatusInternalServerError, Description: "internal server error"}
)

var (
	SQL23505 = "23505"
)
