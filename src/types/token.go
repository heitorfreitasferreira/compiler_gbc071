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
	EOF        TokenType = "EOF"
	IDENTIFIER TokenType = "ID"

	ARIOP_POW  TokenType = "ARIOP_POW"
	ARIOP_MULT TokenType = "ARIOP_MULT"
	ARIOP_SUM  TokenType = "ARIOP_SUM"
	RELOP      TokenType = "RELOP"

	ASSIGN TokenType = ":="

	START_PAREN TokenType = "("
	END_PAREN   TokenType = ")"

	KW_MAIN   TokenType = "MAIN"
	KW_BEGIN  TokenType = "BEGIN"
	KW_END    TokenType = "END"
	KW_TYPE   TokenType = "TYPE"
	KW_IF     TokenType = "IF"
	KW_THEN   TokenType = "THEN"
	KW_ELSE   TokenType = "ELSE"
	KW_WHILE  TokenType = "WHILE"
	KW_DO     TokenType = "DO"
	KW_REPEAT TokenType = "REPEAT"
	KW_UNTIL  TokenType = "UNTIL"

	COMMENT TokenType = "COMMENT"

	KKOMA     TokenType = ","
	SEMICOLON TokenType = ";"
	SEPARATOR TokenType = "SEPARATOR"

	TYPE_SEPARATOR TokenType = ":"

	CONST TokenType = "CONST"

	EMPTY TokenType = ""
)

func (tk Token) String() string {
	return string(tk.TokenType)
}

func (p Position) String() string {
	return fmt.Sprintf("line %d, column %d", p.Line, p.Column)
}
