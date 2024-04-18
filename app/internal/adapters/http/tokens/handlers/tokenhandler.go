package tokenhandlers

import (
	"GophKeeper/app/internal/entity/users"
	"context"
)

type logger interface {
	Info(s string)
	Error(e error)
}

type tokenService interface {
	LockToken(ctx context.Context, token string, userID int) error
	CreateAndSave(ctx context.Context, u users.User, ip, tokenName string) (string, error)
	GetAll(ctx context.Context, userID int) ([]string, error)
}

type userService interface {
	GetUserInfo(ctx context.Context, login, password string) (users.User, error)
}

type TokenHandler struct {
	tokenService tokenService
	userService  userService
	log          logger
}

func NewTokenHandler(l logger, s tokenService, u userService) *TokenHandler {
	return &TokenHandler{
		tokenService: s,
		userService:  u,
		log:          l,
	}
}
