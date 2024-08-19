package lexer

import (
	"bufio"

	"github.com/heitorfreitasferreira/compiler/types"
)

type Lexer struct {
	types.Position
	*bufio.Reader
}

func NewLexer(reader *bufio.Reader) *Lexer {
	return &Lexer{
		Reader: reader,
		Position: types.Position{
			Line:   1,
			Column: 0,
		},
	}
}

func (l *Lexer) GetNextToken() (types.Token, error) {
	return types.Token{}, nil
}

func (l *Lexer) peak() (rune, error) {
	bts, err := l.Reader.Peek(1)
	if err != nil {
		return ' ', err
	}
	return rune(bts[0]), nil
}
