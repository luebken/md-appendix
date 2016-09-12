package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"os"

	_ "github.com/russross/blackfriday"
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
	}
}
