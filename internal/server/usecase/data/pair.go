package data

import (
	"context"

	"GophKeeper/internal/server/entity/data"
	"GophKeeper/internal/server/entity/users"
)

type pairRepo interface {
	Save(ctx context.Context, pair data.Pair, user users.User) error
	Get(ctx context.Context, user users.User, ExternalID string) (data.Pair, error)
	GetAll(ctx context.Context, user users.User) ([]data.Pair, error)
	Delete(ctx context.Context, user users.User, ExternalID string) error
}

type PairService struct {
	repo pairRepo
}

func NewPairService(r pairRepo) *PairService {
	return &PairService{
		repo: r,
	}
}

func (ps *PairService) Save(ctx context.Context, pair data.Pair, user users.User) error {
	err := ps.repo.Save(ctx, pair, user)
	return err
}

func (ps *PairService) Get(ctx context.Context, user users.User, ExternalID string) (data.Pair, error) {
	pair, err := ps.repo.Get(ctx, user, ExternalID)
	return pair, err
}

func (ps *PairService) GetAll(ctx context.Context, user users.User) ([]data.Pair, error) {
	pairs, err := ps.repo.GetAll(ctx, user)
	return pairs, err
}

func (ps *PairService) Delete(ctx context.Context, user users.User, ExternalID string) error {
	err := ps.repo.Delete(ctx, user, ExternalID)
	return err
}
