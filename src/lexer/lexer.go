package lexer

import (
	"bufio"

	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

type Lexer struct {
	charsRead    int
	lastByteRead byte
	statemachine.DFA
	*types.Position
	lastPosition *types.Position
	*bufio.Reader

	*simboltable.SymbolTable
}

func NewLexer(reader *bufio.Reader, st *simboltable.SymbolTable) *Lexer {
	return &Lexer{
		Reader:      reader,
		SymbolTable: st,
		charsRead:   0,
		Position: &types.Position{
			Line:   1,
			Column: 0,
		},
	}
}

func (l *Lexer) GetNextToken() (*types.Token, error) {
	return nil, nil // TODO: Implement
}

func (l *Lexer) DealWithLookAhead() error {
	l.charsRead--
	l.Position.Column = l.lastPosition.Column
	l.Position.Line = l.lastPosition.Line
	err := l.Reader.UnreadRune()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lexer) peak() (rune, error) {
	bts, err := l.Reader.Peek(1)
	if err != nil {
		return ' ', err
	}
	return rune(bts[0]), nil
}
