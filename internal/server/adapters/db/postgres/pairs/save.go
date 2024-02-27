package pairsrepo

import (
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"GophKeeper/internal/server/entity/pairs"
)

func (ps *PairStorage) Save(ctx context.Context, pair pairs.Pair) error {
	query := `INSERT INTO pairs (pair_name, password, login, created_time, fk_user_id) 
				values ($1, $2, $3, $4, $5);`

	_, err := ps.conn.Exec(ctx, query, pair.Name, pair.Password, pair.Login, pair.CreatedTime, pair.User.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return pairs.ErrPairAlreadyExists
		}
		return err
	}

	return nil
}
