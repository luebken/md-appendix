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
	Header2     string
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
		currentSubHeader := ""
		for fileScanner.Scan() {
			line := fileScanner.Text()

			regexSubHeader, _ := regexp.Compile("^\\#\\#\\s*(.*)")
			regexHeader, _ := regexp.Compile("^\\#\\s*(.*)")
			if regexSubHeader.MatchString(line) {
				currentSubHeader = regexSubHeader.FindStringSubmatch(line)[1]
			} else if regexHeader.MatchString(line) {
				currentHeader = regexHeader.FindStringSubmatch(line)[1]
				currentSubHeader = ""
			}

			regexLink, _ := regexp.Compile("\\[(.*)\\]\\((http.*)\\)")
			if regexLink.MatchString(line) {
				find := regexLink.FindStringSubmatch(line)
				links = append(links, Link{find[1], find[2], currentHeader, currentSubHeader})
			}
		}
	}

	// output
	sections := make(map[string][]Link)
	for _, link := range links {
		headerString := link.Header1 + "::" + link.Header2
		sections[headerString] = append(sections[headerString], link)
	}

	fmt.Println("------------------ >8 ------------------\n")
	fmt.Println("# Links:\n")
	for header, links := range sections {
		fmt.Println(header)
		for _, link := range links {
			fmt.Println("* [" + link.Description + "](" + link.URL + ")")

		}
	}
	fmt.Println("\n------------------ 8< ------------------")

}
