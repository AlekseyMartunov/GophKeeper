package filestorage

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (fs *FileStorage) createBucket(ctx context.Context, bucketName string) error {
	err := fs.client.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := fs.client.BucketExists(ctx, bucketName)
		if errBucketExists != nil || !exists {
			return err
		}
	}

	return nil
}
