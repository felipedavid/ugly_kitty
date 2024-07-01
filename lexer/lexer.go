package lexer

import "fmt"

type Lexer struct {
	token  Token
	source []byte

	start   int
	current int
	line    int
}

func New(s []byte) *Lexer {
	return &Lexer{
		source: s,
		line:   1,
	}
}

func (l *Lexer) Next() {
	ch := l.advance()

	token := &l.token
	switch ch {
	case '(':
		token.Kind = LEFT_PAREN
		token.Lexeme = string(l.source[l.start:l.current])
	case ')':
		token.Kind = RIGHT_PAREN
		token.Lexeme = string(l.source[l.start:l.current])
	case 0:
		token.Kind = EOF
	default:
		fmt.Printf("Unknown character '%c'\n", ch)
	}
}

func (l *Lexer) advance() byte {
	if l.isAtEnd() {
		return 0
	}

	ch := l.source[l.current]
	l.current++
	return ch
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}

func (l *Lexer) ParseTokens() []Token {
	var tokens []Token
	for !l.isAtEnd() {
		l.start = l.current
		l.Next()

		tokens = append(tokens, l.token)
	}

	tokens = append(tokens, Token{Kind: EOF})
	return tokens
}
