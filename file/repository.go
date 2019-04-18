package file

import (
	"context"

	"github.com/chrismccluskey/benten-tagreader/models"
)

type Repository interface {
	GetOne(ctx context.Context, id string) (*models.File, error)
}
