package sintatical

import (
	"fmt"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

var proxToken types.Token

type Sintatical struct {
	Lexer lexer.TokenProducer
}

func (sin *Sintatical) Analize() (ConcreteSintaticalTree, error) {
	tree := ConcreteSintaticalTree{
		Root: &types.Node[types.GrammarSymbol]{},
	}

	proxToken = sin.Lexer.GetNextToken()

	err := s(tree.Root, sin.Lexer)
	if err != nil {
		return tree, err
	}
	if proxToken.TokenType != types.EOF {
		return tree, fmt.Errorf("expected EOF at %v", proxToken.Position)
	}
	return tree, nil
}
