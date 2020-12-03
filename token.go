package main

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	TOK_ILLEGAL = "ILLEGAL"
	TOK_EOF     = "EOF"

	TOK_IDENT = "IDENT"
	TOK_NUM   = "NUM"

	TOK_AT        = "@"
	TOK_COMMA     = ","
	TOK_SEMICOLON = ";"
	TOK_ASSIGN    = "="

	TOK_LPAREN   = "("
	TOK_RPAREN   = ")"
	TOK_LBRACE   = "{"
	TOK_RBRACE   = "}"
	TOK_LBRACKET = "["
	TOK_RBRACKET = "]"

	TOK_ARROW = "->"

	TOK_COMMENT = "COMMENT"

	// keywords
	TOK_PROGRAM   = "PROGRAM"
	TOK_INTERFACE = "INTERFACE"
	TOK_IMPORT    = "IMPORT"
	TOK_ENUM      = "ENUM"
	TOK_STRUCT    = "STRUCT"
	TOK_TYPE      = "TYPE"
	TOK_ONEWAY    = "ONEWAY"

	TOK_STRING = "STRING"
)

var keywords = map[string]TokenType{
	"program":   TOK_PROGRAM,
	"interface": TOK_INTERFACE,
	"import":    TOK_IMPORT,
	"enum":      TOK_ENUM,
	"struct":    TOK_STRUCT,
	"string":    TOK_TYPE,
	"uint8":     TOK_TYPE,
	"uint16":    TOK_TYPE,
	"uint32":    TOK_TYPE,
	"int8":      TOK_TYPE,
	"int16":     TOK_TYPE,
	"int32":     TOK_TYPE,
	"void":      TOK_TYPE,
	"bool":      TOK_TYPE,
	"in":        TOK_TYPE,
	"inout":     TOK_TYPE,
	"out":       TOK_TYPE,
	"binary":    TOK_TYPE,
	"oneway":    TOK_ONEWAY,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOK_IDENT
}
