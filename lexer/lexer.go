package lexer

import (
  "monkey/token"
)

type Lexer struct {
  input string
  currentPosition int
  nextPosition int
  currentChar byte
}

func New(input string) *Lexer {
  l := &Lexer{input: input}
  l.readChar()
  return l
}

func (l *Lexer) readChar() {
  if l.nextPosition >= len(l.input) {
    l.currentChar = 0
    return
  }

  l.currentChar = l.input[l.nextPosition]
  l.currentPosition = l.nextPosition
  l.nextPosition += 1
}

func (l *Lexer) NextToken() token.Token {
  var tok token.Token

  switch l.currentChar {
  case '=':
    tok = newToken(token.ASSIGN, '=')
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
  case '{':
    tok = newToken(token.LBRACE, '{')
  case '}':
    tok = newToken(token.RBRACE, '}')
  case 0:
    tok.Literal = ""
    tok.Type = token.EOF
  }

  l.readChar()
  return tok
}

func newToken(tokenType token.TokenType, char byte) token.Token {
  return token.Token{Type: tokenType, Literal: string(char)}
}
