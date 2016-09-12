package main

import (
	"bufio"
	"fmt"
	"path/filepath"
	"regexp"
	"strings"

	"os"
)

type Link struct {
	Description string
	URL         string
	Header1     string
}

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

	links := []Link{}

	for _, file := range fileList {
		fileHandle, _ := os.Open(file)
		defer fileHandle.Close()
		fileScanner := bufio.NewScanner(fileHandle)

		currentHeader := ""
		for fileScanner.Scan() {
			line := fileScanner.Text()

			regexHeader, _ := regexp.Compile("\\#\\s*(.*)")
			regexLink, _ := regexp.Compile("\\[(.*)\\]\\((.*)\\)")
			if regexHeader.MatchString(line) {
				currentHeader = regexHeader.FindStringSubmatch(line)[1]
			}
			if regexLink.MatchString(line) {
				find := regexLink.FindStringSubmatch(line)
				links = append(links, Link{find[1], find[2], currentHeader})
			}
		}
	}

	// output
	for _, link := range links {
		fmt.Println("Header:" + link.Header1)
		fmt.Println("Description:" + link.Description)
		fmt.Println("URL:" + link.URL)
	}

}
