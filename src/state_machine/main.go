package statemachine

import "github.com/heitorfreitasferreira/compiler/types"

// TODO: Quando o diagrama final estiver pronto, implementa-lo aqui
var DefaultDFA *DFA = NewDFA(
	[][]int{
		GetTransition([]types.Tuple[byte, int]{{'<', 1}, {'>', 4}, {'=', 7}, {'!', 9}}),
		GetTransition([]types.Tuple[byte, int]{{'=', 3}}, types.Tuple[[]byte, int]{[]byte{'='}, 2}),
		GetTransition([]types.Tuple[byte, int]{}),
		GetTransition([]types.Tuple[byte, int]{}),
		GetTransition([]types.Tuple[byte, int]{{'=', 6}}, types.Tuple[[]byte, int]{[]byte{'='}, 5}),
		GetTransition([]types.Tuple[byte, int]{}),
		GetTransition([]types.Tuple[byte, int]{}),
		GetTransition([]types.Tuple[byte, int]{{'=', 8}}),
		GetTransition([]types.Tuple[byte, int]{}),
		GetTransition([]types.Tuple[byte, int]{{'=', 10}}),
		GetTransition([]types.Tuple[byte, int]{}),
	},
	map[int]types.Tuple[types.TokenType, bool]{
		2:  {types.RELOP, true},
		3:  {types.RELOP, false},
		5:  {types.RELOP, true},
		6:  {types.RELOP, false},
		8:  {types.RELOP, false},
		10: {types.RELOP, false},
	},
)
