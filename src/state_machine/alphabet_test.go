package statemachine

import (
	"testing"

	"github.com/heitorfreitasferreira/compiler/types"
)

func Test(t *testing.T) {

	positives := []types.Tuple[byte, int]{
		{byte('='), 1},
	}

	transition := GetTransition(positives, 2)

	for letter := range globalAlphabet {
		if letter == '=' {
			if transition[int(letter)] != 1 {
				t.Errorf("Expected transition to 1, got %d", transition[int(letter)])
			}
		} else {
			if transition[int(letter)] != 2 {
				t.Errorf("Expected transition to 2, got %d", transition[int(letter)])
			}
		}
	}
}

func TestEmptyTransitionNotInGlobalAlphabet(t *testing.T) {
	if _, exists := globalAlphabet[emptyTransition]; exists {
		t.Errorf("Empty transition should not be in global alphabet")
	}
}

func TestLetterTransition(t *testing.T) {
	tr := GetTransitionLetterDigit([]types.Tuple[[]byte, int]{
		{Letter, 1},
	})

	for i, transition := range tr {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' {
			if transition != 1 {
				t.Errorf("Expected transition to 1, got %d", transition)
			}
		} else {
			if transition != -1 {
				t.Errorf("Expected transition to 0, got %d", transition)
			}
		}
	}
}
