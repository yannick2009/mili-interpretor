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
	ASSIGN   TokenType = "=" // Assignation sign
	PLUS     TokenType = "+" // Plus (+) sign
	MINUS    TokenType = "-" // Subtraction (-) sign
	BANG     TokenType = "!" // Exclamation
	ASTERISk TokenType = "*" // Multiplication (*) sign
	SLASH    TokenType = "/" // Slash

	LT     TokenType = "<"
	GT     TokenType = ">"
	EQ     TokenType = "=="
	NOT_EQ TokenType = "!="

	// Delimiters
	COMMA     TokenType = "," // comma
	SEMICOLON TokenType = ";" // Semicolon
	LPAREN    TokenType = "(" // Left Parenthese
	RPAREN    TokenType = ")" // Right Parenthese
	LBRACE    TokenType = "{" // Left Brace
	RBRACE    TokenType = "}" // Right Brace

	// Keywords
	FUNCTION TokenType = "FUNCTION" // function Keyword
	LET      TokenType = "LET"      // variable keyword
	TRUE     TokenType = "TRUE"     // true keyword
	FALSE    TokenType = "FALSE"    // false keyword
	IF       TokenType = "IF"       // if condition keyword
	ELSE     TokenType = "ELSE"     // else keyword
	RETURN   TokenType = "RETURN"   // function return keyword
)

// search the right keyword for the input
func LookupIdent(input string) TokenType {
	if tokType, ok := keywords[input]; ok {
		return tokType
	}

	return IDENT
}

// keywords mapping
var keywords = map[string]TokenType{
	"let":    LET,
	"fn":     FUNCTION,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}
