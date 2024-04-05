package filestorage

import (
	"GophKeeper/internal/entity/file"
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"time"
)

func (fs *FileStorage) Get(ctx context.Context, bucketName, fileName string) (*file.File, error) {
	object, err := fs.client.GetObject(ctx, bucketName, fileName, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	defer object.Close()

	info, err := object.Stat()
	if err != nil {
		return nil, err
	}
	buff := make([]byte, info.Size)
	_, err = object.Read(buff)

	if err != nil && err != io.EOF {
		return nil, err
	}


	time, err := time.Parse(time.RFC3339, info.UserMetadata["Created_time"])
	if err != nil {
		return nil, err
	}

	file := file.File{
		Name:        fileName,
		Data:        buff,
		CreatedTime: time,
	}

	return &file, nil
}
