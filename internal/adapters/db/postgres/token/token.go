package tokenrepo

import "github.com/jackc/pgx/v5/pgxpool"

type TokenStorage struct {
	pool *pgxpool.Pool
}

func NewTokenStorage(p *pgxpool.Pool) *TokenStorage {
	return &TokenStorage{pool: p}
}
