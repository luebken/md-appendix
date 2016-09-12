package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"os"

	"github.com/russross/blackfriday"
)

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Printf("usage: md-appendix <dir>")
		os.Exit(0)
	}
	searchDir := argsWithoutProg[0]

	fileList := []string{}

	filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		if strings.HasSuffix(f.Name(), ".md") {
			fileList = append(fileList, path)
		}
		return nil
	})

	for _, file := range fileList {
		fmt.Println(file)
		file, err := ioutil.ReadFile(file) // For read access.
		if err != nil {
			log.Fatal(err)
		}
		output := blackfriday.MarkdownCommon(file)

		fmt.Println(string(output))

	}
}
