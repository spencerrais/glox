package report

import "fmt"

var HadError bool

func LoxError(line int, message string) {
	LoxReport(line, "", message)
}

func LoxReport(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	HadError = true
}
