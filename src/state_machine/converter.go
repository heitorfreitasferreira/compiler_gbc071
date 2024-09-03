package statemachine

import (
	"github.com/heitorfreitasferreira/compiler/types"
)

func ConvertToDFAWithoutEpsilon(dfa *DFA) *DFA {
	newStates := [][]int{}
	finalStates := map[int]types.Tuple[types.TokenType, bool]{}
	stateMapping := map[string]int{}
	stateQueue := [][]int{}

	initialClosure := dfa.epsilon_closure([]int{0})

	// Adicionar o estado inicial ao novo DFA
	stateQueue = append(stateQueue, initialClosure)
	stateMapping[encodeStateSet(initialClosure)] = 0
	newStates = append(newStates, make([]int, len(globalAlphabet)))
	nextStateIndex := 1

	for len(stateQueue) > 0 {
		currentStateSet := stateQueue[0]
		stateQueue = stateQueue[1:]

		for symbol := range globalAlphabet {
			if symbol == emptyTransition {
				continue // Ignorar a transição vazia
			}

			newStateSet := dfa.epsilon_closure(dfa.move(symbol))

			if len(newStateSet) == 0 {
				continue // Nenhum novo estado foi encontrado
			}

			newStateKey := encodeStateSet(newStateSet)

			// Verificar se o novo conjunto de estados já foi adicionado ao DFA
			if _, exists := stateMapping[newStateKey]; !exists {
				stateMapping[newStateKey] = nextStateIndex
				stateQueue = append(stateQueue, newStateSet)
				newStates = append(newStates, make([]int, len(globalAlphabet)))
				nextStateIndex++
			}

			// Adicionar a transição ao novo DFA
			newStates[stateMapping[encodeStateSet(currentStateSet)]][symbol] = stateMapping[newStateKey]
		}
	}

	// Identificar os estados finais
	for encodedStateSet, newStateIndex := range stateMapping {
		originalStates := decodeStateSet(encodedStateSet)
		for _, originalState := range originalStates {
			if dfa.final[originalState] {
				finalStates[newStateIndex] = dfa.out[originalState]
				break // Estado final encontrado no estado resultantes
			}
		}
	}
	return NewDFA(newStates, finalStates)
}

func (dfa DFA) epsilon_closure(states []int) []int {
	closure := make(map[int]bool)
	var explore func(int)
	explore = func(state int) {
		if !closure[state] {
			closure[state] = true
			if dfa.states[state][emptyTransition] != notInAlphabet {
				explore(dfa.states[state][emptyTransition])
			}
		}
	}

	for _, state := range states {
		explore(state)
	}

	closureList := make([]int, 0, len(closure))
	for state := range closure {
		closureList = append(closureList, state)
	}
	return closureList
}

func (dfa DFA) move(c byte) []int {

	move := make([]int, 0)
	for state := 0; state < len(dfa.states); state++ {
		if dfa.states[state][c] != notInAlphabet {
			move = append(move, dfa.states[state][c])
		}
	}

	return move
}
