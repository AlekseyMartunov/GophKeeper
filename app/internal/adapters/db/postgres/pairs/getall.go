package pairsrepo

import (
	"GophKeeper/internal/entity/pairs"
	"context"
)

func (ps *PairStorage) GetAll(ctx context.Context, userID int) ([]pairs.Pair, error) {
	query := `SELECT pair_name, created_time FROM pairs 
			  WHERE fk_user_id = $1`

	rows, err := ps.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	pairsArr := make([]pairs.Pair, 0, 10)
	for rows.Next() {
		p := pairs.Pair{}

		err = rows.Scan(&p.Name, &p.CreatedTime)
		if err != nil {
			return nil, err
		}
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
