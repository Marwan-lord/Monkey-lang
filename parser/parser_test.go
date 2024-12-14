package parser

import (
    "testing"
    "monkeylang/ast"
    "monkeylang/lexer"
)

func TestReturnStatements(t *testing.T) {
    input := `
    return 5;
    return 10;
    return 23234;
    `

    l := lexer.New(input)
    p := New(l)

    prog := p.ParseProgram()
    checkParseErrors(t, p)

    if len(prog.Statements) != 3 {
        t.Fatalf("Program does not have 3 statements. got %d", len(prog.Statements))
    }

    for _, stmt := range prog.Statements {
        returnStmt, ok := stmt.(*ast.ReturnStatement)
        if !ok {
            t.Errorf("statement is not a return statement got %q", stmt)
            continue
        }

        if returnStmt.TokenLiteral() != "return" {
            t.Errorf(
                "return statement Literal expected got %q",
                returnStmt.TokenLiteral())
            }
        }
    }

    func TestLetStatements(t *testing.T) {
        input := `
        let x =  5;
        let y = 10 ;
        let foobar = 100030404;
        `
        l := lexer.New(input)
        p := New(l)
        program := p.ParseProgram()
        checkParseErrors(t, p)
        if program == nil {

            t.Fatalf("Program Parser returned nil")
        }

        if len(program.Statements) != 3 {
            t.Fatalf("Program does not contain 3 statements got %d", 
            len(program.Statements))
        }

        tests := []struct {
            expectedIdent string
        } {
            {"x"},
            {"y"},
            {"foobar"},
        }

        for i, tt := range tests {
            stmt := program.Statements[i]
            if !testLetStatement(t, stmt, tt.expectedIdent) {
                return
            }
        }
    }


    func checkParseErrors(t *testing.T, p *Parser) {
        errors := p.Errors()

        if len(errors) == 0 {
            return 
        }

        t.Errorf("Parser have %d errors", len(errors))
        for _, msg := range errors {
            t.Errorf("Parsing Error: %q", msg)
        }

        t.FailNow()
    }

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
    if s.TokenLiteral() != "let" {
        t.Errorf("Expected token <Let> got <%q> ", s.TokenLiteral())
        return false
    }

    letStmt, ok := s.(*ast.LetStatement)
    if !ok {
        t.Errorf("expected <s>  *ast.LetStatement, got <%T> ", s)
        return false
    }

    if letStmt.Name.Value != name {
        t.Errorf("expected let statement name to be '%s'. got %s",
        name,
        letStmt.Name.Value)
        return false
    }

    if letStmt.Name.Value != name {
        t.Errorf("s.Name not %s is %s",
        name,
        letStmt.Name.Value)
        return false
    }
    return true
}
