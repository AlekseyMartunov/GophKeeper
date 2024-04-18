package pairs

import (
	"errors"
	"time"
)

var (
	ErrPairAlreadyExists   = errors.New("pair already exists")
	ErrPairDoseNotExist    = errors.New("pair dose not exists")
	ErrPairNothingToDelete = errors.New("nothing to delete")
)

type Pair struct {
	Login       string
	Password    string
	Name        string
	UserID      int
	ID          int
	CreatedTime time.Time
}
