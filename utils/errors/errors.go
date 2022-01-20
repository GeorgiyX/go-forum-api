package errors

import (
	"go-forum-api/app/models"
	"net/http"
)

type IAPIErrors interface {
	Error() string
	SetDetails(text string) *models.Message
	Code() int
}

var (
	ErrUserNotFound       IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "не найден юзер"}
	ErrUserUpdateNotFound IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "не найден пользователь для обновления"}
	ErrUserUpdateConflict IAPIErrors = &models.Message{ErrCode: http.StatusConflict, Description: "новые данные профиля пользователя конфликтуют с имеющимися пользователями"}
	ErrUserCreateConflict IAPIErrors = &models.Message{ErrCode: http.StatusConflict, Description: "пользователь c таким email или nickname уже существует"}
	ErrBadRequest         IAPIErrors = &models.Message{ErrCode: http.StatusBadRequest, Description: "bad request"}
	ErrInternalServer     IAPIErrors = &models.Message{ErrCode: http.StatusInternalServerError, Description: "internal server error"}
)

var (
	SQL23505 = "23505"
)
