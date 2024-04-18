package file

import "time"

type File struct {
	Name        string
	BucketName  string
	Data        []byte
	CreatedTime time.Time
	UserID      int
}
