package pairsrepo

import (
	"GophKeeper/internal/entity/pairs"
	"context"
)

func (ps *PairStorage) Delete(ctx context.Context, pairName string, userID int) error {
	query := `DELETE FROM pairs WHERE fk_user_id = $1 AND pair_name = $2`

	tag, err := ps.conn.Exec(ctx, query, userID, pairName)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return pairs.ErrPairNothingToDelete
	}
	return nil
}
