package token
type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

var keywords = map[string] TokenType  {
    "fn": Function,
    "let": Let,
    "if": If,
    "else": Else, 
    "return": Return, 
    "true": True, 
    "false": False, 
}

func LookupIdent(ident string) TokenType {
    if t, ok := keywords[ident]; ok {
        return t
    }
    return Ident;
}

const (
    Illegal = "ILLEGAL"
    EOF = "EOF"
    
    Ident = "IDENT"
    Int = "INT"

    Assign = "="
    Plus = "+"
    Minus = "-"
    Bang = "!"
    Asterisk= "*"
    Slash = "/"
    Comma = ","
    SemiColon = ";"
    GT = ">"
    LT = "<"

    LParen = "("
    RParen = ")"
    LBrace = "{"
    RBrace = "}"


    Function = "Function"
    Let = "Let"
    If = "if"
    Else = "else"
    Return = "return"
    True = "true"
    False = "false"

)
