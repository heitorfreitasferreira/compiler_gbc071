package sintatical

type NonTerminal string

const (
	S          NonTerminal = "S"
	BLOCK      NonTerminal = "Block"
	DECL       NonTerminal = "Decl"
	CMDSEQ     NonTerminal = "CmdSeq"
	LIST       NonTerminal = "List"
	CMD        NonTerminal = "Cmd"
	CMDORBLOCK NonTerminal = "CmdOrBlock"
	CMDSEL     NonTerminal = "CmdSel"
	CMDSELSIMP NonTerminal = "CmdSelSimp"
	CMDSELCOMP NonTerminal = "CmdSelComp"
	CMDREP     NonTerminal = "CmdRep"
	CMDATR     NonTerminal = "CmdAtr"
	EXP        NonTerminal = "Exp"
	EXP1       NonTerminal = "Exp'"
	TERM       NonTerminal = "Term"
	TERM1      NonTerminal = "Term'"
	FACTOR     NonTerminal = "Factor"
	COND       NonTerminal = "Cond"
)

func (nt NonTerminal) String() string {
	return string(nt)
}
