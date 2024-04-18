package cardsrepo

import (
	"GophKeeper/internal/entity/card"
	"context"
)

func (cs *CardStorage) GetAll(ctx context.Context, userID int) ([]card.Card, error) {
	query := `SELECT card_name, created_time FROM cards 
			  WHERE fk_user_id = $1`

	rows, err := cs.conn.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	pairsArr := make([]card.Card, 0, 10)
	for rows.Next() {
		c := card.Card{}

		err = rows.Scan(&c.Name, &c.CreatedTime)
		if err != nil {
			return nil, err
		}
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
