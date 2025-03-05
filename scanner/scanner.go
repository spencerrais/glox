package scanner

import (
	"github.com/spencerrais/glox/report"
	. "github.com/spencerrais/glox/token"
)

var keywords = map[string]TokenType{
	"and":    AND,
	"class":  CLASS,
	"else":   ELSE,
	"false":  FALSE,
	"fun":    FUN,
	"for":    FOR,
	"if":     IF,
	"nil":    NIL,
	"or":     OR,
	"print":  PRINT,
	"return": RETURN,
	"super":  SUPER,
	"this":   THIS,
	"true":   TRUE,
	"var":    VAR,
	"while":  WHILE,
}

type Scanner struct {
	Source string
	Tokens []Token

	start   int
	current int
	line    int
}

func NewScanner(source string) *Scanner {
	return &Scanner{
		Source:  source,
		start:   0,
		current: 0,
		line:    1,
	}
}

func (s *Scanner) ScanTokens() []Token {
	for !s.isAtEnd() {
		s.start = s.current
		s.scanToken()
	}
	s.Tokens = append(s.Tokens, Token{
		Type:    EOF,
		Lexeme:  "",
		Line:    s.line,
		Literal: nil,
	})
	return s.Tokens
}

func (s *Scanner) isAtEnd() bool {
	return s.current >= len(s.Source)
}

func (s *Scanner) scanToken() {
	c := s.advance()
	switch c {
	case '(':
		s.addToken(LEFT_PAREN)
	case ')':
		s.addToken(RIGHT_PAREN)
	case '{':
		s.addToken(LEFT_BRACE)
	case '}':
		s.addToken(RIGHT_BRACE)
	case ',':
		s.addToken(COMMA)
	case '.':
		s.addToken(DOT)
	case '-':
		s.addToken(MINUS)
	case '+':
		s.addToken(PLUS)
	case ';':
		s.addToken(SEMICOLON)
	case '*':
		s.addToken(STAR)
	case '!':
		if s.match('=') {
			s.addToken(BANG_EQUAL)
		} else {
			s.addToken(BANG)
		}
	case '=':
		if s.match('=') {
			s.addToken(EQUAL_EQUAL)
		} else {
			s.addToken(EQUAL)
		}
	case '<':
		if s.match('=') {
			s.addToken(LESS_EQUAL)
		} else {
			s.addToken(LESS)
		}
	case '>':
		if s.match('=') {
			s.addToken(GREATER_EQUAL)
		} else {
			s.addToken(GREATER)
		}
	case '/':
		if s.match('/') {
			for s.peak() != '\n' && !s.isAtEnd() {
				s.advance()
			}
		} else {
			s.addToken(SLASH)
		}
	case ' ', '\r', '\t':
		return
	case '\n':
		s.line += 1
		return
	case '"':
		s.string()
	default:
		if isDigit(c) {
			s.number()
		} else if isAlpha(c) {
			s.identifier()
		} else {
			report.LoxError(s.line, "Unexpected character.")
		}
	}
}

func (s *Scanner) advance() byte {
	c := s.Source[s.current]
	s.current += 1
	return c
}

func (s *Scanner) addToken(tokenType TokenType) {
	s.addTokenWithLiteral(tokenType, nil)
}

func (s *Scanner) addTokenWithLiteral(tokenType TokenType, literal interface{}) {
	text := s.Source[s.start:s.current]
	s.Tokens = append(s.Tokens, Token{
		Type:    tokenType,
		Lexeme:  text,
		Line:    s.line,
		Literal: literal,
	})
}

func (s *Scanner) match(expected byte) bool {
	if s.isAtEnd() {
		return false
	}
	if s.current != int(expected) {
		return false
	}
	s.current += 1
	return true
}

func (s *Scanner) peak() byte {
	if s.isAtEnd() {
		return 0
	}
	return s.Source[s.current]
}

func (s *Scanner) peakNext() byte {
	if s.current+1 >= len(s.Source) {
		return 0
	}
	return s.Source[s.current+1]
}

func (s *Scanner) string() {
	for s.peak() != '"' && !s.isAtEnd() {
		if s.peak() == '\n' {
			s.line += 1
		}
		s.advance()
	}
	if s.isAtEnd() {
		report.LoxError(s.line, "Unterminated string.")
		return
	}
	s.advance()
	value := s.Source[s.start+1 : s.current-1]
	s.addTokenWithLiteral(STRING, value)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func (s *Scanner) number() {
	for isDigit(s.peak()) {
		s.advance()
	}

	if s.peak() == '.' && isDigit(s.peakNext()) {
		s.advance()
		for isDigit(s.peak()) {
			s.advance()
		}
	}
}

func (s *Scanner) identifier() {
	for isAlphaNumeric(s.peak()) {
		s.advance()
	}
	text := s.Source[s.start:s.current]
	tokenType, ok := keywords[text]
	if !ok {
		tokenType = IDENTIFIER
	}
	s.addToken(tokenType)
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || c == '_'
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}
