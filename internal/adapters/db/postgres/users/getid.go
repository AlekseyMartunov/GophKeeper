package usersrepo

import (
	"GophKeeper/internal/entity/users"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (us *UserStorage) GetExternalID(ctx context.Context, user users.User) (string, error) {
	query := `SELECT external_id FROM users WHERE login = $1 AND password = $2`
	res := us.conn.QueryRow(ctx, query, user.Login, user.Password)

	var ID string

	if err := res.Scan(&ID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", users.ErrUserDoseNotExist
		}
		return "", err
	}

	return ID, nil
}
