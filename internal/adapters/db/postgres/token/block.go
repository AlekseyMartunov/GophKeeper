package tokenrepo

import (
	"GophKeeper/internal/entity/token"
	"context"
)

func (ts *TokenStorage) BlockToken(ctx context.Context, t string) error {
	query := `UPDATE tokens SET is_blocked = TRUE WHERE token = $1`

	tag, err := ts.pool.Exec(ctx, query, t)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return token.ErrNoTokenFound
	}

	return nil

}
