package types

type Token struct {
	TokenType
	Position
	Lexeme string
	Attr   map[string]interface{}
}

type Position struct {
	Line   int
	Column int
}

type TokenType string

const (
	EOF        = "EOF"
	TYPE       = "TYPE"
	IDENTIFIER = "IDENTIFIER"

	ARIOP = "ARIOP"
	RELOP = "RELOP"

	ASSIGN = "ASSIGN"

	START_BLOCK = "START_BLOCK"
	END_BLOCK   = "END_BLOCK"
	START_PAREN = "START_PAREN"
	END_PAREN   = "END_PAREN"

	KEYWORD   = "KEYWORD"
	COMMENT   = "COMMENT"
	SEPARATOR = "SEPARATOR"

	TYPE_SEPARATOR = "TYPE_SEPARATOR"

	CONSTANT = "CONSTANT"
)
