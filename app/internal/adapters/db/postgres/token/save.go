package tokenrepo

import (
	"GophKeeper/app/internal/entity/token"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (ts *TokenStorage) SaveToken(ctx context.Context, t token.Token) error {
	query := `INSERT INTO tokens (token_name, client_ip, token, created_time, external_user_id, fk_user_id)
				VALUES ($1, $2, $3, $4, $5, $6) ON CONFLICT (fk_user_id, client_ip)
				DO UPDATE SET
				token_name = $1,
				token = $3,
				created_time = $4`

	_, err := ts.pool.Exec(ctx, query, t.Name, t.IpAddress, t.Token, t.CreatedTime, t.ExternalUserID, t.InternalUserID)
	if err != nil {
		fmt.Println(err)
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return token.ErrTokenAlreadyExists
		}
		return err
	}
	return nil
}
