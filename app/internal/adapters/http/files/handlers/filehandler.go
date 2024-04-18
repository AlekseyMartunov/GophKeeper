package filehandlers

import (
	"GophKeeper/app/internal/entity/file"
	"context"
)

type logger interface {
	Info(s string)
	Error(e error)
}

type fileService interface {
	Save(ctx context.Context, f *file.File) error
	Get(ctx context.Context, userID int, fileName string) (*file.File, error)
	Delete(ctx context.Context, userID int, fileName string) error
	GetAll(ctx context.Context, userID int) ([]*file.File, error)
}

type FileHandler struct {
	service fileService
	log     logger
}

func NewFileHandler(l logger, s fileService) *FileHandler {
	return &FileHandler{
		log:     l,
		service: s,
	}
}
