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

		err := os.WriteFile("out.html", []byte(interpreter.GetDocument()), 0766)
		panicif(err != nil, "could not write to file.")

		fmt.Println("render completed: out.html")
	default:
		panic("no command provided.")
	}
}
