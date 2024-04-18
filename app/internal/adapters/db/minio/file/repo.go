package filestorage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

const location = "us-east-1"

type config interface {
	MinioAccessKeyID() string
	MinioSecretAccessKey() string
	MinioEndpoint() string
}

type FileStorage struct {
	client *minio.Client
}

func NewFileStorage(cfg config) (*FileStorage, error) {
	minioClient, err := minio.New(cfg.MinioEndpoint(), &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.MinioAccessKeyID(), cfg.MinioSecretAccessKey(), ""),
		Secure: false,
	})

	if err != nil {
		return nil, err
	}

	return &FileStorage{client: minioClient}, nil
}
