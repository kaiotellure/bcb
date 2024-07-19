package main

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strings"
)

func dataUri(path, mime string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return "data:" + mime + ";base64," + base64.StdEncoding.EncodeToString(b)
}

func readFile(path string) string {
	b, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(b)
}

// put smaller strings replacements at the top to avoid issues
func remap(table map[string]string, target string) string {
	for key, value := range table {
		target = strings.ReplaceAll(target, key, value)
	}
	return target
}

func panicif(should bool, message string) {
	if should {
		panic(errors.New(message))
	}
}

func f(format string, v ...any) string {
	return fmt.Sprintf(format, v...)
}
