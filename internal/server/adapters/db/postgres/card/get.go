package cardsrepo

import (
	"context"
	"errors"

	"GophKeeper/internal/server/entity/card"
	"GophKeeper/internal/server/entity/users"

	"github.com/jackc/pgx/v5"
)

func (cs *CardStorage) Get(ctx context.Context, cardName string, userID int) (card.Card, error) {
	query := `SELECT card_id, card_name, card_number, owner, cvv, card_date, created_time, fk_user_id FROM cards 
			  WHERE pair_name = $1 AND fk_user_id = $2`

	row := cs.conn.QueryRow(ctx, query, cardName, userID)
	var c card.Card
	var u users.User
	c.User = u

	err := row.Scan(&c.ID, &c.Name, &c.Number, &c.Owner, &c.CVV, c.Date, c.CreatedTime, &c.User.ID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return c, card.ErrCardDoseNotExist
		}
		return c, err
	}

	return c, nil
}
