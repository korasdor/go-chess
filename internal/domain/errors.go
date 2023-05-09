package domain

import "errors"

var (
	ErrBadRequestFormat = errors.New("bad request: error while parsing json package")

	ErrUserNotFound      = errors.New("user doesn't exists")
	ErrUserAlreadyExists = errors.New("user with such email already exists")
)
