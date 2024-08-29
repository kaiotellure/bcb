package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	quick := flag.Bool("quick", false, "--quick limits content to 1000 chars, rendering fast. good for previewing style changes.")
	flag.Parse()

	fmt.Println("reading file: ", flag.Arg(1))
	content := readFile(flag.Arg(1))
	if *quick {
		fmt.Println("quick mode: enabled.")
		content = content[:1000]
	}

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
