package sintatical

import "github.com/heitorfreitasferreira/compiler/types"

var first = map[types.GrammarSymbol][]types.TokenType{
	S:             {types.KW_MAIN},
	BLOCK:         {types.KW_BEGIN},
	LIST:          {types.IDENTIFIER},
	LIST_PRIME:    {types.KKOMA},                                   // WHILE
	DECL:          {types.KW_TYPE},                                 // WHILE
	CMD_SEQ:       {types.KW_IF, types.KW_WHILE, types.IDENTIFIER}, // WHILE
	CMD_OR_BLOCK:  {types.KW_IF, types.KW_WHILE, types.KW_REPEAT, types.IDENTIFIER, types.KW_BEGIN},
	CMD_SEL:       {types.KW_IF},
	CMD_SEL_PRIME: {types.KW_ELSE},
	CMD_REP:       {types.KW_WHILE, types.KW_REPEAT},
	EXP:           {types.START_PAREN, types.CONST},
	EXP_PRIME:     {types.ARIOP_SUM}, //WHILE
	TERM:          {types.START_PAREN, types.CONST},
	TERM_PRIME:    {types.ARIOP_MULT}, //WHILE
	FACTOR:        {types.START_PAREN, types.CONST},
	FACTOR_PRIME:  {types.RELOP}, //WHILE
	COND:          {types.START_PAREN, types.CONST},
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
