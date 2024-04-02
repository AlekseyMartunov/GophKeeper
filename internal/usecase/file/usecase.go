package fileservice

import (
	"GophKeeper/internal/entity/file"
	"context"
	"strconv"
)

type fileStorage interface {
	Save(ctx context.Context, f *file.File) error
	Get(ctx context.Context, bucketName, fileName string) (*file.File, error)
}

type FileService struct {
	repo fileStorage
}

func NewFileService(s fileStorage) *FileService {
	return &FileService{repo: s}
}

func (fs *FileService) Save(ctx context.Context, f *file.File) error {
	return fs.repo.Save(ctx, f)
}

func (fs *FileService) Get(ctx context.Context, userID int, fileName string) (*file.File, error) {
	return fs.repo.Get(ctx, strconv.Itoa(userID), fileName)
}
