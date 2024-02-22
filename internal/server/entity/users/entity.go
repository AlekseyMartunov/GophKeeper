package users

import "errors"

var ErrUserAlreadyExists = errors.New("user already exists")
var ErrUserDoseNotExist = errors.New("user dose not exists")

type User struct {
	Login    string
	Password string

	ID         int
	ExternalID string
}
