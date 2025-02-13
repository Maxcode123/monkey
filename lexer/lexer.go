package lexer

import (
	"monkey/token"
)

type Lexer struct {
	input           string
	currentPosition int
	nextPosition    int
	currentChar     byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.nextPosition >= len(l.input) {
		l.currentChar = 0
	} else {
		l.currentChar = l.input[l.nextPosition]
	}

	l.currentPosition = l.nextPosition
	l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.currentChar {
	case '=':
		if l.nextChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: "=="}
			l.readChar()
		} else {
			tok = newToken(token.ASSIGN, '=')
		}
	case ';':
		tok = newToken(token.SEMICOLON, ';')
	case '(':
		tok = newToken(token.LPAREN, '(')
	case ')':
		tok = newToken(token.RPAREN, ')')
	case ',':
		tok = newToken(token.COMMA, ',')
	case '+':
		tok = newToken(token.PLUS, '+')
	case '-':
		tok = newToken(token.MINUS, '-')
	case '!':
		if l.nextChar() == '=' {
			tok = token.Token{Type: token.NOT_EQ, Literal: "!="}
			l.readChar()
		} else {
			tok = newToken(token.BANG, '!')
		}
	case '*':
		tok = newToken(token.ASTERISK, '*')
	case '/':
		tok = newToken(token.SLASH, '/')
	case '<':
		tok = newToken(token.LT, '<')
	case '>':
		tok = newToken(token.GT, '>')
	case '{':
		tok = newToken(token.LBRACE, '{')
	case '}':
		tok = newToken(token.RBRACE, '}')
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(l.currentChar) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.currentChar) {
			tok.Literal = l.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readIdentifier() string {
	startPosition := l.currentPosition
	for isLetter(l.currentChar) {
		l.readChar()
	}
	return l.input[startPosition:l.currentPosition]
}

func (l *Lexer) readNumber() string {
	startPosition := l.currentPosition
	for isDigit(l.currentChar) {
		l.readChar()
	}
	return l.input[startPosition:l.currentPosition]
}

func (l *Lexer) skipWhitespace() {
	for l.currentChar == ' ' || l.currentChar == '\t' || l.currentChar == '\n' || l.currentChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) nextChar() byte {
	if l.nextPosition >= len(l.input) {
		return 0
	}

	return l.input[l.nextPosition]
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(char)}
}

func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

func isDigit(char byte) bool {
	return '0' <= char && char <= '9'
}
