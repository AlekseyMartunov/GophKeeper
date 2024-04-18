package pairhandlers

import (
	"GophKeeper/internal/entity/pairs"
	"context"
)

//go:generate mockgen -source=pairhandler.go -destination=/mocks/mock.go

type logger interface {
	Info(s string)
	Error(e error)
}

type pairService interface {
	Save(ctx context.Context, pair pairs.Pair) error
	Get(ctx context.Context, pairName string, userID int) (pairs.Pair, error)
	GetAll(ctx context.Context, userID int) ([]pairs.Pair, error)
	Delete(ctx context.Context, pairName string, userID int) error
}

type PairHandler struct {
	service pairService
	log     logger
}

func NewPairHandler(l logger, s pairService) *PairHandler {
	return &PairHandler{
		log:     l,
		service: s,
	}
}
