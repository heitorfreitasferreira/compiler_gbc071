package sintatical

import (
	"fmt"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

func s2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_MAIN {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.IDENTIFIER {
			proxToken = lex.GetNextToken()
			if err := block2(lex); err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("erro s2")
}

func block2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_BEGIN {
		proxToken = lex.GetNextToken()
		if err := listSeq2(lex); err != nil {
			return err
		}
		if err := cmdSeq2(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_END {
			proxToken = lex.GetNextToken()
			return nil
		}
	}
	return fmt.Errorf("erro block2")
}

func listInit2(lex lexer.TokenProducer) error {
    if proxToken.TokenType == types.KW_TYPE {
        proxToken = lex.GetNextToken()
        if proxToken.TokenType == types.TYPE_SEPARATOR {
            proxToken = lex.GetNextToken()
            if proxToken.TokenType == types.IDENTIFIER {
                proxToken = lex.GetNextToken()
                if err := listEnd2(lex); err != nil {
                    return err
                }
                return nil
            }
        }
    }
    return fmt.Errorf("erro listInit2")
}

func listMult2(lex lexer.TokenProducer) error {
    if proxToken.TokenType == types.KKOMA {
        proxToken = lex.GetNextToken()
        if proxToken.TokenType == types.IDENTIFIER {
            proxToken = lex.GetNextToken()
            return nil
        }
    }
    return fmt.Errorf("erro listMult2")
}

func listEnd2(lex lexer.TokenProducer) error {
    for isInFirst(LIST_MULT, proxToken.TokenType) {
        if err := listMult2(lex); err != nil {
            return err
        }
    }
    if proxToken.TokenType == types.SEMICOLON {
        proxToken = lex.GetNextToken()
        return nil
    }
    return fmt.Errorf("erro listEnd2")
}

func listSeq2(lex lexer.TokenProducer) error {
    for isInFirst(LIST_INIT, proxToken.TokenType) {
        if err := listInit2(lex); err != nil {
            return err
        }
    }
    return nil
}

func cmd2(lex lexer.TokenProducer) error {
	if isInFirst(CMD_SEL, proxToken.TokenType) {
		if err := cmdSel2(lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_REP, proxToken.TokenType) {
		if err := cmdRep2(lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(CMD_ATR, proxToken.TokenType) {
		if err := cmdAtr2(lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro cmd2")
}

func cmdSeq2(lex lexer.TokenProducer) error {
	for isInFirst(CMD, proxToken.TokenType) {
		if err := cmd2(lex); err != nil {
			return err
		}
	}
	return nil
}

func cmdOrBlock2(lex lexer.TokenProducer) error {
	if isInFirst(CMD, proxToken.TokenType) {
		if err := cmd2(lex); err != nil {
			return err
		}
		return nil
	}
	if isInFirst(BLOCK, proxToken.TokenType) {
		if err := block2(lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro cmdOrBlock2 %v", proxToken)
}

func cmdSel2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_IF {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			proxToken = lex.GetNextToken()
			if err := cond2(lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.KW_THEN {
					proxToken = lex.GetNextToken()
					if err := cmdOrBlock2(lex); err != nil {
						return err
					}
					if err := cmdSelPrime2(lex); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	return fmt.Errorf("erro cmdSel2")
}

func cmdSelPrime2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_ELSE {
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock2(lex); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func cmdRep2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.KW_WHILE {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.START_PAREN {
			proxToken = lex.GetNextToken()
			if err := cond2(lex); err != nil {
				return err
			}
			if proxToken.TokenType == types.END_PAREN {
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.KW_DO {
                    proxToken = lex.GetNextToken()
					if err := cmdOrBlock2(lex); err != nil {
						return err
					}
					return nil
				}
			}
		}
	}
	if proxToken.TokenType == types.KW_REPEAT {
		proxToken = lex.GetNextToken()
		if err := cmdOrBlock2(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.KW_UNTIL {
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.START_PAREN {
				proxToken = lex.GetNextToken()
				if err := cond2(lex); err != nil {
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

func cmdAtr2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.IDENTIFIER {
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.ASSIGN {
			proxToken = lex.GetNextToken()
			if err := exp2(lex); err != nil {
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

func exp2(lex lexer.TokenProducer) error {
	if err := term2(lex); err != nil {
		return err
	}
	if err := expPrime2(lex); err != nil {
		return err
	}
	return nil
}

func expPrime2(lex lexer.TokenProducer) error {
	for proxToken.TokenType == types.ARIOP_SUM {
		proxToken = lex.GetNextToken()
		if err := term2(lex); err != nil {
			return err
		}
	}
	return nil
}

func term2(lex lexer.TokenProducer) error {
	if err := factor2(lex); err != nil {
		return err
	}
	if err := termPrime2(lex); err != nil {
		return err
	}
	return nil
}

func termPrime2(lex lexer.TokenProducer) error {
	for proxToken.TokenType == types.ARIOP_MULT {
		proxToken = lex.GetNextToken()
		if err := factor2(lex); err != nil {
			return err
		}
	}
	return nil
}

func factor2(lex lexer.TokenProducer) error {
	if proxToken.TokenType == types.START_PAREN {
		proxToken = lex.GetNextToken()
		if err := exp2(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			proxToken = lex.GetNextToken()
			if err := factorPrime2(lex); err != nil {
                return err
                            }
            return nil
		}
	}
	if proxToken.TokenType == types.CONST {
		proxToken = lex.GetNextToken()
		if err := factorPrime2(lex); err != nil {
			return err
		}
		return nil
	}
	if proxToken.TokenType == types.IDENTIFIER {
		proxToken = lex.GetNextToken()
		if err := factorPrime2(lex); err != nil {
			return err
		}
		return nil
	}
	return fmt.Errorf("erro factor2")
}

func factorPrime2(lex lexer.TokenProducer) error {
	for proxToken.TokenType == types.ARIOP_POW {
		proxToken = lex.GetNextToken()
        fmt.Println("foi ariop_pow")
		if err := factor2(lex); err != nil {
			return err
		}
	}
	return nil
}

func cond2(lex lexer.TokenProducer) error {
	if isInFirst(EXP, proxToken.TokenType) {
		if err := exp2(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.RELOP {
			proxToken = lex.GetNextToken()
			if err := exp2(lex); err != nil {
				return err
			}
			return nil
		}
	}
	if proxToken.TokenType == types.START_PAREN {
		proxToken = lex.GetNextToken()
		if err := cond2(lex); err != nil {
			return err
		}
		if proxToken.TokenType == types.END_PAREN {
			proxToken = lex.GetNextToken()
			if proxToken.TokenType == types.RELOP {
				proxToken = lex.GetNextToken()
				if proxToken.TokenType == types.START_PAREN {
					proxToken = lex.GetNextToken()
					if err := cond2(lex); err != nil {
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
