package lexer

import (
	"fmt"
)

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

func (l *Lexer) Next() error {
	ch := l.advance()

	token := &l.token
	switch ch {
	case '(':
		token.Kind = LEFT_PAREN
	case ')':
		token.Kind = RIGHT_PAREN
	case '{':
		token.Kind = LEFT_BRACE
	case '}':
		token.Kind = RIGHT_BRACE
	case ',':
		token.Kind = COMMA
	case '.':
		token.Kind = DOT
	case '-':
		token.Kind = MINUS
	case '+':
		token.Kind = PLUS
	case ';':
		token.Kind = SEMICOLON
	case '*':
		token.Kind = PLUS
	case '!':
		if l.match('=') {
			token.Kind = BANG_EQUAL
		} else {
			token.Kind = BANG
		}
	case '=':
		if l.match('=') {
			token.Kind = EQUAL_EQUAL
		} else {
			token.Kind = EQUAL
		}
	case '<':
		if l.match('=') {
			token.Kind = LESS_EQUAL
		} else {
			token.Kind = LESS
		}
	case '>':
		if l.match('=') {
			token.Kind = GREATER_EQUAL
		} else {
			token.Kind = GREATER
		}
	case 0:
		token.Kind = EOF
	default:
		return fmt.Errorf("unexpected character '%c'", ch)
	}

	token.Lexeme = string(l.source[l.start:l.current])

	return nil
}

// matches the current character to ch, if matches eats one ch and returns true
func (l *Lexer) match(ch byte) bool {
	if l.peek() == ch {
		l.current++
		return true
	}

	return false
}

func (l *Lexer) peek() byte {
	if l.isAtEnd() {
		return 0
	}

	return l.source[l.current]
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
		err := l.Next()
		if err != nil {
			fmt.Println(err.Error())
		}

		tokens = append(tokens, l.token)
	}

	tokens = append(tokens, Token{Kind: EOF})
	return tokens
}
