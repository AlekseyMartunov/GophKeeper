package usersrepo

import (
	"GophKeeper/app/internal/entity/users"
	"context"
	"errors"
)

func (us *UserStorage) GetInternalUserID(ctx context.Context, ExternalID string) (int, error) {
	query := `SELECT user_id FROM users WHERE external_id = $1 `
	res := us.conn.QueryRow(ctx, query, ExternalID)

	var ID int

	if err := res.Scan(&ID); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ID, users.ErrUserDoseNotExist
		}
		return ID, err
	}

	return ID, nil
}
