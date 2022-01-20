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
	ErrBadRequest     IAPIErrors = &models.Message{ErrCode: http.StatusBadRequest, Description: "bad request"}
	ErrInternalServer IAPIErrors = &models.Message{ErrCode: http.StatusInternalServerError, Description: "internal server error"}
)

var (
	ErrUserNotFound       IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "не найден юзер"}
	ErrUserUpdateNotFound IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "не найден пользователь для обновления"}
	ErrUserUpdateConflict IAPIErrors = &models.Message{ErrCode: http.StatusConflict, Description: "новые данные профиля пользователя конфликтуют с имеющимися пользователями"}
	ErrUserCreateConflict IAPIErrors = &models.Message{ErrCode: http.StatusConflict, Description: "пользователь c таким email или nickname уже существует"}
)

var (
	ErrForumUserNotFound  IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "владелец форума не найден"}
	ErrForumAlreadyExists IAPIErrors = &models.Message{ErrCode: http.StatusConflict, Description: "форум уже присутсвует в базе данных"}
	ErrForumNotFound      IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "форум не найден"}
)

var (
	ErrThreadAlreadyExists       IAPIErrors = &models.Message{ErrCode: http.StatusConflict, Description: "тред уже присутсвует в базе данных"}
	ErrThreadUserOrForumNotFound IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "автор треда или форуи не найдены"}
	ErrThreadNotFound            IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "тред не найден"}
	ErrThreadUpdateNotFound      IAPIErrors = &models.Message{ErrCode: http.StatusNotFound, Description: "не найден тред для обновления"}
)

var (
	SQL23505 = "23505" // duplicate key value violates unique constraint
	SQL23503 = "23503" // violates foreign key constraint
	SQL23502 = "23502" // null value violates not-null constraint
)
