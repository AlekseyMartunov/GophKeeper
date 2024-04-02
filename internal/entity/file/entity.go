package file

import "time"

type File struct {
	Name        string
	Data        []byte
	CreatedTime time.Time
	UserID      int
	ID          int
}
