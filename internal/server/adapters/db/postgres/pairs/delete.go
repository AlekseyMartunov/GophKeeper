package pairsrepo

import (
	"GophKeeper/internal/server/entity/pairs"
	"GophKeeper/internal/server/entity/users"
	"context"
)

func (ps *PairStorage) Delete(ctx context.Context, user users.User, name string) error {
	query := `DELETE FROM pairs WHERE fk_user_id = $1 AND pair_name = $2`

	tag, err := ps.conn.Exec(ctx, query, user.ID, name)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pairs.ErrPairNothingToDelete
	}
	return nil
}
