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

type encrypt interface {
	Encrypt(value string) string
	Decrypt(value string) string
	DecryptForAll([]data.Pair) []data.Pair
}

type PairService struct {
	repo      pairRepo
	encrypter encrypt
}

func NewPairService(r pairRepo, e encrypt) *PairService {
	return &PairService{
		repo:      r,
		encrypter: e,
	}
}

func (ps *PairService) Save(ctx context.Context, pair data.Pair, user users.User) error {
	pair.Password = ps.encrypter.Encrypt(pair.Password)
	pair.Login = ps.encrypter.Encrypt(pair.Login)

	err := ps.repo.Save(ctx, pair, user)
	return err
}

func (ps *PairService) Get(ctx context.Context, user users.User, ExternalID string) (data.Pair, error) {
	pair, err := ps.repo.Get(ctx, user, ExternalID)
	if err != nil {
		return pair, err
	}
	pair.Login = ps.encrypter.Decrypt(pair.Login)
	pair.Password = ps.encrypter.Decrypt(pair.Password)

	return pair, nil
}

func (ps *PairService) GetAll(ctx context.Context, user users.User) ([]data.Pair, error) {
	pairs, err := ps.repo.GetAll(ctx, user)
	if err != nil {
		return nil, err
	}

	pairs = ps.encrypter.DecryptForAll(pairs)
	return pairs, err
}

func (ps *PairService) Delete(ctx context.Context, user users.User, ExternalID string) error {
	err := ps.repo.Delete(ctx, user, ExternalID)
	return err
}
