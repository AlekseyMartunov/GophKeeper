package cardhandlers

import (
	"GophKeeper/internal/entity/card"
	"context"
)

//go:generate mockgen -source=pairhandler.go -destination=/mocks/mock.go

type logger interface {
	Info(s string)
	Error(e error)
}

type cardService interface {
	Save(ctx context.Context, card card.Card) error
	Get(ctx context.Context, cardName string, userID int) (card.Card, error)
	GetAll(ctx context.Context, userID int) ([]card.Card, error)
	Delete(ctx context.Context, cardName string, userID int) error
}

type CardHandler struct {
	service cardService
	log     logger
}

func NewCardHandler(l logger, s cardService) *CardHandler {
	return &CardHandler{
		log:     l,
		service: s,
	}
}
