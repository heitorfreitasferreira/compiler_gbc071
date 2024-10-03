package sintatical

type NonTerminal string

const (
	S             NonTerminal = "S"
	BLOCK         NonTerminal = "Block"
	LIST          NonTerminal = "List"
	LIST_PRIME    NonTerminal = "List'"
	DECL_SEQ      NonTerminal = "DeclSeq"
	DECL          NonTerminal = "Decl"
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
	COND          NonTerminal = "Cond"
)

func (nt NonTerminal) String() string {
	return string(nt)
}
