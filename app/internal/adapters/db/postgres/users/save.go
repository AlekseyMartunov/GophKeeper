package usersrepo

import (
	"GophKeeper/internal/entity/users"
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (us *UserStorage) Save(ctx context.Context, user users.User) error {
	query := `INSERT INTO users (login, password) VALUES ($1, $2)`
	_, err := us.conn.Exec(ctx, query, user.Login, user.Password)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return users.ErrUserAlreadyExists
		}
		return err
	}

	return nil
}
