package binarydata

import (
	"GophKeeper/internal/entity/users"
	"time"
)

type BinaryData struct {
	Data []byte

	CreatedTime time.Time
	User        users.User
	ID          int
}
