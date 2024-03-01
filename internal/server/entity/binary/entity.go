package binarydata

import (
	"time"

	"GophKeeper/internal/server/entity/users"
)

type BinaryData struct {
	Data []byte

	CreatedTime time.Time
	User        users.User
	ID          int
}
