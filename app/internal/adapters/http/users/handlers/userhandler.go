package userhandlers

import (
	"GophKeeper/app/internal/entity/users"
	"context"
)

//go:generate mockgen -source=handlers/handler.go -destination=tests/mock/mock.go

type userService interface {
	GetUserInfo(ctx context.Context, login, password string) (users.User, error)
	Save(ctx context.Context, user users.User) error
}

type logger interface {
	Info(s string)
	Error(e error)
}

type UserHandler struct {
	service userService
	log     logger
}

func NewUserHandler(s userService, l logger) *UserHandler {
	return &UserHandler{
		service: s,
		log:     l,
	}
}
