package userhandlers

import (
	"GophKeeper/internal/server/entity/users"
	"context"
)

type userService interface {
	GetExternalID(ctx context.Context, user users.User) (string, error)
	Save(ctx context.Context, user users.User) error
}

type tokenJWTManager interface {
	CreateToken(ID string) (string, error)
}

type logger interface {
	Info(s string)
	Error(e error)
}

type UserHandler struct {
	service userService
	log     logger
	jwt     tokenJWTManager
}

func NewUserHandler(s userService, l logger, t tokenJWTManager) *UserHandler {
	return &UserHandler{
		service: s,
		log:     l,
		jwt:     t,
	}
}
