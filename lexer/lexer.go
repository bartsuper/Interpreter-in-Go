// The lexer turns input into tokens
/* 
Point:
    I've change the newToken() to take in a string instead of a byte.
    That way, I can use newToken() for every switch case in NextToken.
    I'm not sure if there will be any performance difference, so I'm putting this note right here.
    TODO: Benchmark the performance difference between this approach and the approach in the book.
*/
package lexer

import (
    "go-interpreter/token"
)

type Lexer struct {
    input string        
    position int      // Current position in input (points to current char), corresponds to the "ch" byte
    readPosition int  // Current reading position in input (after current char), basically the "peeking" the next char
    ch byte           // Current char under examination
}

func (l *Lexer) NextToken() token.Token {
    var tok token.Token

    l.skipWhitespace()

    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            // The token is "=="
            ch := l.ch
            l.readChar()  // ch would be '=' and l.ch would also '='
            literal := string(ch) + string(l.ch)
            tok = newToken(token.EQUAL, literal)
        } else {
            tok = newToken(token.ASSIGN, string(l.ch))
        }
    case '+':
        tok = newToken(token.PLUS, string(l.ch))
    case '-':
        tok = newToken(token.MINUS, string(l.ch))
    case '*':
        tok = newToken(token.ASTERISK, string(l.ch))
    case '/':
        tok = newToken(token.SLASH, string(l.ch))
    case '!':
        if l.peekChar() == '=' {
            // The token is "!="
            ch := l.ch
            l.readChar()
            literal := string(ch) + string(l.ch)
            tok = newToken(token.NOT_EQUAL, literal)
        } else {
            tok = newToken(token.BANG, string(l.ch))
        }
    case '<':
        tok = newToken(token.LESS_THAN, string(l.ch))
    case '>':
        tok = newToken(token.GREATER_THAN, string(l.ch))
    case ',':
        tok = newToken(token.COMMA, string(l.ch))
    case ';':
        tok = newToken(token.SEMICOLON, string(l.ch))
    case '(':
        tok = newToken(token.L_PAREN, string(l.ch))
    case ')':
        tok = newToken(token.R_PAREN, string(l.ch))
    case '{':
        tok = newToken(token.L_BRACE, string(l.ch))
    case '}':
        tok = newToken(token.R_BRACE, string(l.ch))
    case 0:
        tok = newToken(token.EOF, "")
    default:
        // If the current character is not a symbol but is a letter, read the whole string and return the identifier. Else it's an illegal character
        if isLetter(l.ch) {
            tok.Literal = l.readIdentifier()
            tok.Type = token.LookupIdentifier(tok.Literal)
            return tok
        } else if isDigit(l.ch) {
            tok.Type = token.INT
            tok.Literal = l.readNumber()
            return tok
        } else {
            tok = newToken(token.ILLEGAL, string(l.ch))
        } 
    }

    // Go to the next character
    l.readChar()
    return tok
}

func (l *Lexer) readChar() {
    if l.readPosition >= len(l.input) {
        // If the next position is out of bounds (reached the end of input)
        l.ch = 0 // 0 is the ASCII character for NUL
    } else {
        l.ch = l.input[l.readPosition]
    }
    l.position = l.readPosition
    l.readPosition += 1
}

// Kinda like a constructor?
func newToken(tokenType token.TokenType, literal string) token.Token {
    return token.Token{Type: tokenType, Literal: literal}
}

func (l *Lexer) readIdentifier() string {
    position := l.position
    for isLetter(l.ch) {
        // Reads the string until it reads a non-letter character (uses l.readChar() to move the l.position to the last char of l.ch)
        l.readChar()
    }

    // Returns the whole string identifier
    return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
    position := l.position
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[position:l.position]
}

// Looks ahead to determine if the symbol has more than one character
// peekChar() only returns the next character, but doesn't change the positions. Unlike readChar(), which does
func (l *Lexer) peekChar() byte {
    if l.readPosition >= len(l.input) {
        // You're at the end of the input
        return 0
    } else {
        // Returns(peeks) the next character
        return l.input[l.readPosition]
    }
}

// Check if the current character is a letter (check if the ASCII value of ch is between [a, z] or [A, Z], or if ch is _)
func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// Check if the current character is a isDigit
func isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

// The lexer needs to skip all of the whitespace
func (l *Lexer) skipWhitespace() {
    for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
        l.readChar()
    }
}

// Creates a new lexer (if you can't tell)
func New(input string) *Lexer {
    l := &Lexer{input: input}
    l.readChar()  //  Makes sure that the lexer is initialized
    return l
}


