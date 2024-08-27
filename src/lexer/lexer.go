package lexer

import (
	"bufio"
	"fmt"
	"io"
	"os"

	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

type Lexer struct {
	lastPosition *types.Position

	*statemachine.DFA
	*types.Position
	*bufio.Reader
	*simboltable.SymbolTable
}

func NewLexer(reader *bufio.Reader, st *simboltable.SymbolTable, dfa *statemachine.DFA) *Lexer {
	return &Lexer{
		Reader:      reader,
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
	b := make([]byte, 1)
	n, err := l.Reader.Read(b)
	ch := string(b)
	_ = ch
	if n == 0 || err == io.EOF {
		return types.Token{
			TokenType: types.EOF,
		}
	}
	if err != nil {
		panic(err)
	}

	l.updatePosition(b[0])
	var token *types.Token
	var lookAhead bool
	token, lookAhead, err = l.DFA.Step(b[0])
	if err == statemachine.ErrTransitionNotSupported {
		fmt.Printf("Caractere não suportado: [ %s ]\nLinha: %d\nColuna: %d\n", ch, l.Position.Line, l.Position.Column)
		os.Exit(1)
	}
	if token == nil || token.TokenType == types.COMMENT || token.TokenType == types.SEPARATOR {
		return l.GetNextToken()
	}

	l.updateSimbolTable(token)
	if lookAhead {
		l.dealWithLookAhead()
	}
	// change the reader to next token

	token.Position.Column = l.lastPosition.Column
	token.Position.Line = l.lastPosition.Line
	return *token
}

func (l *Lexer) DealWithLookAhead() error {
	l.Position.Column = l.lastPosition.Column
	l.Position.Line = l.lastPosition.Line
	err := l.Reader.UnreadRune()
	if err != nil {
		return err
	}
	return nil
}

func (l *Lexer) updatePosition(b byte) {
	if b == '\n' {
		l.lastPosition.Column = l.Position.Column
		l.lastPosition.Line = l.Position.Line
		l.Position.Line++
		l.Position.Column = 0
	} else {
		l.lastPosition.Column = l.Position.Column
		l.Position.Column++
	}
}

func (l *Lexer) dealWithLookAhead() {
	l.lastPosition.Column = l.Position.Column
	l.lastPosition.Line = l.Position.Line

	// Undo the last read in the reader
	err := l.Reader.UnreadByte()
	if err != nil {
		panic("erro ao desfazer a leitura devido ao look ahead")
	}
}

func (l *Lexer) updateSimbolTable(tk *types.Token) {
	if tk.TokenType == types.IDENTIFIER || tk.TokenType == types.NUM_CONST {
		tk.Attr[simboltable.ST_KEY] = l.SymbolTable.AddSymbol(tk.Lexeme)
		return
	}
}
