package lexer

import (
	"fmt"
	"io"
	"os"

	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

type Lexer struct {
	lastPosition *types.Position
	peaking      bool
	peakingByte  byte

	*statemachine.DFA
	*types.Position
	io.ByteReader
	*simboltable.SymbolTable
}

func NewLexer(reader io.ByteReader, st *simboltable.SymbolTable, dfa *statemachine.DFA) *Lexer {
	return &Lexer{
		ByteReader:  reader,
		SymbolTable: st,
		DFA:         dfa,
		Position: &types.Position{
			Line:   1,
			Column: 0,
		},
		lastPosition: &types.Position{
			Line:   1,
			Column: 0,
		},
	}
}

func (l *Lexer) GetNextToken() types.Token {
	b, err := l.read()
	if err == io.EOF {
		return types.Token{
			TokenType: types.EOF,
		}
	}
	if err != nil {
		panic(err)
	}

	var token *types.Token
	var lookAhead bool
	token, lookAhead, err = l.DFA.Step(b)
	if err == statemachine.ErrTransitionNotSupported {
		fmt.Printf("Caractere não suportado: [ %s ]\nLinha: %d\nColuna: %d\n", string(b), l.Position.Line, l.Position.Column)
		os.Exit(1)
	}
	if token == nil || token.TokenType == types.COMMENT || token.TokenType == types.SEPARATOR {
		return l.GetNextToken()
	}
	l.updateSimbolTable(token)
	if lookAhead {
		l.dealWithLookAhead()
	}
	return *token
}

func (l *Lexer) dealWithLookAhead() {
	l.Position.Column = l.lastPosition.Column
	l.Position.Line = l.lastPosition.Line

	l.peaking = true
}

func (l *Lexer) updateSimbolTable(tk *types.Token) {
	if tk.TokenType == types.IDENTIFIER || tk.TokenType == types.NUM_CONST {
		*tk.Id = l.SymbolTable.AddSymbol(tk.Lexeme)
		return
	}
}

func (l *Lexer) read() (byte, error) {
	if l.peaking {
		l.peaking = false
		return l.peakingByte, nil
	}
	b, err := l.ByteReader.ReadByte()
	if err != nil {
		return b, err
	}
	if b == '\n' {
		l.lastPosition.Column = l.Position.Column
		l.lastPosition.Line = l.Position.Line
		l.Position.Line++
		l.Position.Column = 0
	} else {
		l.lastPosition.Column = l.Position.Column
		l.Position.Column++
	}
	return l.ByteReader.ReadByte()
}
