package sintatical

import (
	"github.com/heitorfreitasferreira/compiler/types"
)

type ConcreteSintaticalTree types.Tree[types.GrammarSymbol]

func appendDown(currNode *types.Node[types.GrammarSymbol], tk NonTerminal) {

	oldNode := currNode
	currNode = &types.Node[types.GrammarSymbol]{
		Value:    tk,
		Children: []*types.Node[types.GrammarSymbol]{},
	}
	oldNode.Children = append(oldNode.Children, currNode)
}
