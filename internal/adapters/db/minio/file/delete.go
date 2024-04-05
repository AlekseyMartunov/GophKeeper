package filestorage

import (
	"context"
	"github.com/minio/minio-go/v7"
)

func (fs *FileStorage) Delete(ctx context.Context, bucketName, fileName string) error {
	err := fs.client.RemoveObject(ctx, bucketName, fileName, minio.RemoveObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
