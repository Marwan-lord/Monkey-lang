package parser

import (
	"fmt"
	"monkeylang/ast"
	"monkeylang/lexer"
	"testing"
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

func TestIdentifierExpression(t *testing.T) {
    input := "foobar;"
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParseErrors(t, p)
    if len(program.Statements) != 1 {
        t.Fatalf("program does not have enough statements")
    }

    stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statement[0] is not an Expression. got=%T",
        program.Statements[0])
    }

    ident, ok := stmt.Expression.(*ast.Identifier)
    if !ok {
        t.Fatalf("exp not *ast.Identifier. got=%T",
        stmt.Expression)
    }
    if ident.Value != "foobar" {
        t.Errorf("ident.Value is wrong got %s", ident.Value)
    }

    if ident.TokenLiteral() != "foobar" {
        t.Errorf("ident Token literal does not match \"foobar\"got %s",
        ident.TokenLiteral())
    }
}

func TestIntegerExpression(t *testing.T) {
    input := "5;"
    l := lexer.New(input)
    p := New(l)
    program := p.ParseProgram()
    checkParseErrors(t, p)

    if len(program.Statements) != 1 {
        t.Fatalf("program does not have enough statements")
    }

    stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
    if !ok {
        t.Fatalf("program.Statement[0] is not an Expression. got=%T",
        program.Statements[0])
    }

    ident, ok := stmt.Expression.(*ast.IntegerLiteral)
    if !ok {
        t.Fatalf("exp not *ast.IntegerLiteral. got=%T",
        stmt.Expression)
    }

    if ident.Value != 5 {
        t.Errorf("Literal.Value is wrong got %d", ident.Value)
    }

    if ident.TokenLiteral() != "5" {
        t.Errorf("ident Token literal does not match \"5\"got %s",
        ident.TokenLiteral())
    }
}

func TestParsingPrefixExpressions(t *testing.T) {
    prefixTest := []struct {
        input string
        op    string
        intVal int64
    } {
        {"!5", "!", 5 },
        {"-15", "-", 15 },
    }

    for _, tt := range prefixTest {
        l := lexer.New(tt.input)
        p := New(l)
        prog := p.ParseProgram()
        checkParseErrors(t, p)

        if len(prog.Statements) != 1 {
            t.Fatalf("program does not have enough statements")
        }

        stmt, ok := prog.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statement[0] is not an Expression. got=%T",
            prog.Statements[0])
        }

        exp, ok := stmt.Expression.(*ast.PrefixExpression)
        if !ok {
            t.Fatalf("exp not *ast.PrefixExpression . got=%T",
            stmt.Expression)
        }

        if exp.Operator != tt.op {
            t.Fatalf("exp.Operator is wrong expected %s got %s",
            tt.op,exp.Operator)
        }

        if !testIntegerLiteral(t, exp.Right, tt.intVal){
            t.Errorf("ident Token literal does not match \"5\"got %s",
            exp.TokenLiteral())
        }
    }
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
    intg, ok := il.(*ast.IntegerLiteral)
    if !ok {
        t.Errorf("il is not an IntegerLiteral got %T", il)
        return false
    }

    if intg.Value != value {
        t.Errorf("intg.Value is not right expected %d got %d", value, intg.Value)
        return false
    }

    if intg.TokenLiteral() != fmt.Sprintf("%d", value) {
        t.Errorf("intg Token Value is not right expected %d got %s", value, 
    intg.TokenLiteral())
        return false

    }

    return true
}

func TestParsingInfifxExpressions(t *testing.T) {
    infexTest := []struct {
        input string
        leftValue int64
        operator string
        rightValue int64
    } {
        {"5 + 5;", 5, "+", 5},
        {"5 - 5;", 5, "-", 5},
        {"5 * 5;", 5, "*", 5},
        {"5 / 5;", 5, "/", 5},
        {"5 > 5;", 5, ">", 5},
        {"5 < 5;", 5, "<", 5},
        {"5 == 5;", 5, "==", 5},
        {"5 != 5;", 5, "!=", 5},
    }
    for _, tt := range infexTest {
        l := lexer.New(tt.input)
        p := New(l)
        prog := p.ParseProgram()
        checkParseErrors(t, p)

        if len(prog.Statements) != 1 {
            t.Fatalf("program Statements is not enough got %d", len(prog.Statements))
        }

        stmt, ok := prog.Statements[0].(*ast.ExpressionStatement)
        if !ok {
            t.Fatalf("program.Statement[0] is not an Expression. got=%T",
            prog.Statements[0])
        }

        exp, ok := stmt.Expression.(*ast.InfixExpression)
        if !ok {
            t.Fatalf("exp not *ast.InfixExpression . got=%T",
            stmt.Expression)
        }

        if !testIntegerLiteral(t, exp.Left, tt.leftValue) {
            return 
        }

        if exp.Operator != tt.operator {
            t.Fatalf("exp.Operator is wrong expected %s got %s",
            tt.operator ,exp.Operator)
        }

        if !testIntegerLiteral(t, exp.Right, tt.rightValue){
            return 
        }
    }
}

func TestOperatorPrecedenceParsing(t *testing.T) {
    tests := []struct {
        input string
        expeceted string
    } {
        {
         "-a * b", "((-a) * b)",
     },

     {"!-a", "(!(-a))", },
     { "a + b + c", "((a + b) + c)", },
     {"a * b * c", "((a * b) * c)", },
     { "a * b / c", "((a * b) / c)", },
     { "a + b / c", "(a + (b / c))", },
     { "a + b * c + d / e - f", "(((a + (b * c)) + (d / e)) - f)", },
     { "3 + 4; -5 * 5", "(3 + 4)((-5) * 5)", },
     { "5 > 4 == 3 < 4", "((5 > 4) == (3 < 4))", },
     { "5 < 4 != 3 > 4", "((5 < 4) != (3 > 4))", },
     { "3 + 4 * 5 == 3 * 1 + 4 * 5", "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))", },

    }

    for _, tt := range tests {
        l := lexer.New(tt.input)
        p := New(l)
        program := p.ParseProgram()
        checkParseErrors(t, p)
        actual := program.String()
        if actual != tt.expeceted {
            t.Errorf("Parsing Error expected %q got %q", tt.expeceted, actual)
        }

    }
}
