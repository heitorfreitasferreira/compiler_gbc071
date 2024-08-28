package statemachine

import (
	"errors"
	"strings"

	"github.com/heitorfreitasferreira/compiler/types"
)

var (
	ErrTransitionNotSupported error = errors.New("unexpected character")
	ErrDealWithLookAhead      error = errors.New("this error is treated by the lexer by calling DealWithLookAhead")
)

type DFA struct {
	currentState  int
	states        [][]int                              // Cada estado é representado por um índice no slice
	final         []bool                               // Lista de booleanos que indicam se o estado é final
	out           []types.Tuple[types.TokenType, bool] // Lista de tokens e se deve tratar look ahead nos estados finais
	lexemeBuilder strings.Builder
}

func NewDFA(states [][]int, finals map[int]types.Tuple[types.TokenType, bool]) *DFA {
	final := make([]bool, len(states))
	out := make([]types.Tuple[types.TokenType, bool], len(states))
	for i, f := range finals {
		final[i] = true
		out[i] = f
	}

	return &DFA{
		currentState:  0,
		states:        states,
		final:         final,
		out:           out,
		lexemeBuilder: strings.Builder{},
	}
}

func (dfa *DFA) Step(transition byte) (*types.Token, bool, error) {
	next := dfa.states[dfa.currentState][transition]
	if next == notInAlphabet {
		return nil, false, ErrTransitionNotSupported
	}

	dfa.currentState = next

	dfa.lexemeBuilder.WriteByte(transition)

	if dfa.final[dfa.currentState] {
		token, lookAhead := dfa.out[dfa.currentState].First, dfa.out[dfa.currentState].Second
		var lexeme string
		if lookAhead {
			lexeme = dfa.lexemeBuilder.String()[:len(dfa.lexemeBuilder.String())-1]
		} else {
			lexeme = dfa.lexemeBuilder.String()
		}
		dfa.lexemeBuilder.Reset()
		dfa.currentState = 0
		return &types.Token{
			TokenType: token,
			Lexeme:    lexeme,
			Position:  types.Position{},
			Id:        nil,
		}, lookAhead, nil
	}
	return nil, false, nil
}
