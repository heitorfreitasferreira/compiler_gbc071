package sintatical

import "github.com/heitorfreitasferreira/compiler/types"

var first = map[types.GrammarSymbol][]types.TokenType{
	S:             {types.KW_MAIN},
	BLOCK:         {types.KW_BEGIN},
	LIST_INIT:     {types.KW_TYPE},
	LIST_MULT:     {types.KKOMA},
	LIST_END:      {types.KKOMA, types.SEMICOLON},
	LIST_SEQ:      {types.KW_TYPE}, // WHILE
	CMD:           {types.KW_IF, types.KW_WHILE, types.KW_REPEAT, types.IDENTIFIER},
	CMD_SEQ:       {types.KW_IF, types.KW_WHILE, types.IDENTIFIER, types.KW_REPEAT}, // WHILE
	CMD_OR_BLOCK:  {types.KW_IF, types.KW_WHILE, types.KW_REPEAT, types.IDENTIFIER, types.KW_BEGIN},
	CMD_SEL:       {types.KW_IF},
	CMD_SEL_PRIME: {types.KW_ELSE},
	CMD_REP:       {types.KW_WHILE, types.KW_REPEAT},
	CMD_ATR:       {types.IDENTIFIER},
	EXP:           {types.ARIOP_SUM, types.START_PAREN, types.CONST, types.IDENTIFIER},
	EXP_PRIME:     {types.ARIOP_SUM}, //WHILE
	TERM:          {types.ARIOP_SUM, types.START_PAREN, types.CONST, types.IDENTIFIER},
	TERM_PRIME:    {types.ARIOP_MULT}, //WHILE
	FACTOR:        {types.ARIOP_SUM, types.START_PAREN, types.CONST, types.IDENTIFIER},
	FACTOR_PRIME:  {types.ARIOP_POW}, //WHILE
	UNARY_EXP:     {types.ARIOP_SUM},
	ID_OR_CONST:   {types.CONST, types.IDENTIFIER},
	COND:          {types.ARIOP_SUM, types.START_PAREN, types.CONST, types.IDENTIFIER},
}

func isInFirst(nt types.GrammarSymbol, token types.TokenType) bool {
	for _, t := range first[nt] {
		if t == token {
			return true
		}
	}
	return false
}
func getFirst(nts ...types.GrammarSymbol) []string {
	var tokens []string
	for _, nt := range nts {
		for _, token := range first[nt] {
			tokens = append(tokens, string(token))
		}
	}
	return tokens
}
