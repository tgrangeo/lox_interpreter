package main

import (
	"fmt"
	"os"
)

func Scanner(c byte) {
	switch c {
	case '(':
		fmt.Println("LEFT_PAREN ( null")
	case ')':
		fmt.Println("RIGHT_PAREN ) null")
	case '{':
		fmt.Println("LEFT_BRACE { null")
	case '}':
		fmt.Println("RIGHT_BRACE } null")
	}
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
			Scanner(c)
		}
	}
	fmt.Println("EOF  null")
}
