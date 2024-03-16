package tokenrepo

import (
	"GophKeeper/internal/entity/token"
	"context"
	"errors"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (ts *TokenStorage) SaveToken(ctx context.Context, t token.Token) error {
	query := `INSERT INTO tokens (token_name, token, created_time, external_user_id, fk_user_id)
				VALUES ($1, $2, $3, $4, $5)`

	_, err := ts.pool.Exec(ctx, query, t.Name, t.Token, t.CreatedTime, t.ExternalUserID, t.InternalUserID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return token.ErrTokenAlreadyExists
		}
		return err
	}
	return nil
}
