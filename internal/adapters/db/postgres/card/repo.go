package cardsrepo

import "github.com/jackc/pgx/v5/pgxpool"

type CardStorage struct {
	conn *pgxpool.Pool
}

func NewCardStorage(p *pgxpool.Pool) *CardStorage {
	return &CardStorage{
		conn: p,
	}
}
