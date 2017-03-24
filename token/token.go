package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"
	ARR   = "[]"

	ASSIGN   = "="
	ASTERISK = "*"
	BANG     = "!"
	MINUS    = "-"
	PLUS     = "+"
	SLASH    = "/"

	LT     = "<"
	GT     = ">"
	EQ     = "=="
	NOT_EQ = "!="
	INC    = "++"
	DEC    = "--"

	SEMICOLON = ";"
	COMMA     = ","

	LPAREN   = "("
	RPAREN   = ")"
	LBRACE   = "{"
	RBRACE   = "}"
	LBRACKET = "["
	RBRACKET = "]"

	LET    = "LET"
	TRUE   = "TRUE"
	FALSE  = "FALSE"
	IF     = "IF"
	ELSE   = "ELSE"
	STRING = "STRING"
	INPUT  = "INPUT"
	PRINT  = "PRINT"
	FOR    = "FOR"
	WHILE  = "WHILE"
	CONST  = "CONST"
)

var keywords = map[string]TokenType{
	"let":   LET,
	"true":  TRUE,
	"false": FALSE,
	"if":    IF,
	"else":  ELSE,
	"input": INPUT,
	"print": PRINT,
	"for":   FOR,
	"while": WHILE,
	"const": CONST,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
