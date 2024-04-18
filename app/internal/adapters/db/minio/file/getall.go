package filestorage

import (
	"GophKeeper/app/internal/entity/file"
	"context"
	"time"
)

func (fs *FileStorage) GetAll(ctx context.Context, bucketName string) ([]*file.File, error) {
	var files []*minio.Object

	for obj := range fs.client.ListObjects(ctx, bucketName, minio.ListObjectsOptions{WithMetadata: true}) {
		if obj.Err != nil {
			return nil, obj.Err
		}
		object, err := fs.client.GetObject(ctx, bucketName, obj.Key, minio.GetObjectOptions{})
		if err != nil {
			return nil, err
		}

		files = append(files, object)
	}

	var resultFiles []*file.File

	for _, f := range files {
		info, err := f.Stat()
		if err != nil {
			return nil, err
		}

		time, err := time.Parse(time.RFC3339, info.UserMetadata["Created_time"])
		if err != nil {
			return nil, err
		}

		res := file.File{
			Name:        info.Key,
			CreatedTime: time,
		}

		resultFiles = append(resultFiles, &res)
	}

	return resultFiles, nil

}
