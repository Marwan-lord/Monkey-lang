package parser

import (
	"fmt"
	"monkeylang/ast"
	"monkeylang/lexer"
	"monkeylang/token"
	"strconv"
)

const (
     _ int = iota 
     LOWEST
     EQUALS // == 
     LESSGREATER // > or < 
     SUM // + 
     PRODUCT // * 
     PREFIX // -X or !X 
     CALL // myFunction(X)
)

var precedences = map[token.TokenType]int {
    token.EqualTo: EQUALS,
    token.NotEqualTo: EQUALS,
    token.LT: LESSGREATER,
    token.GT: LESSGREATER,
    token.Plus: SUM,
    token.Minus: SUM,
    token.Slash: PRODUCT,
    token.Asterisk: PRODUCT,
}

type (
    prefixParseFn func() ast.Expression
    infixParseFn func(ast.Expression) ast.Expression
)

type Parser struct {
    l *lexer.Lexer
    curToken token.Token
    peekToken token.Token
    errors []string
    prefixParseFn map[token.TokenType]prefixParseFn
    infixParseFn map[token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
    p.prefixParseFn[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
    p.infixParseFn[tokenType] = fn
}

func (p *Parser) peekPrecedence() int {
    if p, ok := precedences[p.peekToken.Type]; ok {
        return p
    }

    return LOWEST
}

func (p *Parser) currentPrecedence() int {
    if p, ok := precedences[p.curToken.Type]; ok {
        return p
    }

    return LOWEST
}

func New(l *lexer.Lexer) *Parser {
    p := &Parser{
        l: l,
        errors: []string{},
    }
    p.prefixParseFn = make(map[token.TokenType]prefixParseFn)
    p.infixParseFn = make(map[token.TokenType]infixParseFn)

    p.registerPrefix(token.Ident, p.parseIdentifier)
    p.registerPrefix(token.Int, p.parseIntegerLiteral)
    p.registerPrefix(token.Bang, p.parsePrefixExpression)
    p.registerPrefix(token.Minus, p.parsePrefixExpression)

    p.registerInfix(token.Plus, p.parseInfixExpression)
    p.registerInfix(token.Minus, p.parseInfixExpression)
    p.registerInfix(token.Slash, p.parseInfixExpression)
    p.registerInfix(token.Asterisk, p.parseInfixExpression)
    p.registerInfix(token.EqualTo, p.parseInfixExpression)
    p.registerInfix(token.NotEqualTo, p.parseInfixExpression)
    p.registerInfix(token.LT, p.parseInfixExpression)
    p.registerInfix(token.GT, p.parseInfixExpression)

    // We properly set up the curToken and peekToken fields
    p.nextToken()
    p.nextToken()

    return p
}

func (p *Parser) parsePrefixExpression() ast.Expression {
    exp := &ast.PrefixExpression{
        Token: p.curToken,
        Operator: p.curToken.Literal,
    }
    p.nextToken()
    exp.Right = p.parseExpression(PREFIX)

    return exp
}

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
    expression := &ast.InfixExpression{
        Token: p.curToken,
        Operator: p.curToken.Literal,
        Left: left,
    }

    precedence := p.currentPrecedence()
    p.nextToken()
    expression.Right = p.parseExpression(precedence)

    return expression
}

func (p *Parser) parseIdentifier() ast.Expression {
    return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) Errors() []string {
    return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
    msg := fmt.Sprintf("Expected %s , got %s instead", 
    t, p.peekToken.Type)
    p.errors = append(p.errors, msg)
}

func (p *Parser) nextToken() {
    p.curToken = p.peekToken
    p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
    program := &ast.Program{}
    program.Statements = []ast.Statement{}

    for p.curToken.Type != token.EOF {
        stmt := p.parseStatement()

        if stmt != nil {
            program.Statements = append(program.Statements, stmt)
        }
        p.nextToken()
    }

    return program
}

func (p *Parser) parseStatement() ast.Statement {
    switch p.curToken.Type {

    case token.Let: 
        return p.parseLetStatement()
    case token.Return:
        return p.parseReturnStatement()

    default:
        return p.parseExpressionStatement()
    }
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
    stmt := &ast.ExpressionStatement{Token: p.curToken}
    stmt.Expression = p.parseExpression(LOWEST)
    if p.peekTokenIs(token.SemiColon) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseExpression(prec int) ast.Expression {
    prefix := p.prefixParseFn[p.curToken.Type]

    if prefix == nil {
        p.noPrefixParseError(p.curToken.Type)
        return nil
    }
    leftExp := prefix()

    for !p.peekTokenIs(token.SemiColon) && prec < p.peekPrecedence() {
        infix := p.infixParseFn[p.peekToken.Type]
        if infix == nil {
            return leftExp
        }
        p.nextToken()

        leftExp = infix(leftExp)
    }

    return leftExp
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement{
    stmt := &ast.ReturnStatement{Token: p.curToken}
    p.nextToken()

    // TODO: Actually parse any thing not just skip
    for !p.curTokenIs(token.SemiColon) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
    lit := &ast.IntegerLiteral{Token: p.curToken}
    value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
    if err != nil {
        msg := fmt.Sprintf("couldn't parse %q as integer", p.curToken.Literal)
        p.errors = append(p.errors, msg)
        return nil
    }
    lit.Value = value
    return lit
}

func (p *Parser) parseLetStatement() *ast.LetStatement{
    stmt := &ast.LetStatement{Token: p.curToken}
    if !p.expectPeek(token.Ident) {
        return nil
    }

    stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

    if !p.expectPeek(token.Assign) {
        return nil
    }

    // TODO: Actually parse any thing not just skip
    for !p.curTokenIs(token.SemiColon) {
        p.nextToken()
    }

    return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
    return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
    return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
    if p.peekTokenIs(t) {
        p.nextToken()
        return true
    } else {
        p.peekError(t)
        return false
    }
}

func (p *Parser) noPrefixParseError(t token.TokenType) {
    msg := fmt.Sprintf("no prefix parser function for %s found", t)
    p.errors = append(p.errors, msg)
}

