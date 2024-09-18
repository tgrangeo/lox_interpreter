package main

import (
	"fmt"
	"os"
)

var tokens = map[rune]string{
	'(':  "LEFT_PAREN ( null",
	')':  "RIGHT_PAREN ) null",
	'{':  "LEFT_BRACE { null",
	'}':  "RIGHT_BRACE } null",
	'\n': "EOF  null",
	'*':  "STAR * null",
	'.':  "DOT . null",
	',':  "COMMA , null",
	'+':  "PLUS + null",
	'-':  "MINUS - null",
	';':  "SEMICOLON ; null",
}

func main() {
	if len(os.Args) < 3 {
		fmt.Fprintln(os.Stderr, "Usage: ./your_program.sh tokenize <filename>")
		os.Exit(1)
	}
	command := os.Args[1]
	if command != "tokenize" {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", command)
		os.Exit(1)
	}
	filename := os.Args[2]
	fileContents, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading file: %v\n", err)
		os.Exit(1)
	}

	if len(fileContents) > 0 {
		for _, c := range fileContents {
			fmt.Println(tokens[rune(c)])
		}
	}
	fmt.Println("EOF  null")
}
