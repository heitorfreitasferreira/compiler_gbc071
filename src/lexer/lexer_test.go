package lexer_test

import (
	"bufio"
	"strings"
	"testing"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

func TestNewLexer(t *testing.T) {
	reader := bufio.NewReader(strings.NewReader("test"))
	lexer := lexer.NewLexer(reader)
	if lexer.Line != 1 {
		t.Errorf("Expected line to be 1, got %d", lexer.Line)
	}
	if lexer.Column != 0 {
		t.Errorf("Expected column to be 0, got %d", lexer.Column)
	}
}

func TestTokenAndAtr(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  types.Token
	}{
		{
			desc: "ARITHMETIC OPERATOR",
			in:   "+",
			out: types.Token{
				TokenType: types.ARIOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme":   "+",
					"priority": 2,
				},
			},
		},
		{
			desc: "ARITHMETIC OPERATOR",
			in:   "-",
			out: types.Token{
				TokenType: types.ARIOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme":   "-",
					"priority": 2,
				},
			},
		},
		{
			desc: "ARITHMETIC OPERATOR",
			in:   "*",
			out: types.Token{
				TokenType: types.ARIOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme":   "*",
					"priority": 1,
				},
			},
		},
		{
			desc: "ARITHMETIC OPERATOR",
			in:   "/",
			out: types.Token{
				TokenType: types.ARIOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme":   "/",
					"priority": 1,
				},
			},
		},
		{
			desc: "ARITHMETIC OPERATOR",
			in:   "**",
			out: types.Token{
				TokenType: types.ARIOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme":   "**",
					"priority": 0,
				},
			},
		},
		{
			desc: "RELATIONAL OPERATOR",
			in:   ">",
			out: types.Token{
				TokenType: types.RELOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme": ">",
				},
			},
		},
		{
			desc: "RELATIONAL OPERATOR",
			in:   "<",
			out: types.Token{
				TokenType: types.RELOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme": "<",
				},
			},
		},
		{
			desc: "RELATIONAL OPERATOR",
			in:   ">=",
			out: types.Token{
				TokenType: types.RELOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme": ">=",
				},
			},
		},
		{
			desc: "RELATIONAL OPERATOR",
			in:   "<=",
			out: types.Token{
				TokenType: types.RELOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme": "<=",
				},
			},
		},
		{
			desc: "RELATIONAL OPERATOR",
			in:   "==",
			out: types.Token{
				TokenType: types.RELOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme": "==",
				},
			},
		},
		{
			desc: "RELATIONAL OPERATOR",
			in:   "!=",
			out: types.Token{
				TokenType: types.RELOP,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"lexeme": "!=",
				},
			},
		},
		{
			desc: "ASSIGNMENT OPERATOR",
			in:   ":=",
			out: types.Token{
				TokenType: types.ASSIGN,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
			},
		},
		{
			desc: "CONSTANT LITERAL",
			in:   "\"Hello, World!\"",
			out: types.Token{
				TokenType: types.CONSTANT,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": "Hello, World!",
					"type":  "literal",
				},
			},
		},
		{
			desc: "CONSTANT INTEGER",
			in:   "123",
			out: types.Token{
				TokenType: types.CONSTANT,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": 123,
					"type":  "integer",
				},
			},
		},
		{
			desc: "CONSTANT SCIENTIFIC NOTATION",
			in:   "0.1E-2",
			out: types.Token{
				TokenType: types.CONSTANT,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": 0.1e-2,
					"type":  "float",
				},
			},
		},
		{
			desc: "CONSTANT SCIENTIFIC NOTATION",
			in:   "1.23E+10",
			out: types.Token{
				TokenType: types.CONSTANT,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": 1.23e+10,
					"type":  "float",
				},
			},
		},
		{
			desc: "CONSTANT SCIENTIFIC NOTATION",
			in:   "'a'",
			out: types.Token{
				TokenType: types.CONSTANT,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": 'a',
					"type":  "char",
				},
			},
		},
		{
			desc: "IDENTIFIER",
			in:   "v",
			out: types.Token{
				TokenType: types.IDENTIFIER,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": "v",
				},
			},
		},
		{
			desc: "IDENTIFIER",
			in:   "var",
			out: types.Token{
				TokenType: types.IDENTIFIER,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": "var",
				},
			},
		},
		{
			desc: "IDENTIFIER",
			in:   "V",
			out: types.Token{
				TokenType: types.IDENTIFIER,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": "V",
				},
			},
		},
		{
			desc: "IDENTIFIER",
			in:   "var123",
			out: types.Token{
				TokenType: types.IDENTIFIER,
				Position: types.Position{
					Line:   1,
					Column: 0,
				},
				Attr: map[string]interface{}{
					"value": "var123",
				},
			},
		},
		{
			desc: "SEPARATOR",
			in:   ",",
			out: types.Token{
				TokenType: types.TYPE_SEPARATOR,
			},
		},
		{
			desc: "SEPARATOR",
			in:   ";",
			out: types.Token{
				TokenType: types.EOF,
			},
		},
		{
			desc: "TYPE SEPARATOR",
			in:   ":",
			out: types.Token{
				TokenType: types.TYPE_SEPARATOR,
			},
		},
		{
			desc: "SEPARATOR",
			in:   "\n",
			out: types.Token{
				TokenType: types.EOF,
			},
		},
		{
			desc: "SEPARATOR",
			in:   "\t",
			out: types.Token{
				TokenType: types.EOF,
			},
		},
		{
			desc: "SEPARATOR",
			in:   "",
			out: types.Token{
				TokenType: types.EOF,
			},
		},
		{
			desc: "COMMENT",
			in:   "{ This is a comment }",
			out: types.Token{
				TokenType: types.EOF,
			},
		},
		{
			desc: "KEYWORD",
			in:   "if",
			out: types.Token{
				TokenType: types.KEYWORD,
				Attr: map[string]interface{}{
					"lexeme": "if",
					"type":   "conditional",
				},
			},
		},
		{
			desc: "KEYWORD",
			in:   "else",
			out: types.Token{
				TokenType: types.KEYWORD,
				Attr: map[string]interface{}{
					"lexeme": "else",
					"type":   "conditional",
				},
			},
		},
		{
			desc: "KEYWORD",
			in:   "while",
			out: types.Token{
				TokenType: types.KEYWORD,
				Attr: map[string]interface{}{
					"lexeme": "while",
					"type":   "loop",
				},
			},
		},
		{
			desc: "KEYWORD",
			in:   "do",
			out: types.Token{
				TokenType: types.KEYWORD,
				Attr: map[string]interface{}{
					"lexeme": "do",
					"type":   "loop",
				},
			},
		},
		{
			desc: "KEYWORD",
			in:   "repeat",
			out: types.Token{
				TokenType: types.KEYWORD,
				Attr: map[string]interface{}{
					"lexeme": "repeat",
					"type":   "loop",
				},
			},
		},
		{
			desc: "KEYWORD",
			in:   "until",
			out: types.Token{
				TokenType: types.KEYWORD,
				Attr: map[string]interface{}{
					"lexeme": "until",
					"type":   "loop",
				},
			},
		},
		{
			desc: "TYPE",
			in:   "int",
			out: types.Token{
				TokenType: types.TYPE,
				Attr: map[string]interface{}{
					"lexeme": "int",
				},
			},
		},
		{
			desc: "TYPE",
			in:   "float",
			out: types.Token{
				TokenType: types.TYPE,
				Attr: map[string]interface{}{
					"lexeme": "float",
				},
			},
		},
		{
			desc: "TYPE",
			in:   "char",
			out: types.Token{
				TokenType: types.TYPE,
				Attr: map[string]interface{}{
					"lexeme": "char",
				},
			},
		},
		{
			desc: "EOF",
			in:   "",
			out: types.Token{
				TokenType: types.EOF,
			},
		},
		{
			desc: "START BLOCK",
			in:   "begin",
			out: types.Token{
				TokenType: types.START_BLOCK,
			},
		},
		{
			desc: "END BLOCK",
			in:   "end",
			out: types.Token{
				TokenType: types.END_BLOCK,
			},
		},
		{
			desc: "START PAREN",
			in:   "(",
			out: types.Token{
				TokenType: types.START_PAREN,
			},
		},
		{
			desc: "END PAREN",
			in:   ")",
			out: types.Token{
				TokenType: types.END_PAREN,
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := lexer.NewLexer(bufio.NewReader(strings.NewReader(tC.in)))
			tk, err := l.GetNextToken()
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if tk.TokenType != tC.out.TokenType {
				t.Errorf("Expected token type to be %s, got %s", tC.out.TokenType, tk.TokenType)
			}
		})
	}
}

func TestNeverReturnCommentsAndSeparators(t *testing.T) {
	testCases := []struct {
		desc string
		in   string
		out  []types.Token
	}{
		{
			desc: "main",
			in:   "  \t\tmain\n\t\t { This is a comment } \n\t\t begin \n\t\t int:4;\n end",
			out: []types.Token{
				{
					TokenType: types.IDENTIFIER,
				},
				{
					TokenType: types.START_BLOCK,
				},
				{
					TokenType: types.TYPE,
				},
				{
					TokenType: types.TYPE_SEPARATOR,
				},
				{
					TokenType: types.CONSTANT,
				},
				{
					TokenType: types.TYPE_SEPARATOR,
				},
				{
					TokenType: types.END_BLOCK,
				},
				{
					TokenType: types.EOF,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			l := lexer.NewLexer(bufio.NewReader(strings.NewReader(tC.in)))
			for _, out := range tC.out {
				tk, err := l.GetNextToken()
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if tk.TokenType != out.TokenType {
					t.Errorf("Expected token type to be %s, got %s", out.TokenType, tk.TokenType)
				}
			}
		})
	}
}
