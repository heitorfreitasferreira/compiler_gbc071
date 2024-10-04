package sintatical

import (
	"fmt"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

func s(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_MAIN {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.IDENTIFIER {
			proxToken = lex.GetNextToken()
			if err := block(lex); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("erro s2")
}

func block(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_BEGIN {
		proxToken = lex.GetNextToken()
		if err := listSeq(lex); err != nil {
			return err
		}
		if err := cmdSeq(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_END {
			proxToken = lex.GetNextToken()
			return nil
		}
	}
	return fmt.Errorf("erro block2")
}

func listInit(lex lexer.TokenProducer) error {
    if proxToken.TokenType == types.KW_TYPE {
        proxToken = lex.GetNextToken()
        if proxToken.TokenType == types.TYPE_SEPARATOR {
            proxToken = lex.GetNextToken()
            if proxToken.TokenType == types.IDENTIFIER {
                proxToken = lex.GetNextToken()
                if err := listEnd(lex); err != nil {
                    return err
                }
                return nil
            }
        }
    }
    return fmt.Errorf("erro listInit2")
}

func listMult(lex lexer.TokenProducer) error {
    if proxToken.TokenType == types.KKOMA {
        proxToken = lex.GetNextToken()
        if proxToken.TokenType == types.IDENTIFIER {
            proxToken = lex.GetNextToken()
            return nil
        }
    }
    return fmt.Errorf("erro listMult2")
}

func listEnd(lex lexer.TokenProducer) error {
    for isInFirst(LIST_MULT, proxToken.TokenType) {
        if err := listMult(lex); err != nil {
            return err
        }
    }
    if proxToken.TokenType == types.SEMICOLON {
        proxToken = lex.GetNextToken()
        return nil
    }
    return fmt.Errorf("erro listEnd2")
}

func listSeq(lex lexer.TokenProducer) error {
    for isInFirst(LIST_INIT, proxToken.TokenType) {
        if err := listInit(lex); err != nil {
            return err
        }
    }
    return nil
}

func cmd(lex lexer.TokenProducer) error {
	if isInFirst(CMD_SEL, proxToken.TokenType) {
		if err := cmdSel(lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_REP, proxToken.TokenType) {
		if err := cmdRep(lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_ATR, proxToken.TokenType) {
		if err := cmdAtr(lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro cmd2")
}

func cmdSeq(lex lexer.TokenProducer) error {
	for isInFirst(CMD, proxToken.TokenType) {
		if err := cmd(lex); err != nil {
			return err
		}
	}
	return nil
}

func cmdOrBlock(lex lexer.TokenProducer) error {
	if isInFirst(CMD, proxToken.TokenType) {
		if err := cmd(lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(BLOCK, proxToken.TokenType) {
		if err := block(lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro cmdOrBlock2 %v", proxToken)
}

func cmdSel(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_IF {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			proxToken = lex.GetNextToken()
			if err := cond(lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.KW_THEN {
					proxToken = lex.GetNextToken()
					if err := cmdOrBlock(lex); err != nil {
						return err
					}
					if err := cmdSelPrime(lex); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("erro cmdSel2")
}

func cmdSelPrime(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_ELSE {
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock(lex); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func cmdRep(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_WHILE {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			proxToken = lex.GetNextToken()
			if err := cond(lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.KW_DO {
                    proxToken = lex.GetNextToken()
					if err := cmdOrBlock(lex); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	if proxToken.TokenType == types.KW_REPEAT {
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_UNTIL {
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.START_PAREN {
				proxToken = lex.GetNextToken()
				if err := cond(lex); err != nil {
					return err
				}
				if proxToken.TokenType == types.END_PAREN {
					proxToken = lex.GetNextToken()
					if proxToken.TokenType == types.SEMICOLON {
						proxToken = lex.GetNextToken()
						return nil
					}
				}
			}
		}
	}
	return fmt.Errorf("erro cmdRep2 %v", proxToken)
}

func cmdAtr(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.IDENTIFIER {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.ASSIGN {
			proxToken = lex.GetNextToken()
			if err := exp(lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.SEMICOLON {
				proxToken = lex.GetNextToken()
				return nil
			}
		}
	}
    return fmt.Errorf("erro cmdAtr2: %v", proxToken)
}

func exp(lex lexer.TokenProducer) error {
	if err := term(lex); err != nil {
		return err
	}
	if err := expPrime(lex); err != nil {
		return err
	}
	return nil
}

func expPrime(lex lexer.TokenProducer) error {
	for proxToken.TokenType == types.ARIOP_SUM {
		proxToken = lex.GetNextToken()
		if err := term(lex); err != nil {
			return err
		}
	}
	return nil
}

func term(lex lexer.TokenProducer) error {
	if err := factor(lex); err != nil {
		return err
	}
	if err := termPrime(lex); err != nil {
		return err
	}
	return nil
}

func termPrime(lex lexer.TokenProducer) error {
	for proxToken.TokenType == types.ARIOP_MULT {
		proxToken = lex.GetNextToken()
		if err := factor(lex); err != nil {
			return err
		}
	}
	return nil
}

func idOrConst(lex lexer.TokenProducer) error {
    if proxToken.TokenType == types.IDENTIFIER {
        proxToken = lex.GetNextToken()
        return nil
    }
    if proxToken.TokenType == types.CONST {
        proxToken = lex.GetNextToken()
        return nil
    }
    return fmt.Errorf("erro idOrConst")
}

func unaryExp(lex lexer.TokenProducer) error {
    if proxToken.TokenType == types.ARIOP_SUM {
        proxToken = lex.GetNextToken()
        if err := idOrConst(lex); err != nil {
            return err
        }
        return nil
    }
    return fmt.Errorf("erro unaryExp")
}

func factor(lex lexer.TokenProducer) error {
    if isInFirst(UNARY_EXP, proxToken.TokenType) {
        if err := unaryExp(lex); err != nil {
            return err
        }
        if err := factorPrime(lex); err != nil {
            return err
        }
        return nil
    }
	if proxToken.TokenType == types.START_PAREN {
		proxToken = lex.GetNextToken()
		if err := exp(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			proxToken = lex.GetNextToken()
			if err := factorPrime(lex); err != nil {
                return err
                            }
            return nil
		}
	}
	if isInFirst(ID_OR_CONST, proxToken.TokenType) {
		proxToken = lex.GetNextToken()
		if err := factorPrime(lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro factor2")
}

func factorPrime(lex lexer.TokenProducer) error {
	for proxToken.TokenType == types.ARIOP_POW {
		proxToken = lex.GetNextToken()
		if err := factor(lex); err != nil {
			return err
		}
	}
	return nil
}

func cond(lex lexer.TokenProducer) error {
	if isInFirst(EXP, proxToken.TokenType) {
		if err := exp(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.RELOP {
			proxToken = lex.GetNextToken()
			if err := exp(lex); err != nil {
				return err
			}
			return nil
		}
	}
	if proxToken.TokenType == types.START_PAREN {
		proxToken = lex.GetNextToken()
		if err := cond(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.RELOP {
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.START_PAREN {
					proxToken = lex.GetNextToken()
					if err := cond(lex); err != nil {
						return err
					}
					if proxToken.TokenType == types.END_PAREN {
						proxToken = lex.GetNextToken()
						return nil
					}
				}
			}
		}
	}
	return fmt.Errorf("erro cond2")
}
