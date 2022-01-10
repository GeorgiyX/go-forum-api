package errors

import "errors"

var (
	ErrUserNotFound = errors.New("не найден юзер")
	ErrUserUpdate   = errors.New("новые данные профиля пользователя конфликтуют с имеющимися пользователями")
	ErrBadRequest   = errors.New("bad request")
)
