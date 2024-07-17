package main

import (
	"os"
	"strings"
)

type ToTXT struct {}

func isChapter(line string) bool {
	if isUpper(line[0]) {
		for i, char := range line {
			if char == ' ' {
				if isUpper(line[i+1]) {
					return true
				}
				return false
			}
		}
	}
	return false
}

var whitelist []string = []string{
	"Robert",
}

func (i *ToTXT) FeedLine(line string) {
	if isChapter(line) {
		for _, v := range whitelist {
			if strings.HasPrefix(line, v) {
				os.Stdout.WriteString(line + "\n")
			}
		}
	}
}
