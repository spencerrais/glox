package token

type Token struct {
	Type    TokenType
	Lexeme  string
	Line    int
	Literal interface{}
}

func (t Token) String() string {
	return fmt.Sprintf("%v %v %v", t.Type, t.Lexeme, t.Literal)
}
