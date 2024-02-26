package pairs

import (
	"errors"
	"time"

	"GophKeeper/internal/server/entity/users"
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
	User        users.User
	CreatedTime time.Time
	ID          int
}
