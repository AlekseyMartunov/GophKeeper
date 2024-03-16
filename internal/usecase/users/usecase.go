package usersservice

import (
	"GophKeeper/internal/entity/users"
	"context"
)

type userRepo interface {
	Save(ctx context.Context, user users.User) error
	GetUserInfo(ctx context.Context, login, password string) (users.User, error)
	GetInternalUserID(ctx context.Context, ExternalID string) (int, error)
}

type hasher interface {
	Hash(value, salt string) string
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
	user.Password = us.hasher.Hash(user.Password, user.Login)
	err := us.repo.Save(ctx, user)
	return err
}

func (us *UserService) GetUserInfo(ctx context.Context, login, password string) (users.User, error) {
	password = us.hasher.Hash(password, login)
	return us.repo.GetUserInfo(ctx, login, password)
}

func (us *UserService) GetInternalUserID(ctx context.Context, ExternalID string) (int, error) {
	return us.repo.GetInternalUserID(ctx, ExternalID)
}
