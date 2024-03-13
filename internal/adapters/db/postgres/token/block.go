package tokenrepo

import (
	"context"
)

func (ts *TokenStorage) BlockToken(ctx context.Context, t string) error {
	query := `UPDATE tokens SET is_blocked = TRUE WHERE token = $1`

	_, err := ts.pool.Exec(ctx, query, t)
	if err != nil {
		return err
	}

	return nil

}
