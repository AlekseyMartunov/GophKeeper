package pairservice

import (
	"GophKeeper/app/internal/entity/pairs"
	"context"
	"time"
)

type pairRepo interface {
	Save(ctx context.Context, pair pairs.Pair) error
	Get(ctx context.Context, pairName string, userID int) (pairs.Pair, error)
	GetAll(ctx context.Context, userID int) ([]pairs.Pair, error)
	Delete(ctx context.Context, pairName string, userID int) error
}

type PairService struct {
	repo pairRepo
}

func NewPairService(r pairRepo) *PairService {
	return &PairService{
		repo: r,
	}
}

func (ps *PairService) Save(ctx context.Context, pair pairs.Pair) error {
	pair.CreatedTime = time.Now()
	err := ps.repo.Save(ctx, pair)
	return err
}

func (ps *PairService) Get(ctx context.Context, pairName string, userID int) (pairs.Pair, error) {
	pair, err := ps.repo.Get(ctx, pairName, userID)
	return pair, err
}

func (ps *PairService) GetAll(ctx context.Context, userID int) ([]pairs.Pair, error) {
	pairs, err := ps.repo.GetAll(ctx, userID)
	return pairs, err
}

func (ps *PairService) Delete(ctx context.Context, pairName string, userID int) error {
	err := ps.repo.Delete(ctx, pairName, userID)
	return err
}
