package fileservice

import (
	"GophKeeper/app/internal/entity/file"
	"context"
	"fmt"
)

type fileStorage interface {
	Save(ctx context.Context, f *file.File) error
	Get(ctx context.Context, bucketName, fileName string) (*file.File, error)
	Delete(ctx context.Context, bucketName, fileName string) error
	GetAll(ctx context.Context, bucketName string) ([]*file.File, error)
}

type FileService struct {
	repo fileStorage
}

func NewFileService(s fileStorage) *FileService {
	return &FileService{repo: s}
}

func (fs *FileService) Save(ctx context.Context, f *file.File) error {
	f.BucketName = fmt.Sprintf("%dbucket", f.UserID)
	return fs.repo.Save(ctx, f)
}

func (fs *FileService) Get(ctx context.Context, userID int, fileName string) (*file.File, error) {
	bucketName := fmt.Sprintf("%dbucket", userID)
	return fs.repo.Get(ctx, bucketName, fileName)
}

func (fs *FileService) Delete(ctx context.Context, userID int, fileName string) error {
	bucketName := fmt.Sprintf("%dbucket", userID)
	return fs.repo.Delete(ctx, bucketName, fileName)
}

func (fs *FileService) GetAll(ctx context.Context, userID int) ([]*file.File, error) {
	bucketName := fmt.Sprintf("%dbucket", userID)
	return fs.repo.GetAll(ctx, bucketName)
}
