package main

import (
	"fmt"
	"os"
)

var line = 1

func printToken(content []byte, i int) (int, error) {
	switch content[i] {
	case ' ', '\t':
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
			return i - 1, nil
		} else {
			fmt.Println("SLASH / null")
		}
	case '"':
		res := ""
		for i += 1; i < len(content) && content[i] != '"'; i++ {
			res += string(content[i])
		}
		if i == len(content) {
			fmt.Fprintf(os.Stderr, "[line %d] Error: Unterminated string.\n", line)
			return i, fmt.Errorf("bad char")
		} else {
			fmt.Printf("STRING \"%s\" %s\n", res, res)
		}
	case '\n':
		line++
	default:
		fmt.Fprintf(os.Stderr, "[line %d] Error: Unexpected character: %c\n", line, content[i])
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
