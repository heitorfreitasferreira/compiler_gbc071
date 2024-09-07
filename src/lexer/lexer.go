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
	lastCol       int
	startOfLexeme types.Position
	peaking       bool
	peakingByte   byte

	*statemachine.DFA
	types.Position
	io.ByteReader
	*simboltable.SymbolTable
}

func NewLexer(reader io.ByteReader, st *simboltable.SymbolTable, dfa *statemachine.DFA) *Lexer {
	return &Lexer{
		ByteReader:  reader,
		SymbolTable: st,
		DFA:         dfa,
		Position: types.Position{
			Line:   1,
			Column: 1,
		},
		startOfLexeme: types.Position{
			Line:   1,
			Column: 1,
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
		panic(fmt.Errorf("error reading file: %v", err))
	}

	l.peakingByte = b
	token, lookAhead, err := l.DFA.Step(b)
	if err == statemachine.ErrTransitionNotSupported {
		fmt.Printf("%v [ %s ]\nLinha: %d\nColuna: %d\n", err, string(b), l.Position.Line, l.Position.Column)
		os.Exit(1)
	}
	if token == nil || token.TokenType == types.COMMENT || token.TokenType == types.SEPARATOR {
		if token != nil {
			l.assignPosition(token)
		}
		return l.GetNextToken()
	}
	l.updateSimbolTable(token)
	if lookAhead {
		l.dealWithLookAhead()
	}
	l.assignPosition(token)
	return *token
}

func (l *Lexer) dealWithLookAhead() {
	if l.peakingByte == '\n' {
		l.Position.Line--
		l.Position.Column = l.lastCol
	} else {
		l.Position.Column--
	}
	l.peaking = true
}

func (l *Lexer) updateSimbolTable(tk *types.Token) {
	if tk.TokenType == types.IDENTIFIER || tk.TokenType == types.CONST {
		tk.Id = new(int)
		*tk.Id = -1
		*tk.Id = l.SymbolTable.AddSymbol(tk.Lexeme)
		return
	}
}

func (l *Lexer) read() (byte, error) {
	if l.peaking {
		l.peaking = false
		l.updatePosition(l.peakingByte)
		return l.peakingByte, nil
	}
	b, err := l.ByteReader.ReadByte()
	l.updatePosition(b)
	return b, err
}

func (l *Lexer) assignPosition(tk *types.Token) {
	tk.Position.Line = l.startOfLexeme.Line
	tk.Position.Column = l.startOfLexeme.Column

	l.startOfLexeme.Line = l.Position.Line
	l.startOfLexeme.Column = l.Position.Column
}

func (l *Lexer) updatePosition(b byte) {
	if b == '\n' {
		l.Position.Line++
		l.Position.Column = 1
	} else {
		l.Position.Column++
	}
}
