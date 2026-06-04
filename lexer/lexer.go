package lexer

import (
	"log"
	"mili/token"
)

type LexerImpl interface {
	NextToken() token.Token
}

// lexer structure model
type Lexer struct {
	input        string // the source code input
	position     int    // current position in input (points to current char)
	readPosition int    // current reading position in input (after current char)
	ch           byte   // current char under examination
}

// Creates a new Lexer by taking an input (original source code)
func NewLexer(input string) LexerImpl {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

// read the current character of the input
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

// read identfier and return it
func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// read number and return it
func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

// skip white space
func (l *Lexer) skipWhiteSpace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\t' || l.ch == '\r' {
		l.readChar()
	}
}

// Get the current Token
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhiteSpace()

	log.Println(string(l.ch))
	switch l.ch {
	case '=':
		tok = newToken(token.ASSIGN, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok = newToken(token.EOF, l.ch)
	default:
		if isLetter(l.ch) {
			output := l.readIdentifier()
			return token.Token{Type: token.LookupIdent(output), Literal: output}
		} else if isDigit(l.ch) {
			output := l.readNumber()
			return token.Token{Type: token.INTEGER, Literal: output}
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}
