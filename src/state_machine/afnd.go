package statemachine

import "fmt"

type stateAFND struct {
	transitions map[*byte][]*stateAFND
	epsilon     []*stateAFND
	isFinal     bool
}

type AFND struct {
	states        []*stateAFND
	initialState  *stateAFND
	currentStates []*stateAFND
	finalStates   []*stateAFND
}

func newAFNDState(aceita bool) *stateAFND {
	return &stateAFND{
		transitions: make(map[*byte][]*stateAFND),
		epsilon:     []*stateAFND{},
		isFinal:     aceita,
	}
}

func (e *stateAFND) addTransition(simbolo *byte, estado *stateAFND) {
	if _, existe := alphabet[*simbolo]; !existe {
		panic(fmt.Sprintf("Simbolo %c não pertence ao alfabeto", *simbolo))
	}
	e.transitions[simbolo] = append(e.transitions[simbolo], estado)
}

func (e *stateAFND) addAlphabetTransition(estado *stateAFND) {
	aIndex := 97
	AIndex := 65
	for delta := range 26 {
		e.addTransition(&[]byte{byte(aIndex + delta)}[0], estado)
		e.addTransition(&[]byte{byte(AIndex + delta)}[0], estado)
	}
}

func (e *stateAFND) addDigitTransitions(state *stateAFND) {
	for delta := range 10 {
		e.addTransition(&[]byte{byte(48 + delta)}[0], state)
	}
}

func (e *stateAFND) addDigitsTransitions(state *stateAFND) {
	interm := newAFNDState(false)
	e.addDigitTransitions(interm)
	interm.addDigitTransitions(state)
}

func (e *stateAFND) addEmptyTransition(estado *stateAFND) {
	e.epsilon = append(e.epsilon, estado)
}

func NewAFND(firstState *stateAFND) *AFND {
	afnd := &AFND{
		states:       []*stateAFND{firstState},
		initialState: firstState,
	}
	afnd.resetCurrentStates()
	return afnd
}

func (a *AFND) resetCurrentStates() {
	a.currentStates = []*stateAFND{}
	a.expandEpsilon(a.initialState)
}

func (a *AFND) expandEpsilon(estado *stateAFND) {
	a.currentStates = append(a.currentStates, estado)
	for _, proxEstado := range estado.epsilon {
		a.expandEpsilon(proxEstado)
	}
}

func (a *AFND) Read(b byte) {
	novosEstados := []*stateAFND{}
	// Para cada estado atual
	for _, estado := range a.currentStates {
		// Se a transição existir
		if transicoes, existe := estado.transitions[&b]; existe {
			// Adiciona o proximo estado e expande as transições epsilon
			for _, proxEstado := range transicoes {
				novosEstados = append(novosEstados, proxEstado)
				a.expandEpsilon(proxEstado)
			}
		}
	}
	a.currentStates = novosEstados
}

// Verifica se algum dos estados atuais é um estado de aceitação
func (a *AFND) IsInFinalState() bool {
	for _, estado := range a.currentStates {
		if estado.isFinal {
			return true
		}
	}
	return false
}
func (a *AFND) Merge(other *AFND) {
	newStart := newAFNDState(false)
	newStart.addEmptyTransition(a.initialState)
	newStart.addEmptyTransition(other.initialState)
	a.initialState = newStart

}

func (a *AFND) FechoEpsilon(estados []*stateAFND) []*stateAFND {
	pilha := []*stateAFND{}
	fecho := map[*stateAFND]bool{}

	for _, t := range estados {
		pilha = append(pilha, t)
		fecho[t] = true
	}

	for len(pilha) > 0 {
		t := pilha[len(pilha)-1]
		pilha = pilha[:len(pilha)-1]

		// Para cada estado s acessível por transição epsilon
		for _, s := range t.epsilon {
			if !fecho[s] { // Se s não está no fecho
				fecho[s] = true          // Adiciona s ao fecho
				pilha = append(pilha, s) // Empilha s
			}
		}
	}

	resultado := []*stateAFND{}
	for estado := range fecho {
		resultado = append(resultado, estado)
	}

	return resultado
}

func move(states []*stateAFND, b byte) []*stateAFND {
	result := []*stateAFND{}
	for _, estado := range states {
		if transicoes, existe := estado.transitions[&b]; existe {
			for _, proxEstado := range transicoes {
				result = append(result, proxEstado)
			}
		}
	}
	return result
}
