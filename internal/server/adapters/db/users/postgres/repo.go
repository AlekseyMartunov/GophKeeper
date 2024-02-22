package postgres

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserStorage struct {
	conn *pgxpool.Pool
}

func MewUserStorage(p *pgxpool.Pool) *UserStorage {
	return &UserStorage{
		conn: p,
	}
}
