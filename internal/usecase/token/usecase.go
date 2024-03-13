package tokenservice

import (
	"GophKeeper/internal/entity/token"
	"context"
)

type storage interface {
	Save(ctx context.Context, t token.Token) error
	GetUserIDByExternalID(ctx context.Context, externalID string) (int, error)
	BlockToken(ctx context.Context, token string) error
}

type TokenService struct {
	repo storage
}

func (ts *TokenService) Save(ctx context.Context, t token.Token) error {
	return ts.repo.Save(ctx, t)
}

func (ts *TokenService) GetUserIDByExternalID(ctx context.Context, externalID string) (int, error) {
	return ts.repo.GetUserIDByExternalID(ctx, externalID)
}

func (ts *TokenService) BlockToken(ctx context.Context, token string) error {
	return ts.repo.BlockToken(ctx, token)
}
