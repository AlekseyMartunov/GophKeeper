package filehandlers

import (
	"GophKeeper/internal/entity/file"
	"context"
)

type logger interface {
	Info(s string)
	Error(e error)
}

type fileService interface {
	Save(ctx context.Context, f *file.File) error
	Get(ctx context.Context, userID int, fileName string) (*file.File, error)
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
