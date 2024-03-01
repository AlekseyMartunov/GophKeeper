package pairsrepo

import (
	"context"

	"GophKeeper/internal/server/entity/pairs"
	"GophKeeper/internal/server/entity/users"
)

func (ps *PairStorage) GetAll(ctx context.Context, userID int) ([]pairs.Pair, error) {
	query := `SELECT pair_id, pair_name, password, login, created_time, fk_user_id FROM pairs 
			  WHERE fk_user_id = $1`

	rows, err := ps.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	pairsArr := make([]pairs.Pair, 0, 10)
	for rows.Next() {
		p := pairs.Pair{}
		u := users.User{}

		err := rows.Scan(&p.ID, &p.Name, &p.Password, &p.Login, &p.CreatedTime, &u.ID)
		if err != nil {
			return nil, err
		}
		p.User = u
		pairsArr = append(pairsArr, p)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(pairsArr) == 0 {
		return nil, pairs.ErrPairDoseNotExist
	}

	return pairsArr, nil
}
