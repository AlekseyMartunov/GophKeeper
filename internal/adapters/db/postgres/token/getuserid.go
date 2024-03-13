package tokenrepo

import "context"

func (ts *TokenStorage) GetUserIDByExternalID(ctx context.Context, t string) (int, error) {
	query := `SELECT fk_user_id FROM tokens WHERE token = $1`

	row := ts.pool.QueryRow(ctx, query, t)

	var userID int
	err := row.Scan(&userID)
	if err != nil {
		return userID, err
	}

	return userID, nil
}
