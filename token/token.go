package token


type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	// Identifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
    MINUS = "-"
    BANG = "!"
    ASTERISK = "*"
    SLASH = "/"

    LESS_THAN = "<"
    GREATER_THAN = ">"
    
    EQUAL = "=="
    NOT_EQUAL = "!="

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	L_PAREN = "("
	R_PAREN = ")"
	L_BRACE = "{"
	R_BRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET = "LET"
    TRUE = "TRUE"
    FALSE = "FALSE"
    IF = "IF"
    ELSE = "ELSE"
    RETURN = "RETURN"
)

var keywords = map[string]TokenType {
    "fn": FUNCTION,
    "let": LET,
    "true": TRUE,
    "false": FALSE,
    "if": IF,
    "else": ELSE,
    "return": RETURN,
}

// Check if identifier is a keyword
func LookupIdentifier(identifier string) TokenType {
    if tok, ok := keywords[identifier]; ok {
        return tok
    }
    return IDENTIFIER
}
