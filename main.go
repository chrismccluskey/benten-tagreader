package main

import (
	"fmt"
	"log"
	"os"
	"github.com/dhowden/tag"
	"path/filepath"
)

func main() {
	var files []string

	root := "./"
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if len(file) > 3 {
			if file[len(file)-3:len(file)] == "mp3" {

				fileContents, err := os.Open(file) // For read access.
				if err != nil {
					log.Fatal(err)
				}

				m, err := tag.ReadFrom(fileContents)
				if err != nil {
					log.Fatal(err)
				}
				for k, v := range m.Raw() {
					if k == "TIT1" && v == "DELETE" {
						//fmt.Printf("%s: %s\n", k, v)
						fmt.Println(file)
					}
				}
			}
		}
	}

}
