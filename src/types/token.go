package types

import "fmt"

type Token struct {
	TokenType
	Position
	Lexeme string
	Id     *int
}

type Position struct {
	Line   int
	Column int
}

type TokenType string

const (
	EOF        = "EOF"
	IDENTIFIER = "IDENTIFIER"

	ARIOP_POW  = "ARIOP_POW"
	ARIOP_MULT = "ARIOP_MULT"
	ARIOP_SUM  = "ARIOP_SUM"
	RELOP      = "RELOP"

	ASSIGN = "ASSIGN"

	START_PAREN = "START_PAREN"
	END_PAREN   = "END_PAREN"

	KW_MAIN   = "MAIN"
	KW_BEGIN  = "START_BLOCK"
	KW_END    = "END_BLOCK"
	KW_TYPE   = "TYPE"
	KW_IF     = "IF"
	KW_THEN   = "THEN"
	KW_ELSE   = "ELSE"
	KW_WHILE  = "WHILE"
	KW_DO     = "DO"
	KW_REPEAT = "REPEAT"
	KW_UNTIL  = "UNTIL"

	COMMENT = "COMMENT"

	KKOMA     = "COMMA"
	SEMICOLON = "SEMICOLON"
	SEPARATOR = "SEPARATOR"

	TYPE_SEPARATOR = "TYPE_SEPARATOR"

	CONST = "CONST"
)

func (tk Token) String() string {
	return tk.Lexeme
}

func (p Position) String() string {
	return fmt.Sprintf("line %d, column %d", p.Line, p.Column)
}
