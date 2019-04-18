package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	_fileGraphQLDeliver "github.com/chrismccluskey/benten-tagreader/file/delivery/graphql"
	_fileRepo "github.com/chrismccluskey/benten-tagreader/file/repository"
	_fileUsecase "github.com/chrismccluskey/benten-tagreader/file/usecase"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		fmt.Println("Running in DEBUG mode")
	}
}

func main() {
	// --- init
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	fr := _fileRepo.NewFilesystemFileRepo("")
	fu := _fileUsecase.NewFileUsecase(fr, timeoutContext)
	_fileGraphQLDeliver.NewFileGraphQLHandler(fu)

	root := "./"

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {

		if len(path) > 3 && path[len(path)-4:len(path)] == ".mp3" {
			fmt.Println(path)

			// get file information

			// could push to graphql
		}

		return nil

	})

	if err != nil {
		log.Fatal(err)
	}

}
