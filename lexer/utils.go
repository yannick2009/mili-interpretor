package lexer

import "mili/token"

// create a new token
func newToken(tokenType token.TokenType, ch byte) token.Token {
	var chStr string
	if ch == 0 {
		chStr = ""
	} else {
		chStr = string(ch)
	}

	return token.Token{Type: tokenType, Literal: chStr}
}

// tell us if the current character is a Letter
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// tell us if the current character is a digit
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
