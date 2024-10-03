package sintatical

import (
	"fmt"
	"strings"

	"github.com/heitorfreitasferreira/compiler/types"
)

type ConcreteSintaticalTree types.Tree[types.GrammarSymbol]

func (t ConcreteSintaticalTree) stringify(n *types.Node[types.GrammarSymbol], depth int) string {
	if n == nil {
		return ""
	}
	result := fmt.Sprintf("%s%v\n", strings.Repeat("  ", depth), n.Value)
	for _, child := range n.Children {
		result += t.stringify(child, depth+1)
	}
	return result
}

func (t ConcreteSintaticalTree) String() string {
	return t.stringify(t.Root, 0)
}

func (cst ConcreteSintaticalTree) IsEqual(other *ConcreteSintaticalTree) bool {
	return cst.Root.IsEqual(other.Root)
}
