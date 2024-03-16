package tokenservice

import (
	"GophKeeper/internal/entity/token"
	"GophKeeper/internal/entity/users"
	"context"
	"time"
)

type config interface {
	Salt() string
}

type storage interface {
	SaveToken(ctx context.Context, t token.Token) error
	LockToken(ctx context.Context, token string, status bool) error
	GetTokenInfo(ctx context.Context, t string) (token.Token, error)
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

func (ts *TokenService) CreateAndSave(ctx context.Context, u users.User, tokenName string) (string, error) {
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
	}

	err = ts.repo.SaveToken(ctx, tokenEntity)
	if err != nil {
		return "", err
	}

	return t, nil
}

func (ts *TokenService) GetTokenInfo(ctx context.Context, t string) (token.Token, error) {
	return ts.repo.GetTokenInfo(ctx, ts.hasher.Hash(t, ts.cfg.Salt()))
}

func (ts *TokenService) LockToken(ctx context.Context, tokenString string, status bool, userID int) error {
	t, err := ts.GetTokenInfo(ctx, tokenString)
	if err != nil {
		return err
	}
	if t.InternalUserID != userID {
		return token.ErrNoTokenFound
	}
	hasToken := ts.hasher.Hash(tokenString, ts.cfg.Salt())
	return ts.repo.LockToken(ctx, hasToken, status)
}

func (ts *TokenService) GetExternalUserID(token string) (string, error) {
	return ts.manager.GetExternalUserID(token)
}
