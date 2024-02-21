package data

import (
	"time"

	"GophKeeper/internal/server/entity/users"
)

type Pair struct {
	Login    string
	Password string

	User       users.User
	ID         int
	ExternalID string
}

type CreditCard struct {
	Number int
	Owner  string
	Date   time.Time
	CVV    int

	User       users.User
	ID         int
	ExternalID string
}

type BinaryData struct {
	Data []byte

	User       users.User
	ID         int
	ExternalID string
}
