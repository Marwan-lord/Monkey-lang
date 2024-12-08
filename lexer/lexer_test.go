package lexer 

import (
    "testing"
    "monkeylang/token"
)

func TestNextToken(t *testing.T) {
    input := `let five = 5;
    let ten = 10;
    let add = fn(x, y) {
        x + y;
    };

    let result = add(five, ten);
    !-/*5;
    5 < 10 > 5;

    if (5 < 10)  {
        return true;
    } else {
        return false;
    }

    10 == 10;
    10 != 9;
    `


    tests := []struct {
        expectedType token.TokenType
        expectedLiteral string
    } {
        {token.Let, "let"},
        {token.Ident, "five"},
        {token.Assign, "="},
        {token.Int, "5"},
        {token.SemiColon, ";"},
        {token.Let, "let"},
        {token.Ident, "ten"},
        {token.Assign , "="},
        {token.Int, "10"},
        {token.SemiColon, ";"},
        {token.Let, "let"},
        {token.Ident, "add"},
        {token.Assign, "="},
        {token.Function, "fn"},
        {token.LParen, "("},
        {token.Ident, "x"},
        {token.Comma, ","},
        {token.Ident, "y"},
        {token.RParen, ")"},
        {token.LBrace, "{"},
        {token.Ident, "x"},
        {token.Plus, "+"},
        {token.Ident, "y"},
        {token.SemiColon, ";"},
        {token.RBrace, "}"},
        {token.SemiColon, ";"},
        {token.Let, "let"},
        {token.Ident, "result"},
        {token.Assign, "="},
        {token.Ident, "add"},
        {token.LParen, "("},
        {token.Ident, "five"},
        {token.Comma, ","},
        {token.Ident, "ten"},
        {token.RParen, ")"},
        {token.SemiColon, ";"},
        {token.Bang, "!"},
        {token.Minus, "-"},
        {token.Slash, "/"},
        {token.Asterisk, "*"},
        {token.Int, "5"},
        {token.SemiColon, ";"},
        {token.Int, "5"},
        {token.LT, "<"},
        {token.Int, "10"},
        {token.GT, ">"},
        {token.Int, "5"},
        {token.SemiColon, ";"},
        {token.If, "if"},
        {token.LParen, "("},
        {token.Int, "5"},
        {token.LT, "<"},
        {token.Int, "10"},
        {token.RParen, ")"},
        {token.LBrace, "{"},
        {token.Return, "return"},
        {token.True, "true"},
        {token.SemiColon, ";"},
        {token.RBrace, "}"},
        {token.Else, "else"},
        {token.LBrace, "{"},
        {token.Return, "return"},
        {token.False, "false"},
        {token.SemiColon, ";"},
        {token.RBrace, "}"},
        {token.Int, "10"},
        {token.EqualTo, "=="},
        {token.Int, "10"},
        {token.SemiColon, ";"},
        {token.Int, "10"},
        {token.NotEqualTo, "!="},
        {token.Int, "9"},
        {token.SemiColon, ";"},
        {token.EOF, ""},
    }

    l := New(input)
    
    for i ,tt := range tests {
        tok := l.NextToken() 

        if tok.Type != tt.expectedType {
            t.Fatalf("Error: t[%d] token type wrong expected: %q. got: %q ",i, tt.expectedType, tok.Type)
        }

        if tok.Literal!= tt.expectedLiteral {
            t.Fatalf("Error: t[%d] Literal type wrong expected: %q. got: %q ",i, tt.expectedType, tok.Type)
        }

    }
}
