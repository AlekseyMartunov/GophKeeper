package cardsrepo

import (
	"GophKeeper/internal/server/entity/card"
	"context"
)

func (cs *CardStorage) Delete(ctx context.Context, cardName string, userID int) error {
	query := `DELETE FROM cards WHERE fk_user_id = $1 AND card_name = $2`

	tag, err := cs.conn.Exec(ctx, query, userID, cardName)
	if err != nil {
		return err
	}
	if tag.RowsAffected() == 0 {
		return card.ErrCardNothingToDelete
	}
	return nil
}
