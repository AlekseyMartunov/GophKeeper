package tokenservice

import (
	"GophKeeper/app/internal/entity/token"
	"GophKeeper/app/internal/entity/users"
	"context"
	"time"
)

type config interface {
	Salt() string
}

type storage interface {
	SaveToken(ctx context.Context, t token.Token) error
	LockToken(ctx context.Context, tokenName string, userID int, status bool) error
	GetTokenInfo(ctx context.Context, t string) (token.Token, error)
	GetAll(ctx context.Context, userID int) ([]string, error)
}

type tokenManager interface {
	CreateToken(ID string) (string, error)
	GetExternalUserID(tokenString string) (string, error)
}

type hasher interface {
	Hash(value, salt string) string
}
type TokenService struct {
	repo    storage
	manager tokenManager
	hasher  hasher
	cfg     config
}

func NewTokenService(s storage, h hasher, m tokenManager, c config) *TokenService {
	return &TokenService{
		repo:    s,
		manager: m,
		hasher:  h,
		cfg:     c,
	}
}

func (ts *TokenService) CreateAndSave(ctx context.Context, u users.User, ip, tokenName string) (string, error) {
	t, err := ts.manager.CreateToken(u.ExternalID)
	if err != nil {
		return "", err
	}

	tokenEntity := token.Token{
		Name:           tokenName,
		Token:          ts.hasher.Hash(t, ts.cfg.Salt()),
		CreatedTime:    time.Now(),
		ExternalUserID: u.ExternalID,
		InternalUserID: u.ID,
		IsBlocked:      false,
		IpAddress:      ip,
	}

	err = ts.repo.SaveToken(ctx, tokenEntity)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (ts *TokenService) GetAll(ctx context.Context, userID int) ([]string, error) {
	return ts.repo.GetAll(ctx, userID)
}

func (ts *TokenService) GetTokenInfo(ctx context.Context, t string) (token.Token, error) {
	return ts.repo.GetTokenInfo(ctx, ts.hasher.Hash(t, ts.cfg.Salt()))
}

func (ts *TokenService) GetExternalUserID(token string) (string, error) {
	return ts.manager.GetExternalUserID(token)
}

func (ts *TokenService) LockToken(ctx context.Context, tokenName string, userID int) error {
	return ts.repo.LockToken(ctx, tokenName, userID, true)
}
