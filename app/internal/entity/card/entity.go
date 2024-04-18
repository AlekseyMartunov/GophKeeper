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
	Number string
	Owner  string
	Date   string
	CVV    string

	CreatedTime time.Time
	UserID      int
	Name        string
	ID          int
}
