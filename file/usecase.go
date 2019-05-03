package file

import (
	"context"
	"os"

	"github.com/chrismccluskey/benten-tagreader/models"
)

type Usecase interface {
	GetOne(ctx context.Context, path string, info os.FileInfo) (*models.File, error)
	GetAll(ctx context.Context, root string) (*[]models.File, error)
}
