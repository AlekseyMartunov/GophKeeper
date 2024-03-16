package userhandlers

import (
	"GophKeeper/internal/entity/users"
	"context"
)

//go:generate mockgen -source=handlers/handler.go -destination=tests/mock/mock.go

type userService interface {
	GetUserInfo(ctx context.Context, login, password string) (users.User, error)
	Save(ctx context.Context, user users.User) error
}

type tokenService interface {
	CreateAndSave(ctx context.Context, user users.User, tokenName string) (string, error)
}

type logger interface {
	Info(s string)
	Error(e error)
}

type UserHandler struct {
	service      userService
	log          logger
	tokenService tokenService
}

func NewUserHandler(s userService, l logger, t tokenService) *UserHandler {
	return &UserHandler{
		service:      s,
		log:          l,
		tokenService: t,
	}
}
