package statemachine

import (
	"testing"

	"github.com/heitorfreitasferreira/compiler/types"
)

func TestDeterministicWithAssign(t *testing.T) {
	input := []byte(":=")

	transitions := [][]int{
		GetTransition([]types.Tuple[byte, int]{{':', 1}}),
		GetTransition([]types.Tuple[byte, int]{{'=', 2}}),
		GetTransition([]types.Tuple[byte, int]{}),
	}
	finals := map[int]types.Tuple[types.TokenType, bool]{
		2: {types.ASSIGN, false},
	}
	dfa := NewDFA(
		transitions,
		finals,
	)

	shouldBeNil, lookAhead, err := dfa.Step(input[0])
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if shouldBeNil != nil {
		t.Errorf("Expected token to be nil, got %v", shouldBeNil)
	}
	if lookAhead {
		t.Errorf("Expected lookAhead to be false, got %v", lookAhead)
	}
	if dfa.currentState != 1 {
		t.Errorf("Expected current state to be 1, got %d", dfa.currentState)
	}

	shouldBeAssign, lookAhead, err := dfa.Step(input[1])
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if shouldBeAssign.TokenType != types.ASSIGN {
		t.Errorf("Expected token to be ASSIGN, got %v", shouldBeAssign)
	}
	if lookAhead {
		t.Errorf("Expected lookAhead to be false, got %v", lookAhead)
	}
	if dfa.currentState != 0 {
		t.Errorf("Expected current state to be 0, got %d", dfa.currentState)
	}
}

func TestDealingWithLookAhead(t *testing.T) {
	input := []byte(">=<!===")
	expectedLexemes := []string{">=", "<", "!=", "=="}
	gotLexemes := make([]string, 0)
	transitions := [][]int{
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
	}
	finals := map[int]types.Tuple[types.TokenType, bool]{
		2:  {types.RELOP, true},
		3:  {types.RELOP, false},
		5:  {types.RELOP, true},
		6:  {types.RELOP, false},
		8:  {types.RELOP, false},
		10: {types.RELOP, false},
	}
	dfa := NewDFA(
		transitions,
		finals,
	)
	for i := 0; i < len(input); i++ {
		tk, lookAhead, err := dfa.Step(input[i])
		if err != nil {
			t.Fatalf("Expected no error, got %v", err)
		}
		if lookAhead {
			i--
		}
		if tk != nil {
			gotLexemes = append(gotLexemes, tk.Lexeme)
		}
	}

	if len(gotLexemes) != len(expectedLexemes) {
		t.Errorf("Expected %d lexemes, got %d", len(expectedLexemes), len(gotLexemes))
	}

	for lx := range expectedLexemes {
		if gotLexemes[lx] != expectedLexemes[lx] {
			t.Errorf("Expected lexeme %s, got %s", expectedLexemes[lx], gotLexemes[lx])
		}
	}
}
