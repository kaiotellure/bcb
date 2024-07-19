package main

import (
	"flag"
	"os"
)

func main() {
	flag.Parse()
	content := readFile(flag.Arg(1))

	switch flag.Arg(0) {
	case "render":
		interpreter := ToHTML{}
		NewParser(content, &interpreter).Process()
		os.WriteFile("out.html", []byte(interpreter.GetDocument()), 0644)
	default:
		panic("no command provided.")
	}
}
