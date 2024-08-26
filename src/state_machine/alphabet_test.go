package statemachine

import "testing"

func Test(t *testing.T) {

	oldTransition := map[byte]int{
		'q': 1,
	}
	transitions := addNegationTransitions(oldTransition, []byte{'q'}, 2)

	if len(transitions) == len(oldTransition) {
		t.Errorf("Expected more than 1 transition, got %d", len(transitions))
	}
}
