package lexer_test

import (
	"strings"
	"testing"

	"github.com/heitorfreitasferreira/compiler/lexer"
	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

func TestPositionMoving(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  []types.Position
	}{
		{
			desc: "sem espaco",
			in:   "((((((((",
			out: []types.Position{
				{Line: 1, Column: 1},
				{Line: 1, Column: 2},
				{Line: 1, Column: 3},
				{Line: 1, Column: 4},
				{Line: 1, Column: 5},
				{Line: 1, Column: 6},
				{Line: 1, Column: 7},
				{Line: 1, Column: 8},
			},
		},
		{
			desc: "com identificador antes",
			in:   "var ((",
			out: []types.Position{
				{Line: 1, Column: 1},
				{Line: 1, Column: 5},
				{Line: 1, Column: 6},
			},
		},
		{
			desc: "com separador no meio",
			in:   "a\nb\n12\n",
			out: []types.Position{
				{Line: 1, Column: 1},
				{Line: 2, Column: 1},
				{Line: 3, Column: 1},
			},
		},
		{
			desc: "com separador e token com lexema com len > 1",
			in:   "var a\nb\n12\nint : v, u;\n",
			out: []types.Position{
				{Line: 1, Column: 1},  // var
				{Line: 1, Column: 5},  // a
				{Line: 2, Column: 1},  // b
				{Line: 3, Column: 1},  // 12
				{Line: 4, Column: 1},  // int
				{Line: 4, Column: 5},  // :
				{Line: 4, Column: 7},  // v
				{Line: 4, Column: 8},  // ,
				{Line: 4, Column: 10}, // u
				{Line: 4, Column: 11}, // ;
			},
		},
		{
			desc: "tokens com lexemas grandes",
			in:   "begin int: variavel_muito_grande; end\n",
			out: []types.Position{
				{Line: 1, Column: 1},  // begin
				{Line: 1, Column: 7},  // int
				{Line: 1, Column: 10}, // :
				{Line: 1, Column: 12}, // variavel_muito_grande
				{Line: 1, Column: 33}, // ;
				{Line: 1, Column: 35}, // end
			},
		},
		{
			desc: "declaracao de lista de variaveis inteiras",
			in:   "int: v, u;\n",
			out: []types.Position{
				{Line: 1, Column: 1},  // int
				{Line: 1, Column: 4},  // :
				{Line: 1, Column: 6},  // v
				{Line: 1, Column: 7},  // ,
				{Line: 1, Column: 9},  // u
				{Line: 1, Column: 10}, // ;
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lex := lexer.NewLexer(
				strings.NewReader(tC.in),
				simboltable.NewSymbolTable(),
				statemachine.DefaultDFA,
			)

			got := make([]types.Position, 0)

			for {
				tok := lex.GetNextToken()
				if tok.TokenType == types.EOF {
					break
				}
				got = append(got, tok.Position)
			}

			if len(got) != len(tC.out) {
				t.Fatalf("Test case: %s, era esperado reconhecer %d tokens, mas foram reconhecidos %d", tC.desc, len(tC.out), len(got))
			}

			for i := range got {
				if got[i].Column != tC.out[i].Column || got[i].Line != tC.out[i].Line {
					t.Errorf("Test case: %s, era esperado %v, mas foi reconhecido %v", tC.desc, tC.out[i], got[i])
				}
			}
		})
	}
}
