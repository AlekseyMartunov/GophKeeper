package pairsrepo

import (
	"GophKeeper/internal/entity/pairs"
	"GophKeeper/internal/entity/users"
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
)

func (ps *PairStorage) Get(ctx context.Context, pairName string, userID int) (pairs.Pair, error) {
	query := `SELECT pair_id, pair_name, password, login, created_time, fk_user_id FROM pairs 
			  WHERE pair_name = $1 AND fk_user_id = $2`

	row := ps.conn.QueryRow(ctx, query, pairName, userID)
	var p pairs.Pair
	var u users.User
	p.User = u

	err := row.Scan(&p.ID, &p.Name, &p.Password, &p.Login, &p.CreatedTime, &p.User.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return p, pairs.ErrPairDoseNotExist
		}
		return p, err
	}

	return p, nil
}
