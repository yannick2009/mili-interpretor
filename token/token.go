package token

type TokenType string

// Token structure model
type Token struct {
	Type    TokenType // Token Type
	Literal string    // initial value. keep it here because we don't want it to be lost
}

// declaration of the different type of tokens here
const (
	ILLEGAL TokenType = "ILLEGAL" // it means we dont know this token/charather
	EOF     TokenType = "EOF"     // end of file "EOF"

	// Identifier + literals
	IDENT   TokenType = "IDENT" // add, foobar, x, y...
	INTEGER TokenType = "INT"   // 12345678

	// Operators
	ASSIGN         TokenType = "=" // Assignation sign
	PLUS           TokenType = "+" // Plus (+) sign
	SUBTRACT       TokenType = "-" // Subtraction (-) sign
	MULTIPLICATION TokenType = "*" // Multiplication (*) sign
	DIVISION       TokenType = "/" // Division (/) sign

	// Delimiters
	COMMA     TokenType = "," // comma
	SEMICOLON TokenType = ";" // Semicolon
	LPAREN    TokenType = "(" // Left Parenthese
	RPAREN    TokenType = ")" // Right Parenthese
	LBRACE    TokenType = "{" // Left Brace
	RBRACE    TokenType = "}" // Right Brace

	// Keywords
	FUNCTION = "FUNCTION" // function
	LET      = "LET"      // let for variable declaration
)

// search the right keyword for the input
func LookupIdent(input string) TokenType {
	if tokType, ok := keywords[input]; ok {
		return tokType
	}

	return IDENT
}

var keywords = map[string]TokenType{
	"let": LET,
	"fn":  FUNCTION,
}
