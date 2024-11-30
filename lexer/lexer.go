package lexer

import "monkeylang/token"

type Lexer struct {
	input   string
	pos     int // The current character position wich ch is currently on 
    readpos int // The Next character position
	ch      byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readpos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readpos]
	}
	l.pos = l.readpos

	l.readpos += 1
}

func (l *Lexer) peekChar() byte {
    if l.readpos >= len(l.input) {
        return 0
    } else {
        return l.input[l.readpos]
    }
}

func (l *Lexer) NextToken() token.Token {
    var t token.Token 
    l.skipWhiteSpace()

    switch l.ch {
    case '=':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            t = token.Token{Type: token.EqualTo, Literal: string(ch) + string(l.ch) }
        } else {
            t = newToken(token.Assign, l.ch)
        }
    case '!':
        if l.peekChar() == '=' {
            ch := l.ch
            l.readChar()
            t = token.Token{ Type: token.NotEqualTo, Literal: string(ch) + string(l.ch) }
        } else {
            t = newToken(token.Bang, l.ch)
        }
    case ';':
        t = newToken(token.SemiColon, l.ch)
    case '(':
        t = newToken(token.LParen, l.ch)
    case ')':
        t = newToken(token.RParen, l.ch)
    case '{':
        t = newToken(token.LBrace, l.ch)
    case '}':
        t = newToken(token.RBrace, l.ch)
    case ',':
        t = newToken(token.Comma, l.ch)
    case '+':
        t = newToken(token.Plus, l.ch)
    case '*':
        t = newToken(token.Asterisk, l.ch)
    case '/':
        t = newToken(token.Slash, l.ch)
    case '<':
        t = newToken(token.LT, l.ch)
    case '>':
        t = newToken(token.GT, l.ch)
    case '-':
        t = newToken(token.Minus, l.ch)
    case 0:
        t.Literal = ""
        t.Type = token.EOF

    default: 
        if isLetter(l.ch) {
            t.Literal = l.readIdent()
            t.Type = token.LookupIdent(t.Literal)
            return t } else if isDigit(l.ch) {
            t.Type = token.Int
            t.Literal = l.readNumber()
            return t
        } else {
            t = newToken(token.Illegal, l.ch)
        }
    }

    l.readChar()
    return t
}

func (l *Lexer) readNumber() string {
    pos := l.pos
    for isDigit(l.ch) {
        l.readChar()
    }

    return l.input[pos:l.pos]
}

func  isDigit(ch byte) bool {
    return '0' <= ch && ch <= '9'
}

func (l *Lexer) skipWhiteSpace() { 
    for l.ch == ' ' || l.ch == '\r' || l.ch == '\n' || l.ch == '\t' { 
        l.readChar()
    }
}

func (l *Lexer) readIdent() string {
    pos := l.pos
    for isLetter(l.ch) {
        l.readChar()
    }
    return l.input[pos:l.pos]
}

func isLetter(ch byte) bool {
    return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch  == '_'
}

func newToken(tt token.TokenType, ch byte) token.Token {
    return token.Token{Type: tt, Literal: string(ch)}
}
