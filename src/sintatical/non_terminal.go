package sintatical

type NonTerminal string

const (
	S             NonTerminal = "S"
	BLOCK         NonTerminal = "Block"
	LIST_INIT     NonTerminal = "ListInit"
	LIST_MULT     NonTerminal = "ListMult"
	LIST_END      NonTerminal = "ListEnd"
	LIST_SEQ      NonTerminal = "ListSeq"
	CMD           NonTerminal = "Cmd"
	CMD_SEQ       NonTerminal = "CmdSeq"
	CMD_OR_BLOCK  NonTerminal = "CmdOrBlock"
	CMD_SEL       NonTerminal = "CmdSel"
	CMD_SEL_PRIME NonTerminal = "CmdSel'"
	CMD_REP       NonTerminal = "CmdRep"
	CMD_ATR       NonTerminal = "CmdAtr"
	EXP           NonTerminal = "Exp"
	EXP_PRIME     NonTerminal = "Exp'"
	TERM          NonTerminal = "Term"
	TERM_PRIME    NonTerminal = "Term'"
	FACTOR        NonTerminal = "Factor"
	FACTOR_PRIME  NonTerminal = "Factor'"
    UNARY_EXP      NonTerminal = "UnaryExp"
    ID_OR_CONST     NonTerminal = "IdOrConst"
	COND          NonTerminal = "Cond"
)

func (nt NonTerminal) String() string {
	return string(nt)
}
