package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var hadError bool

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

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return
	}
	run(string(bytes))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("> ")
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
	}
}

func run(source string) {
	scanner := NewScanner(source) // to implement still
	tokens := scanner.ScanTokens()

	for _, token := range tokens {
		fmt.Println(token)
	}
}

func error(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	hadError = true
}
