package cardservice

import (
	"GophKeeper/internal/entity/card"
	"context"
)

type cardRepo interface {
	Save(ctx context.Context, card card.Card) error
	Get(ctx context.Context, cardName string, userID int) (card.Card, error)
	GetAll(ctx context.Context, userID int) ([]card.Card, error)
	Delete(ctx context.Context, cardName string, userID int) error
}

type CardService struct {
	repo cardRepo
}

func NewCardService(r cardRepo) *CardService {
	return &CardService{repo: r}
}

func (cs *CardService) Save(ctx context.Context, card card.Card) error {
	return cs.repo.Save(ctx, card)
}

func (cs *CardService) Get(ctx context.Context, cardName string, userID int) (card.Card, error) {
	return cs.repo.Get(ctx, cardName, userID)
}

func (cs *CardService) GetAll(ctx context.Context, userID int) ([]card.Card, error) {
	return cs.repo.GetAll(ctx, userID)
}

func (cs *CardService) Delete(ctx context.Context, cardName string, userID int) error {
	return cs.repo.Delete(ctx, cardName, userID)
}
