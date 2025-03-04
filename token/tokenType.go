package token

import (
	"strconv"
)

// This entire file is just to implement a Java-like enum for TokenType
type TokenType int

const (
	// Single character tokens
	LEFT_PAREN  TokenType = iota // 0
	RIGHT_PAREN                  // 1
	LEFT_BRACE                   // 2
	RIGHT_BRACE                  // 3
	COMMA                        // 4
	DOT                          // 5
	MINUS                        // 6
	PLUS                         // 7
	SEMICOLON                    // 8
	SLASH                        // 9
	STAR                         // 10

	// One or two character tokens
	BANG          // 11
	BANG_EQUAL    // 12
	EQUAL         // 13
	EQUAL_EQUAL   // 14
	GREATER       // 15
	GREATER_EQUAL // 16
	LESS          // 17
	LESS_EQUAL    // 18

	// Literals
	IDENTIFIER // 19
	STRING     // 20
	NUMBER     // 21

	// Keywords
	AND    // 22
	CLASS  // 23
	ELSE   // 24
	FALSE  // 25
	FUN    // 26
	FOR    // 27
	IF     // 28
	NIL    // 29
	OR     // 30
	PRINT  // 31
	RETURN // 32
	SUPER  // 33
	THIS   // 34
	TRUE   // 35
	VAR    // 36
	WHILE  // 37

	EOF // 38
)

const _tokenTypeName = "LEFT_PARENRIGHT_PARENLEFT_BRACERIGHT_BRACECOMMADOTMINUSPLUSSEMICOLONSLASHSTARBANGBANG_EQUALEQUALEQUAL_EQUALGREATERGREATER_EQUALLESSLESS_EQUALIDENTIFIERSTRINGNUMBERANDCLASSELSEFALSEFUNFORIFNILORPRINTRETURNSUPERTHISTRUEVARWHILEEOF"

var _tokenTypeIndex = [...]uint8{
	0, 10, 21, 31, 42, 47, 50, 55, 59, 68,
	73, 77, 81, 91, 96, 107, 114, 127, 131, 141,
	151, 157, 163, 166, 171, 175, 180, 183, 186, 188,
	191, 193, 198, 204, 209, 213, 217, 220, 225, 228,
}

// Helper function to convert TokenType to string for debugging
func (t TokenType) String() string {
	if t < 0 || int(t) >= len(_tokenTypeIndex)-1 {
		return "TokenType(" + strconv.FormatInt(int64(t), 10) + ")"
	}
	return _tokenTypeName[_tokenTypeIndex[t]:_tokenTypeIndex[t+1]]
}
