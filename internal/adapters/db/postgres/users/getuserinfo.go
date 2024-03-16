package usersrepo

import (
	"GophKeeper/internal/entity/users"
	"context"
	"errors"
	"github.com/jackc/pgx/v5"
)

func (us *UserStorage) GetUserInfo(ctx context.Context, login, password string) (users.User, error) {
	query := `SELECT user_id, external_id  FROM users WHERE login = $1 AND password = $2`
	res := us.conn.QueryRow(ctx, query, login, password)

	var u users.User

	if err := res.Scan(&u.ID, &u.ExternalID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return u, users.ErrUserDoseNotExist
		}
		return u, err
	}

	return u, nil
}
