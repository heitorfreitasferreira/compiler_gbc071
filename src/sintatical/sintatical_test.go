package sintatical_test

import (
	"testing"

	"github.com/heitorfreitasferreira/compiler/sintatical"
	"github.com/heitorfreitasferreira/compiler/types"
)

type lexerMock struct {
	tokens []types.Token
}

func newLexerMock(tts ...types.TokenType) *lexerMock {
	tokenList := make([]types.Token, len(tts))
	for i, tt := range tts {
		tokenList[i] = types.Token{
			TokenType: tt,
		}
	}
	return &lexerMock{
		tokens: tokenList,
	}
}

func (lex *lexerMock) GetNextToken() types.Token {
	token := lex.tokens[0]
	lex.tokens = lex.tokens[1:]
	return token
}

func Test(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []types.TokenType
		expected sintatical.ConcreteSintaticalTree
	}{
		{
			desc:  "input5",
			input: []types.TokenType{types.KW_MAIN, types.IDENTIFIER, types.KW_END, types.EOF},
			expected: sintatical.ConcreteSintaticalTree{
				Root: &types.Node[types.GrammarSymbol]{
					Value: sintatical.S,
					Children: []*types.Node[types.GrammarSymbol]{
						{Value: types.Token{TokenType: types.KW_MAIN}},
						{Value: types.Token{TokenType: types.IDENTIFIER}},
						{Value: sintatical.BLOCK, Children: []*types.Node[types.GrammarSymbol]{
							{Value: types.Token{TokenType: types.KW_BEGIN}},
							{Value: types.Token{TokenType: types.KW_END}},
						}},
						{Value: types.Token{TokenType: types.EOF}},
					},
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			lexer := newLexerMock(tC.input...)
			sin := sintatical.Sintatical{
				Lexer: lexer,
			}
			tree, err := sin.Analize()
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if !tree.IsEqual(&tC.expected) {
				t.Fatalf("expected %v, got %v", tC.expected, tree)
			}
		})
	}
}
