package main

import (
	"fmt"
	"os"
)

func printToken(content []byte, i int) (int, error) {
	switch content[i] {
	case '(':
		fmt.Println("LEFT_PAREN ( null")
	case ')':
		fmt.Println("RIGHT_PAREN ) null")
	case '{':
		fmt.Println("LEFT_BRACE { null")
	case '}':
		fmt.Println("RIGHT_BRACE } null")
	case '*':
		fmt.Println("STAR * null")
	case '.':
		fmt.Println("DOT . null")
	case ',':
		fmt.Println("COMMA , null")
	case '+':
		fmt.Println("PLUS + null")
	case '-':
		fmt.Println("MINUS - null")
	case ';':
		fmt.Println("SEMICOLON ; null")
	case '=':
		if i+1 < len(content) && content[i+1] == '=' {
			fmt.Println("EQUAL_EQUAL == null")
			i += 1
		} else {
			fmt.Println("EQUAL = null")
		}
	case '!':
		if i+1 < len(content) && content[i+1] == '=' {
			fmt.Println("BANG_EQUAL != null")
			i += 1
		} else {
			fmt.Println("BANG ! null")
		}
	case '<':
		if i+1 < len(content) && content[i+1] == '=' {
			fmt.Println("LESS_EQUAL <= null")
			i += 1
		} else {
			fmt.Println("LESS < null")
		}
	case '>':
		if i+1 < len(content) && content[i+1] == '=' {
			fmt.Println("GREATER_EQUAL >= null")
			i += 1
		} else {
			fmt.Println("GREATER > null")
		}
	case '/':
		if i+1 < len(content) && content[i+1] == '/' {
			for i < len(content) && content[i] != '\n' {
				i++
			}
			return i, nil
		} else {
			fmt.Println("SLASH / null")
		}
	default:
		fmt.Fprintf(os.Stderr, "[line 1] Error: Unexpected character: %c\n", content[i])
		return i, fmt.Errorf("bad char")
	}
	return i, nil
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
		for i := 0; i < len(fileContents); i++ {
			i, err = printToken(fileContents, i)
			if err != nil {
				defer os.Exit(65)
			}
		}
	}
	fmt.Println("EOF  null")
}
