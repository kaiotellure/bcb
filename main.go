package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	fmt.Println("reading file: ", flag.Arg(1))
	content := readFile(flag.Arg(1))

	switch flag.Arg(0) {
	case "render":
		interpreter := NewHTMLInterpreter()
		NewParser(content, interpreter).Process()

		os.WriteFile("out.html", []byte(interpreter.GetDocument()), 0644)
		fmt.Println("render completed: out.html")
	case "debug":
		interpreter := NewHTMLInterpreter()
		NewParser(content, interpreter).Process()

		os.WriteFile("debug.html", []byte(interpreter.GetDocument()), 0644)
		fmt.Println("render completed: debug.html")
	default:
		panic("no command provided.")
	}
}
