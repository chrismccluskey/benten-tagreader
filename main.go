package main

import (
	"fmt"
	"os"
	"github.com/mikkyang/id3-go"
	"path/filepath"
)

func main() {
	fmt.Println("Reading files in directory")
	var files []string
	var id3v2Frames = []string{"AENC","APIC","ASPI","COMM","COMR","ENCR","EQU2","ETCO","GEOB","GRID","LINK","MCDI","MLLT","OWNE","PRIV","PCNT","POPM","POSS","RBUF","RVA2","RVRB","SEEK","SIGN","SYLT","SYTC","TALB","TBPM","TCOM","TCON","TCOP","TDEN","TDLY","TDOR","TDRC","TDRL","TDTG","TENC","TEXT","TFLT","TIPL","TIT1","TIT2","TIT3","TKEY","TLAN","TLEN","TMCL","TMED","TMOO","TOAL","TOFN","TOLY","TOPE","TOWN","TPE1","TPE2","TPE3","TPE4","TPOS","TPRO","TPUB","TRCK","TRSN","TRSO","TSOA","TSOP","TSOT","TSRC","TSSE","TSST","TXXX","UFID","USER","USLT","WCOM","WCOP","WOAF","WOAR","WOAS","WORS","WPAY","WPUB","WXXX"}

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
				fmt.Println(file)
				mp3File, err := id3.Open(file)
				defer mp3File.Close()

				if err != nil {
					fmt.Println(err)
				}

				for _, frameName := range id3v2Frames {
					id3Frame := mp3File.Frame(frameName)
					if id3Frame != nil {
						fmt.Printf("%s: %s\n", frameName, id3Frame.String())
					} else {
						fmt.Printf("%s: (empty)\n", frameName)
					}
				}
				fmt.Println()
			}
		}
	}

}
