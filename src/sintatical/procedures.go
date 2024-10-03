package sintatical

import (
	"fmt"
	"strings"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

func s(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	currNode.Value = S
	currNode.Children = []*types.Node[types.GrammarSymbol]{}

	if proxToken.TokenType == types.KW_MAIN {
		currNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.IDENTIFIER {
			currNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := block(currNode, lex); err != nil {
				return err
			}
			proxToken = lex.GetNextToken()
			return nil
		}
		return fmt.Errorf("expected identifier at %v", proxToken.Position)
	}
	return fmt.Errorf("expected 'main' at %v", proxToken.Position)
}

func block(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	blockNode := currNode.AddChild(BLOCK)
	if proxToken.TokenType == types.KW_BEGIN {
		blockNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := decl(blockNode, lex); err != nil {
			return err
		}
		if err := cmdSeq(blockNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_END {
			blockNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			return nil
		}
		return fmt.Errorf("expected 'end' at %v", proxToken.Position)
	}
	return fmt.Errorf("expected 'begin' at %v", proxToken.Position)
}

func list(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	listNode := currNode.AddChild(LIST)
	if proxToken.TokenType == types.IDENTIFIER {
		listNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := listPrime(listNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.SEMICOLON {
			listNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			return nil
		}
		return fmt.Errorf("expected ';' at %v", proxToken.Position)
	}
	return fmt.Errorf("expected identifier at %v", proxToken.Position)
}

func listPrime(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KKOMA {
		listPrimeNode := currNode.AddChild(LIST_PRIME)
		listPrimeNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := list(listPrimeNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func decl(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_TYPE {
		declNode := currNode.AddChild(DECL)
		proxToken = lex.GetNextToken()
		declNode.AddChild(proxToken)
		if proxToken.TokenType == types.TYPE_SEPARATOR {
			declNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := list(declNode, lex); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("expected %v at %v", types.TYPE_SEPARATOR, proxToken.Position)
	}
	return nil
}

func cmd(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	cmdNode := currNode.AddChild(CMD)
	if isInFirst(CMD_SEL, proxToken.TokenType) {
		if err := cmdSel(cmdNode, lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_REP, proxToken.TokenType) {
		if err := cmdRep(cmdNode, lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_ATR, proxToken.TokenType) {
		if err := cmdAtr(cmdNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("expected one of %v at %v", strings.Join(getFirst(CMD_SEL, CMD_REP, CMD_ATR), ", "), proxToken.Position)
}

func cmdSeq(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	if isInFirst(CMD, proxToken.TokenType) {
		cmdSeqNode := currNode.AddChild(CMD_SEQ)
		if err := cmd(cmdSeqNode, lex); err != nil {
			return err
		}
		if err := cmdSeq(cmdSeqNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func cmdOrBlock(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	cmdOrBlockNode := currNode.AddChild(CMD_OR_BLOCK)
	if isInFirst(CMD, proxToken.TokenType) {
		if err := cmd(cmdOrBlockNode, lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(BLOCK, proxToken.TokenType) {
		if err := block(cmdOrBlockNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("expected one of %v at %v", strings.Join(getFirst(CMD, BLOCK), ", "), proxToken.Position)
}

func cmdSel(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_IF {
		cmdSelNode := currNode.AddChild(CMD_SEL)
		cmdSelNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			cmdSelNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := cond(cmdSelNode, lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				cmdSelNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if err := cmdOrBlock(cmdSelNode, lex); err != nil {
					return err
				}
				if err := cmdSelPrime(cmdSelNode, lex); err != nil {
					return err
				}
				return nil
			}
			return fmt.Errorf("expected ')' at %v", proxToken.Position)
		}
		return fmt.Errorf("expected '(' at %v", proxToken.Position)
	}
	return fmt.Errorf("expected 'if' at %v", proxToken.Position)
}

func cmdSelPrime(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_ELSE {
		cmdSelPrimeNode := currNode.AddChild(CMD_SEL_PRIME)
		cmdSelPrimeNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock(cmdSelPrimeNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func cmdRep(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	cmdRepNode := currNode.AddChild(CMD_REP)
	if proxToken.TokenType == types.KW_WHILE {
		cmdRepNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			cmdRepNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := cond(cmdRepNode, lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				cmdRepNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if err := cmdOrBlock(cmdRepNode, lex); err != nil {
					return err
				}
				return nil
			}
			return fmt.Errorf("expected ')' at %v", proxToken.Position)
		}
		return fmt.Errorf("expected '(' at %v", proxToken.Position)
	}
	if proxToken.TokenType == types.KW_REPEAT {
		if err := block(cmdRepNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_UNTIL {
			cmdRepNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.START_PAREN {
				cmdRepNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if err := cond(cmdRepNode, lex); err != nil {
					return err
				}
				if proxToken.TokenType == types.END_PAREN {
					cmdRepNode.AddChild(proxToken)
					proxToken = lex.GetNextToken()
					return nil
				}
				return fmt.Errorf("expected ')' at %v", proxToken.Position)
			}
			return fmt.Errorf("expected '(' at %v", proxToken.Position)
		}
		return fmt.Errorf("expected 'until' at %v", proxToken.Position)
	}
	return fmt.Errorf("expected 'while' or 'repeat' at %v", proxToken.Position)
}

func cmdAtr(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	cmdAtrNode := currNode.AddChild(CMD_ATR)
	if proxToken.TokenType == types.IDENTIFIER {
		cmdAtrNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.ASSIGN {
			cmdAtrNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := exp(cmdAtrNode, lex); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("expected %v at %v", types.ASSIGN, proxToken.Position)
	}
	return fmt.Errorf("expected identifier at %v", proxToken.Position)
}

func exp(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	expNode := currNode.AddChild(EXP)
	if err := term(expNode, lex); err != nil {
		return err
	}
	if err := expPrime(expNode, lex); err != nil {
		return err
	}
	return nil
}

func expPrime(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	expPrimeNode := currNode.AddChild(EXP_PRIME)
	if proxToken.TokenType == types.ARIOP_SUM {
		expPrimeNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := term(expPrimeNode, lex); err != nil {
			return err
		}
		if err := expPrime(expPrimeNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func term(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	termNode := currNode.AddChild(TERM)
	if err := factor(termNode, lex); err != nil {
		return err
	}
	if err := termPrime(termNode, lex); err != nil {
		return err
	}
	return nil
}

func termPrime(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	termPrimeNode := currNode.AddChild(TERM_PRIME)
	if proxToken.TokenType == types.ARIOP_MULT {
		termPrimeNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := factor(termPrimeNode, lex); err != nil {
			return err
		}
		if err := termPrime(termPrimeNode, lex); err != nil {
			return err
		}
	}
	return nil
}

func factor(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	factorNode := currNode.AddChild(FACTOR)
	if proxToken.TokenType == types.START_PAREN {
		factorNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := exp(factorNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			factorNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			return nil
		}
		return fmt.Errorf("expected ')' at %v", proxToken.Position)
	}
	if proxToken.TokenType == types.CONST {
		factorNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := factorPrime(factorNode, lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("expected '(' or constant at %v", proxToken.Position)
}

func factorPrime(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	factorPrimeNode := currNode.AddChild(FACTOR_PRIME)
	if proxToken.TokenType == types.ARIOP_POW {
		factorPrimeNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := factor(factorPrimeNode, lex); err != nil {
			return err
		}
		if err := factorPrime(factorPrimeNode, lex); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func cond(currNode *types.Node[types.GrammarSymbol], lex lexer.TokenProducer) error {
	condNode := currNode.AddChild(COND)
	if isInFirst(EXP, proxToken.TokenType) {
		if err := exp(condNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.RELOP {
			condNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if err := exp(condNode, lex); err != nil {
				return err
			}
			return nil
		}
		return fmt.Errorf("expected relational operator at %v", proxToken.Position)
	}
	if proxToken.TokenType == types.START_PAREN {
		condNode.AddChild(proxToken)
		proxToken = lex.GetNextToken()
		if err := cond(condNode, lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			condNode.AddChild(proxToken)
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.RELOP {
				condNode.AddChild(proxToken)
				proxToken = lex.GetNextToken()
				if err := cond(condNode, lex); err != nil {
					return err
				}
				if proxToken.TokenType == types.START_PAREN {
					condNode.AddChild(proxToken)
					proxToken = lex.GetNextToken()
					if err := cond(condNode, lex); err != nil {
						return err
					}
					if proxToken.TokenType == types.END_PAREN {
						condNode.AddChild(proxToken)
						proxToken = lex.GetNextToken()
						return nil
					}
					return fmt.Errorf("expected ')' at %v", proxToken.Position)
				}
				return fmt.Errorf("expected '(' at %v", proxToken.Position)
			}
			return fmt.Errorf("expected relational operator at %v", proxToken.Position)
		}
		return fmt.Errorf("expected ')' at %v", proxToken.Position)
	}
	return fmt.Errorf("expected one of %v at %v", strings.Join(append(getFirst(EXP), "("), ", "), proxToken.Position)
}
