package main

import (
	"fmt"
	"log"
	"os"
	"flag"
	"path/filepath"

	"github.com/dhowden/tag"
	"gopkg.in/alessio/shellescape.v1"
)

func main() {
	var files []string

	verbose := flag.Bool("verbose", false, "increase verbosity")
	printFrames := flag.Bool("print-frames", false, "print all id3 frames in file")
	matchText := flag.String("match-text", "", "find text in frame (requires match-frame)")
	matchFrame := flag.String("match-frame", "", "find text in frame (requires match-frame)")

	flag.Parse()

	if *verbose {
		fmt.Printf("Searching for %s = %s\n\n", *matchFrame, *matchText)
	}

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

						if *printFrames {
							fmt.Println()
							fmt.Println(file)
						}
						for k, v := range m.Raw() {

							if *printFrames {
								fmt.Printf("%s: %s\n", k, v)
							} else {
								if k == *matchFrame && v == *matchText {
									fmt.Println(shellescape.Quote(file))
								}
							}

						}
					}
				}
			}
		}
	}

}
