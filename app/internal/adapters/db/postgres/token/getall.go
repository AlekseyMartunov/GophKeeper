package tokenrepo

import (
	"GophKeeper/app/internal/entity/token"
	"context"
	"fmt"
)

func (ts *TokenStorage) GetAll(ctx context.Context, userID int) ([]string, error) {
	query := `SELECT token_name FROM tokens WHERE fk_user_id = $1`
	fmt.Printf("get all token user id: %d", userID)

	rows, err := ts.pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	res := make([]string, 0, 10)
	var s string

	for rows.Next() {
		rows.Scan(&s)
		res = append(res, s)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, token.ErrNoTokenFound
	}

	return res, nil

}
