package statemachine

import (
	"testing"

	"github.com/heitorfreitasferreira/compiler/types"
)

func TestDeterministicWithAssign(t *testing.T) {
	input := []byte(":=")
	dfa := NewDFA(
		[]map[byte]int{
			map[byte]int{':': 1},
			map[byte]int{'=': 2},
			map[byte]int{},
		},
		map[int]func(string) (types.Token, error){
			2: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.ASSIGN,
				}, nil
			},
		},
	)
	shouldBeNil, err := dfa.Step(input[0])

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if shouldBeNil != nil {
		t.Errorf("Expected token to be nil, got %v", shouldBeNil)
	}
	if dfa.current != 1 {
		t.Errorf("Expected current state to be 1, got %d", dfa.current)
	}
	if string(dfa.lexemeBuilder) != ":" {
		t.Errorf("Expected lexemeBuilder to be ':', got %s", string(dfa.lexemeBuilder))
	}

	shouldBeAssign, err := dfa.Step(input[1])
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if shouldBeAssign.TokenType != types.ASSIGN {
		t.Errorf("Expected token to be ASSIGN, got %v", shouldBeAssign)
	}

	if dfa.current != 0 {
		t.Errorf("Expected current state to be 0, got %d", dfa.current)
	}
	if len(dfa.lexemeBuilder) != 0 {
		t.Errorf("Expected lexemeBuilder to be empty, got %s", string(dfa.lexemeBuilder))
	}
}

func TestDealingWithLookAhead(t *testing.T) {
	input := []byte(">=<!===")
	expectedLexemes := []string{">=", "<", "!=", "=="}
	gotLexemes := make([]string, 0)

	dfa := NewDFA(
		[]map[byte]int{
			map[byte]int{'<': 1, '>': 4, '=': 7, '!': 9},      //0
			alphabetNot(map[byte]int{'=': 3}, []byte{'='}, 2), //1
			map[byte]int{}, //2
			map[byte]int{}, //3
			alphabetNot(map[byte]int{'=': 6}, []byte{'='}, 5), //4
			map[byte]int{},        //5
			map[byte]int{},        //6
			map[byte]int{'=': 8},  //7
			map[byte]int{},        //8
			map[byte]int{'=': 10}, //9
			map[byte]int{},        //10
		},
		map[int]func(string) (types.Token, error){
			2: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.RELOP,
					Lexeme:    lexeme,
				}, DealWithLookAheadError
			},
			3: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.RELOP,
					Lexeme:    lexeme,
				}, nil
			},
			5: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.RELOP,
					Lexeme:    lexeme,
				}, DealWithLookAheadError
			},
			6: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.RELOP,
					Lexeme:    lexeme,
				}, nil
			},
			8: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.RELOP,
					Lexeme:    lexeme,
				}, nil
			},
			10: func(lexeme string) (types.Token, error) {
				return types.Token{
					TokenType: types.RELOP,
					Lexeme:    lexeme,
				}, nil
			},
		},
	)

	for i := 0; i < len(input); i++ {
		tk, err := dfa.Step(input[i])
		if err == DealWithLookAheadError {
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
