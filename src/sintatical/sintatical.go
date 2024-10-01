package sintatical

import (
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

	return tree, err
}
