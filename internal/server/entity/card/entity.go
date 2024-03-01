package card

import (
	"errors"
	"time"

	"GophKeeper/internal/server/entity/users"
)

var (
	ErrCardAlreadyExists   = errors.New("card already exists")
	ErrCardDoseNotExist    = errors.New("card dose not exists")
	ErrCardNothingToDelete = errors.New("nothing to delete")
)

type Card struct {
	Number int
	Owner  string
	Date   time.Time
	CVV    int

	CreatedTime time.Time
	User        users.User
	Name        string
	ID          int
}
