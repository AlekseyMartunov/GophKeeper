package tokenrepo

import (
	"GophKeeper/internal/entity/token"
	"context"
)

func (ts *TokenStorage) LockToken(ctx context.Context, t string, status bool) error {
	query := `UPDATE tokens SET is_blocked = $1 WHERE token = $2`

	tag, err := ts.pool.Exec(ctx, query, status, t)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return token.ErrNoTokenFound
	}

	return nil
}
