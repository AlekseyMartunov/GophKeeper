package cardsrepo

import (
	"GophKeeper/internal/entity/card"
	"GophKeeper/internal/entity/users"
	"context"
)

func (cs *CardStorage) GetAll(ctx context.Context, userID int) ([]card.Card, error) {
	query := `SELECT card_id, card_name, card_number, owner, cvv, card_date, created_time, fk_user_id FROM cards 
			  WHERE fk_user_id = $1`

	rows, err := cs.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	pairsArr := make([]card.Card, 0, 10)
	for rows.Next() {
		c := card.Card{}
		u := users.User{}

		err := rows.Scan(&c.ID, &c.Name, &c.Number, &c.Owner, &c.CVV, c.Date, c.CreatedTime, &c.User.ID)
		if err != nil {
			return nil, err
		}
		c.User = u
		pairsArr = append(pairsArr, c)
	}

	if rows.Err() != nil {
		return nil, err
	}

	if len(pairsArr) == 0 {
		return nil, card.ErrCardDoseNotExist
	}

	return pairsArr, nil
}
