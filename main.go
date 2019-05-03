package main

import (
	"context"
	"fmt"
	"log"
	"os"
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

	root := os.Args[1]

	files, err := fr.GetAll(context.TODO(), root)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(files)

}
