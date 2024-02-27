package data

import (
	"time"

	"GophKeeper/internal/server/entity/users"
)

type CreditCard struct {
	Number int
	Owner  string
	Date   time.Time
	CVV    int

	User users.User
	ID   int
}

type BinaryData struct {
	Data []byte

	User users.User
	ID   int
}
