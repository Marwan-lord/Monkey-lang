package ast

import (
	"monkeylang/token"
	"testing"
)


func TestString(t *testing.T) {
    program := &Program{
        Statements: []Statement {
            &LetStatement{
                Token: token.Token{Type: token.Let, Literal: "let"},
                Name: &Identifier {
                    Token: token.Token{Type: token.Ident, Literal: "myVar"},
                    Value: "myVar",
                },
                Value: &Identifier{
                    Token: token.Token{Type: token.Ident, Literal: "anotherVar"},
                    Value: "anotherVar",
                },
            },
        },
    }
    if program.String() != "let myVar = anotherVar;" {
        t.Errorf("Program.String is wrong; got=%q", program.String())
    }
}
