package tokenrepo

import (
	"GophKeeper/internal/entity/token"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (ts *TokenStorage) GetTokenInfo(ctx context.Context, tokenString string) (token.Token, error) {
	query := `SELECT token_id, token_name, created_time, is_blocked, external_user_id, fk_user_id FROM tokens
				WHERE token = $1`

	row := ts.pool.QueryRow(ctx, query, tokenString)

	t := token.Token{}
	err := row.Scan(&t.TokenID, &t.Name, &t.CreatedTime, &t.IsBlocked, &t.ExternalUserID, &t.InternalUserID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return t, token.ErrNoTokenFound
		}
		return t, err
	}

	return t, nil
}
