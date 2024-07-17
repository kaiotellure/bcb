package main

import (
	"flag"
	"os"
)

type Interpreter interface {
	FeedLine(line string)
}

func parse(content string, interpreter Interpreter) {
	var clip string
	var previous rune
	var insideQuote bool

	for i, char := range content {

		if isNewLine(char) {
			previous = char
			continue // skip new lines
		}

		if char == '.' && !insideQuote { 
			// prevent multiples lines with elipisis
			if len(content) > i+1 && rune(content[i+1]) == '.' || previous == '.' {
				previous = char
				clip += string(char)
				continue
			}
			interpreter.FeedLine(clip + ".")
			clip = ""
			previous = char
			continue
		} else if isSpace(char) {
			if isSpace(previous) {
				previous = char
				continue
			}
			// spaceCount++
		} else if char == '-' && (isSpace(rune(content[i+1])) || isSpace(previous)) {
			char = '—'
		} else if isQuote(char) {
			insideQuote = !insideQuote
		}

		if len(clip) == 0 && isSpace(char) {
			previous = char
			continue
		}

		clip += string(char)
		previous = char
	}
}

func main() {
	outputAsHTML := flag.Bool("html", false, "uses template.html to produce a standalone file.")
	flag.Parse()

	panicif(flag.Arg(0) == "", "please, provide a file to process")
	content := readFile(flag.Arg(0))

	if *outputAsHTML {
		interpreter := &ToHTML{}
		parse(content, interpreter)

		document := remap(map[string]string{
			"$$LOGO$$":  dataUri("./assets/images/logo.png", "image/jpg"),
			"$$LINES$$": interpreter.lines,
		}, readFile("./assets/template.html"))

		os.Stdout.WriteString(document)
		return
	}

	parse(content, &ToTXT{})
}
