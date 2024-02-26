package pairsrepo

import "github.com/jackc/pgx/v5/pgxpool"

type PairStorage struct {
	conn *pgxpool.Pool
}

func NewPairsStorage(p *pgxpool.Pool) *PairStorage {
	return &PairStorage{
		conn: p,
	}
}
