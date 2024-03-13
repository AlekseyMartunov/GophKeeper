package card

import (
	"errors"
	"time"
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
	UserID      int
	Name        string
	ID          int
}
