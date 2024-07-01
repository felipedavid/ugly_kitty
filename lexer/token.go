package lexer

import "fmt"

type TokenKind int

const (
	NONE TokenKind = iota
	LEFT_PAREN
	RIGHT_PAREN
	EOF
)

var tokenName = []string{
	NONE:        "NONE",
	LEFT_PAREN:  "LEFT_PAREN",
	RIGHT_PAREN: "RIGHT_PAREN",
	EOF:         "EOF",
}

type Token struct {
	Kind   TokenKind
	Lexeme string
	Value  any
	Line   int
}

func (t *Token) String() string {
	return fmt.Sprintf("[kind: '%s'] [lexeme: '%s'] [value: '%v']", tokenName[t.Kind], t.Lexeme, t.Value)
}
