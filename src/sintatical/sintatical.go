package sintatical

import (
	"fmt"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

var proxToken types.Token

type Sintatical struct {
	lex *lexer.Lexer
}

func (sin *Sintatical) Analize() (ConcreteSintaticalTree, error) {
	tree := ConcreteSintaticalTree{
		Root: &types.Node[types.GrammarSymbol]{},
	}

	proxToken = sin.lex.GetNextToken()

	err := s(tree.Root, sin.lex)
	if err != nil {
		return tree, err
	}
	proxToken = sin.lex.GetNextToken()
	if proxToken.TokenType != types.EOF {
		return tree, fmt.Errorf("expected EOF at %v", proxToken.Position)
	}
	return tree, nil
}
