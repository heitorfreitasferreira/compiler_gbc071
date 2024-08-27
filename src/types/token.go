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
	IDENTIFIER = "IDENTIFIER"

	ARIOP_POW  = "ARIOP_POW"
	ARIOP_MULT = "ARIOP_MULT"
	ARIOP_ADD  = "ARIOP_ADD"
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

	TYPE_SEPARATOR = "TYPE_SEPARATOR"

	NUM_CONST = "NUM_CONST"
)
