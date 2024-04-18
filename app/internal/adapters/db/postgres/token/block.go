package tokenrepo

import (
	"GophKeeper/internal/entity/token"
	"context"
)

func (ts *TokenStorage) LockToken(ctx context.Context, tokenName string, userID int, status bool) error {
	query := `UPDATE tokens SET is_blocked = $1 WHERE token_name = $2 AND fk_user_id = $3`

	tag, err := ts.pool.Exec(ctx, query, status, tokenName, userID)
	if err != nil {
		return err
	}

	if tag.RowsAffected() == 0 {
		return token.ErrNoTokenFound
	}

	return nil
}
