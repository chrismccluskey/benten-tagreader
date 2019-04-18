package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"
	"os"

	"github.com/chrismccluskey/benten-tagreader/file"
	"github.com/chrismccluskey/benten-tagreader/models"
)

type filesystemFileRepo struct {
	FS string
}

func NewFilesystemFileRepo(fs string) file.Repository {
	return &filesystemFileRepo{
		FS: fs,
	}
}

func (fr *filesystemFileRepo) GetOne(ctx context.Context, path string) (*models.File, error) {

	hasher := sha256.New()

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	name := f.Name()

	fi, err := f.Stat()
	if err != nil {
		log.Fatal(err)
	}

	size := fi.Size()

	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}

	id := hex.EncodeToString(hasher.Sum(nil))

	file := &models.File{ID: id, Name: name, Size: size}

	return file, nil

}
