package statemachine

import (
	"testing"

	"github.com/heitorfreitasferreira/compiler/types"
)

func Test(t *testing.T) {

	positives := []types.Tuple[byte, int]{
		{byte('='), 1},
	}
	negatives := []types.Tuple[[]byte, int]{
		{[]byte{'='}, 2},
	}

	transition := GetTransition(positives, negatives...)

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
