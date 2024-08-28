package simboltable

import (
	"testing"

	"github.com/heitorfreitasferreira/compiler/types"
)

func TestMultipleInsertsOfTheSameTOKEN_ID(t *testing.T) {

}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		in       []types.Token
		expected map[string]int
	}{
		{
			desc: "same identifier multiple times should result in only one entry in the symbol table",
			in: []types.Token{
				{TokenType: types.IDENTIFIER, Lexeme: "a"},
				{TokenType: types.IDENTIFIER, Lexeme: "a"},
				{TokenType: types.IDENTIFIER, Lexeme: "a"},
				{TokenType: types.IDENTIFIER, Lexeme: "a"},
				{TokenType: types.IDENTIFIER, Lexeme: "a"},
			},
			expected: map[string]int{
				"a": 0,
			},
		},
		{
			desc: "different identifiers should result in different entries in the symbol table",
			in: []types.Token{
				{TokenType: types.IDENTIFIER, Lexeme: "a"},
				{TokenType: types.IDENTIFIER, Lexeme: "b"},
				{TokenType: types.IDENTIFIER, Lexeme: "c"},
			},
			expected: map[string]int{
				"a": 0,
				"b": 1,
				"c": 2,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			st := NewSymbolTable()
			for _, token := range tC.in {
				st.AddSymbol(token.Lexeme)
			}
			for lexeme, id := range st.table {
				if id != tC.expected[lexeme] {
					t.Errorf("Expected %s to have id %d, got %d", lexeme, tC.expected[lexeme], id)
				}
			}
		})
	}
}
