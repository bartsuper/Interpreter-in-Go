// The lexer turns input into tokens
package lexer

import (
	"testing"
    "go-interpreter/token"
)



func TestNextToken(t *testing.T) {
    input := `
        let five = 5;
        let ten = 10;
        let add = fn(x, y) {
            x + y;
        };
        
        let result = add(five, ten);
        !-/*5;
        5 < 10 > 5;

        if (5 < 10) {
            return true;
        } else {
            return false;
        }

        10 == 10;
        10 != 9;
    `

    // Tests every token case.
    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
    }{
        {token.LET, "let"},
        {token.IDENTIFIER, "five"},
        {token.ASSIGN, "="}, 
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENTIFIER, "ten"},
        {token.ASSIGN, "="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENTIFIER, "add"},
        {token.ASSIGN, "="},
        {token.FUNCTION, "fn"},
        {token.L_PAREN, "("},
        {token.IDENTIFIER, "x"},
        {token.COMMA, ","},
        {token.IDENTIFIER, "y"},
        {token.R_PAREN, ")"},
        {token.L_BRACE, "{"},
        {token.IDENTIFIER, "x"},
        {token.PLUS, "+"},
        {token.IDENTIFIER, "y"},
        {token.SEMICOLON, ";"},
        {token.R_BRACE, "}"},
        {token.SEMICOLON, ";"},
        {token.LET, "let"},
        {token.IDENTIFIER, "result"},
        {token.ASSIGN, "="},
        {token.IDENTIFIER, "add"},
        {token.L_PAREN, "("},
        {token.IDENTIFIER, "five"},
        {token.COMMA, ","},
        {token.IDENTIFIER, "ten"},
        {token.R_PAREN, ")"},
        {token.SEMICOLON, ";"},
        {token.BANG, "!"},
        {token.MINUS, "-"},
        {token.SLASH, "/"},
        {token.ASTERISK, "*"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.INT, "5"},
        {token.LESS_THAN, "<"},
        {token.INT, "10"},
        {token.GREATER_THAN, ">"},
        {token.INT, "5"},
        {token.SEMICOLON, ";"},
        {token.IF, "if"},
        {token.L_PAREN, "("},
        {token.INT, "5"},
        {token.LESS_THAN, "<"},
        {token.INT, "10"},
        {token.R_PAREN, ")"},
        {token.L_BRACE, "{"},
        {token.RETURN, "return"},
        {token.TRUE, "true"},
        {token.SEMICOLON, ";"},
        {token.R_BRACE, "}"},
        {token.ELSE, "else"},
        {token.L_BRACE, "{"},
        {token.RETURN, "return"},
        {token.FALSE, "false"},
        {token.SEMICOLON, ";"},
        {token.R_BRACE, "}"},
        {token.INT, "10"},
        {token.EQUAL, "=="},
        {token.INT, "10"},
        {token.SEMICOLON, ";"},
        {token.INT, "10"},
        {token.NOT_EQUAL, "!="},
        {token.INT, "9"},
        {token.SEMICOLON, ";"},
        {token.EOF, ""},
    }

    l := New(input)
    
    // tt as in token test, I guess
    for i, tt := range tests {
        tok := l.NextToken()        

        if tok.Type != tt.expectedType {
            t.Fatalf("tests[%d] - tokentype wrong, expected=%q, got=%q", i, tt.expectedType, tok.Type) // Type is of type TokenType, which is of type string
        }

        if tok.Literal!= tt.expectedLiteral {
            t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
        }
    }

}

