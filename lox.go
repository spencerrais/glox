package main

import (
	"bufio"
	"fmt"
	"github.com/spencerrais/glox/report"
	"github.com/spencerrais/glox/scanner"
	"os"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		runFile(os.Args[1])
	} else {
		runPrompt()
	}
}

// run a lox file based on its path
func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	run(string(bytes))
	if report.HadError {
		os.Exit(65)
	}
}

// runs a REPL line by line
func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("> ")
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}
		run(line)
		fmt.Println("> ")
		report.HadError = false
	}
}

func run(source string) {
	scanner := scanner.NewScanner(source)
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}
