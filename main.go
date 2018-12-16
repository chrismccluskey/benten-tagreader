package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dhowden/tag"
	"gopkg.in/alessio/shellescape.v1"
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
			if file[len(file)-4:len(file)] == ".mp3" {

				fileContents, err := os.Open(file) // For read access.
				if err != nil {
					log.Print(err)
				} else {

					m, err := tag.ReadFrom(fileContents)
					if err != nil {
						log.Print(err)
					} else {

						//fmt.Println(file)
						for k, v := range m.Raw() {

							//fmt.Printf("%s: %s\n", k, v)

							if k == "TIT1" && v == "DELETE" {
								fmt.Println(shellescape.Quote(file))
							}
						}
						//fmt.Println()
					}
				}
			}
		}
	}

}
