package file

import (
	"context"

	"github.com/chrismccluskey/benten-tagreader/models"
)

type Usecase interface {
	GetOne(ctx context.Context, path string) (*models.File, error)
}
