package filestorage

import (
	"GophKeeper/internal/entity/file"
	"bytes"
	"context"
	"github.com/minio/minio-go/v7"
	"time"
)

func (fs *FileStorage) Save(ctx context.Context, f *file.File) error {
	err := fs.createBucket(ctx, f.BucketName)
	if err != nil {
		return err
	}

	_, err = fs.client.PutObject(
		ctx,
		f.BucketName,
		f.Name,
		bytes.NewReader(f.Data),
		int64(len(f.Data)),
		minio.PutObjectOptions{
			UserMetadata: map[string]string{
				"Created_time": f.CreatedTime.Format(time.RFC3339),
			},
		},
	)

	if err != nil {
		return err
	}

	return nil
}
