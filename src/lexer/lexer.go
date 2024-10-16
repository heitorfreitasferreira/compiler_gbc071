package lexer

import (
	"fmt"
	"io"
	"os"

	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

type TokenProducer interface {
	GetNextToken() types.Token
}
type Lexer struct {
	lastCol       int
	startOfLexeme types.Position
	peaking       bool
	peakingByte   byte

	dfa      *statemachine.DFA
	position types.Position
	br       io.ByteReader
	st       *simboltable.SymbolTable
}

func NewLexer(reader io.ByteReader, st *simboltable.SymbolTable, dfa *statemachine.DFA) *Lexer {
	return &Lexer{
		br:  reader,
		st:  st,
		dfa: dfa,
		position: types.Position{
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
	token, lookAhead, err := l.dfa.Step(b)
	if err == statemachine.ErrTransitionNotSupported {
		fmt.Printf("%v [ %s ]\nLinha: %d\nColuna: %d\n", err, string(b), l.position.Line, l.position.Column)
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

	if token.Id != nil {
		fmt.Printf("ID: %v, Lexema: %v\n", *token.Id, token.Lexeme)
	} else {
		fmt.Printf("ID: nil, Lexema: %v\n", token.Lexeme)
	}

	return *token
}

func (l *Lexer) dealWithLookAhead() {
	if l.peakingByte == '\n' {
		l.position.Line--
		l.position.Column = l.lastCol
	} else {
		l.position.Column--
	}
	l.peaking = true
}

func (l *Lexer) updateSimbolTable(tk *types.Token) {
	if tk.TokenType == types.IDENTIFIER || tk.TokenType == types.CONST {
		tk.Id = new(int)
		*tk.Id = -1
		*tk.Id = l.st.AddSymbol(tk.Lexeme)
		return
	}
}

func (l *Lexer) read() (byte, error) {
	if l.peaking {
		l.peaking = false
		l.updatePosition(l.peakingByte)
		return l.peakingByte, nil
	}
	b, err := l.br.ReadByte()
	l.updatePosition(b)
	return b, err
}

func (l *Lexer) assignPosition(tk *types.Token) {
	tk.Position.Line = l.startOfLexeme.Line
	tk.Position.Column = l.startOfLexeme.Column

	l.startOfLexeme.Line = l.position.Line
	l.startOfLexeme.Column = l.position.Column
}

func (l *Lexer) updatePosition(b byte) {
	if b == '\n' {
		l.position.Line++
		l.position.Column = 1
	} else {
		l.position.Column++
	}
}
