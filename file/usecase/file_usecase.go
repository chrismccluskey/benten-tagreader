package usecase

import (
	"context"
	"os"
	"time"

	"github.com/chrismccluskey/benten-tagreader/file"
	"github.com/chrismccluskey/benten-tagreader/models"
)

type fileUsecase struct {
	fileRepo       file.Repository
	contextTimeout time.Duration
}

func NewFileUsecase(f file.Repository, timeout time.Duration) file.Usecase {
	return &fileUsecase{
		fileRepo:       f,
		contextTimeout: timeout,
	}
}

func (f *fileUsecase) GetOne(c context.Context, path string, info os.FileInfo) (*models.File, error) {

	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	file, err := f.fileRepo.GetOne(ctx, path, info)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (f *fileUsecase) GetAll(c context.Context, root string) (*[]models.File, error) {

	ctx, cancel := context.WithTimeout(c, f.contextTimeout)
	defer cancel()

	file, err := f.fileRepo.GetAll(ctx, root)
	if err != nil {
		return nil, err
	}

	return file, nil
}
