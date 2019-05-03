package repository

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

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

func (fr *filesystemFileRepo) GetOne(ctx context.Context, path string, info os.FileInfo) (*models.File, error) {

	if !info.IsDir() {

		hasher := sha256.New()

		f, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		name := f.Name()

		//fi, err := f.Stat()
		//if err != nil {
		//	log.Fatal(err)
		//}

		size := info.Size()

		if _, err := io.Copy(hasher, f); err != nil {
			log.Fatal(err)
		}

		id := hex.EncodeToString(hasher.Sum(nil))

		file := &models.File{ID: id, Name: name, Size: size}

		return file, nil

	}

	return nil, errors.New("Cannot read file properties of directory")

}

func (fr *filesystemFileRepo) GetAll(ctx context.Context, root string) (*[]models.File, error) {

	var files []models.File

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if len(path) > 3 && path[len(path)-4:len(path)] == ".mp3" { //TODO: move filetype logic elsewhere
			fmt.Println(path)

			// get file information
			file, err := fr.GetOne(context.TODO(), path, info)
			if err != nil {
				panic(err)
			}
			fmt.Println(file)
			files = append(files, *file)

		}

		return nil

	})

	return &files, err

}
