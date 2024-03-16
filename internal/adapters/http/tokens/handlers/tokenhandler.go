package tokenhandlers

import "context"

type logger interface {
	Info(s string)
	Error(e error)
}

type tokenService interface {
	LockToken(ctx context.Context, token string, status bool, userID int) error
}

type TokenHandler struct {
	service tokenService
	log     logger
}

func NewTokenHandler(l logger, s tokenService) *TokenHandler {
	return &TokenHandler{
		service: s,
		log:     l,
	}
}
