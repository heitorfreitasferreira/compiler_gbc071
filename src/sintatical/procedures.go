package sintatical

import (
	"fmt"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

func s(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	node.Value = S
	node.Children = []*types.Node[types.GrammarSymbol]{}
	if proxToken.TokenType == types.KW_MAIN {
		node.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.IDENTIFIER {
			node.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := block(node, lex); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("erro s2")
}

func block(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(BLOCK)
	if proxToken.TokenType == types.KW_BEGIN {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := listSeq(nextNode, lex); err != nil {
			return err
		}
		if err := cmdSeq(nextNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_END {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			return nil
		}
	}
	return fmt.Errorf("erro block2")
}

func listInit(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(LIST_INIT)
	if proxToken.TokenType == types.KW_TYPE {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.TYPE_SEPARATOR {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.IDENTIFIER {
				nextNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if err := listEnd(nextNode, lex); err != nil {
					return err
				}
				return nil
			}
		}
	}
	return fmt.Errorf("erro listInit2")
}

func listMult(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(LIST_MULT)
	if proxToken.TokenType == types.KKOMA {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.IDENTIFIER {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			return nil
		}
	}
	return fmt.Errorf("erro listMult2")
}

func listEnd(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node
	for isInFirst(LIST_MULT, proxToken.TokenType) {
		nextNode = nextNode.AddChild(LIST_END)
		if err := listMult(nextNode, lex); err != nil {
			return err
		}
	}
	if proxToken.TokenType == types.SEMICOLON {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		return nil
	}
	return fmt.Errorf("erro listEnd2")
}

func listSeq(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node
	for isInFirst(LIST_INIT, proxToken.TokenType) {
		nextNode = nextNode.AddChild(LIST_SEQ)
		if err := listInit(nextNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func cmd(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(CMD)
	if isInFirst(CMD_SEL, proxToken.TokenType) {
		if err := cmdSel(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_REP, proxToken.TokenType) {
		if err := cmdRep(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_ATR, proxToken.TokenType) {
		if err := cmdAtr(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro cmd2")
}

func cmdSeq(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node
	for isInFirst(CMD, proxToken.TokenType) {
		nextNode = nextNode.AddChild(CMD_SEQ)
		if err := cmd(nextNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func cmdOrBlock(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(CMD_OR_BLOCK)
	if isInFirst(CMD, proxToken.TokenType) {
		if err := cmd(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(BLOCK, proxToken.TokenType) {
		if err := block(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro cmdOrBlock2 %v", proxToken)
}

func cmdSel(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(CMD_SEL)
	if proxToken.TokenType == types.KW_IF {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := cond(nextNode, lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				nextNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.KW_THEN {
					nextNode.AddChild(proxToken)
					proxToken = lex.GetNextToken()
					if err := cmdOrBlock(nextNode, lex); err != nil {
						return err
					}
					if err := cmdSelPrime(nextNode, lex); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("erro cmdSel2")
}

func cmdSelPrime(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(CMD_SEL_PRIME)
	if proxToken.TokenType == types.KW_ELSE {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func cmdRep(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(CMD_REP)
	if proxToken.TokenType == types.KW_WHILE {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := cond(nextNode, lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				nextNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.KW_DO {
					nextNode.AddChild(proxToken)
					proxToken = lex.GetNextToken()
					if err := cmdOrBlock(nextNode, lex); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	if proxToken.TokenType == types.KW_REPEAT {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock(nextNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_UNTIL {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.START_PAREN {
				nextNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if err := cond(nextNode, lex); err != nil {
					return err
				}
				if proxToken.TokenType == types.END_PAREN {
					nextNode.AddChild(proxToken)
					proxToken = lex.GetNextToken()
					if proxToken.TokenType == types.SEMICOLON {
						nextNode.AddChild(proxToken)
						proxToken = lex.GetNextToken()
						return nil
					}
				}
			}
		}
	}
	return fmt.Errorf("erro cmdRep2 %v", proxToken)
}

func cmdAtr(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(CMD_ATR)
	if proxToken.TokenType == types.IDENTIFIER {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.ASSIGN {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := exp(nextNode, lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.SEMICOLON {
				nextNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				return nil
			}
		}
	}
	return fmt.Errorf("erro cmdAtr2: %v", proxToken)
}

func exp(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(EXP)
	if err := term(nextNode, lex); err != nil {
		return err
	}
	if err := expPrime(nextNode, lex); err != nil {
		return err
	}
	return nil
}

func expPrime(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node
	for proxToken.TokenType == types.ARIOP_SUM {
		nextNode = nextNode.AddChild(EXP_PRIME)
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := term(nextNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func term(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(TERM)
	if err := factor(nextNode, lex); err != nil {
		return err
	}
	if err := termPrime(nextNode, lex); err != nil {
		return err
	}
	return nil
}

func termPrime(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node
	for proxToken.TokenType == types.ARIOP_MULT {
		nextNode = nextNode.AddChild(TERM_PRIME)
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := factor(nextNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func idOrConst(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(ID_OR_CONST)
	if proxToken.TokenType == types.IDENTIFIER {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		return nil
	}
	if proxToken.TokenType == types.CONST {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		return nil
	}
	return fmt.Errorf("erro idOrConst")
}

func unaryExp(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(UNARY_EXP)
	if proxToken.TokenType == types.ARIOP_SUM {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := idOrConst(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro unaryExp")
}

func factor(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(FACTOR)
	if isInFirst(UNARY_EXP, proxToken.TokenType) {
		if err := unaryExp(nextNode, lex); err != nil {
			return err
		}
		if err := factorPrime(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	if proxToken.TokenType == types.START_PAREN {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := exp(nextNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := factorPrime(nextNode, lex); err != nil {
				return err
			}
			return nil
		}
	}
	if isInFirst(ID_OR_CONST, proxToken.TokenType) {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := factorPrime(nextNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro factor2")
}

func factorPrime(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node
	for proxToken.TokenType == types.ARIOP_POW {
		nextNode = nextNode.AddChild(FACTOR_PRIME)
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := factor(nextNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func cond(node *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	nextNode := node.AddChild(COND)
	if isInFirst(EXP, proxToken.TokenType) {
		if err := exp(nextNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.RELOP {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := exp(nextNode, lex); err != nil {
				return err
			}
			return nil
		}
	}

	if proxToken.TokenType == types.START_PAREN {
		nextNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := cond(nextNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			nextNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.RELOP {
				nextNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.START_PAREN {
					nextNode.AddChild(proxToken)
					proxToken = lex.GetNextToken()
					if err := cond(nextNode, lex); err != nil {
						return err
					}
					if proxToken.TokenType == types.END_PAREN {
						nextNode.AddChild(proxToken)
						proxToken = lex.GetNextToken()
						return nil
					}
				}
			}
		}
	}
	return fmt.Errorf("erro cond2")
}
