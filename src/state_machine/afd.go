package statemachine

import (
	"errors"

	"github.com/heitorfreitasferreira/compiler/types"
)

var (
	TransitionNotSupported error = errors.New("empty transition not supported")
	DealWithLookAheadError error = errors.New("this error is treated by the lexer by calling DealWithLookAhead")
)

type AFD struct {
	current       int
	prev          int
	states        []map[byte]int                      // Cada estado é representado por um índice no slice
	final         []bool                              // Lista de booleanos que indicam se o estado é final
	cleanUp       []func(string) (types.Token, error) // Funções para limpar o estado, como por exemplo, tratar lookahead ou desfazer a leitura de um byte
	lexemeBuilder []byte
}

func NewAFD(states []map[byte]int, finals map[int]func(string) (types.Token, error)) *AFD {
	final := make([]bool, len(states))
	cleanUp := make([]func(string) (types.Token, error), len(states))
	for i, f := range finals {
		final[i] = true
		cleanUp[i] = f
	}

	return &AFD{
		current:       0, // Estado inicial sempre é 0
		prev:          0,
		states:        states,
		final:         final,
		cleanUp:       cleanUp,
		lexemeBuilder: make([]byte, 0),
	}
}

func (afd *AFD) Step(b byte) (*types.Token, error) {
	next, ok := afd.states[afd.current][b]
	ch := string(b)
	_ = ch
	if !ok {
		return nil, TransitionNotSupported
	}
	afd.prev = afd.current
	afd.current = next

	afd.lexemeBuilder = append(afd.lexemeBuilder, b)

	if afd.final[afd.current] {
		lexeme := string(afd.lexemeBuilder)
		tk, err := afd.cleanUp[afd.current](lexeme)
		if err == DealWithLookAheadError {
			// Remove last byte from tk.Atr["lexeme"]
			tk.Lexeme = tk.Lexeme[:len(tk.Lexeme)-1]
		}
		afd.lexemeBuilder = make([]byte, 0)
		afd.current = 0
		return &tk, err
	}
	return nil, nil
}

func (afd *AFD) GoBackOneState() {
	afd.current = afd.prev
}
