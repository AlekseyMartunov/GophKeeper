package users

import (
	"context"

	"GophKeeper/internal/server/entity/users"
)

type userRepo interface {
	Save(ctx context.Context, user users.User) error
	GetExternalID(ctx context.Context, user users.User) (string, error)
}

type hasher interface {
	Hash(value string) string
}

type UserService struct {
	hasher hasher
	repo   userRepo
}

func NewUserService(r userRepo, h hasher) *UserService {
	return &UserService{
		hasher: h,
		repo:   r,
	}
}

func (us *UserService) Save(ctx context.Context, user users.User) error {
	user.Password = us.hasher.Hash(user.Password)
	err := us.repo.Save(ctx, user)
	return err
}

func (us *UserService) GetExternalID(ctx context.Context, user users.User) (string, error) {
	ID, err := us.repo.GetExternalID(ctx, user)
	return ID, err
}
