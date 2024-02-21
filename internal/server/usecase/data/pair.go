package data

import (
	"GophKeeper/internal/server/entity/data"
	"GophKeeper/internal/server/entity/users"
	"context"
)

type pairRepo interface {
	Save(ctx context.Context, pair data.Pair, user users.User) error
	Get(ctx context.Context, user users.User, ExternalID string) (data.Pair, error)
	GetAll(ctx context.Context, user users.User) ([]data.Pair, error)
	Delete(ctx context.Context, user users.User, ExternalID string) error
}

type coder interface {
	Encode(value string) string
	Decode(value string) string
	DecodeForAll([]data.Pair) []data.Pair
}

type PairService struct {
	repo  pairRepo
	coder coder
}

func NewPairService(r pairRepo, c coder) *PairService {
	return &PairService{
		repo:  r,
		coder: c,
	}
}

func (ps *PairService) Save(ctx context.Context, pair data.Pair, user users.User) error {
	pair.Password = ps.coder.Encode(pair.Password)
	pair.Login = ps.coder.Encode(pair.Login)

	err := ps.repo.Save(ctx, pair, user)
	return err
}

func (ps *PairService) Get(ctx context.Context, user users.User, ExternalID string) (data.Pair, error) {
	pair, err := ps.repo.Get(ctx, user, ExternalID)
	if err != nil {
		return pair, err
	}
	pair.Login = ps.coder.Decode(pair.Login)
	pair.Password = ps.coder.Decode(pair.Password)

	return pair, nil
}

func (ps *PairService) GetAll(ctx context.Context, user users.User) ([]data.Pair, error) {
	pairs, err := ps.repo.GetAll(ctx, user)
	if err != nil {
		return nil, err
	}

	pairs = ps.coder.DecodeForAll(pairs)
	return pairs, err
}

func (ps *PairService) Delete(ctx context.Context, user users.User, ExternalID string) error {
	err := ps.repo.Delete(ctx, user, ExternalID)
	return err
}
