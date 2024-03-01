package cardsrepo

import (
	"context"
	"errors"

	"GophKeeper/internal/server/entity/card"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
)

func (cs *CardStorage) Save(ctx context.Context, c card.Card) error {
	query := `INSERT INTO cards (card_name, card_number, owner, cvv, card_date, created_time,fk_user_id) 
				values ($1, $2, $3, $4, $5);`

	_, err := cs.conn.Exec(ctx, query, c.Name, c.Number, c.Owner, c.CVV, c.Date, c.CreatedTime, c.User.ID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgerrcode.IsIntegrityConstraintViolation(pgErr.Code) {
			return card.ErrCardAlreadyExists
		}

		return err
	}

	return nil
}
