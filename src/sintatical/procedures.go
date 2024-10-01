package sintatical

import (
	"fmt"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/types"
)

func s(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	currNode.Value = S
	currNode.Children = []*types.Node[types.GrammarSymbol]{}

	if proxToken.TokenType == types.KW_MAIN {
		currNode.AddNode(proxToken)
		proxToken = lex.GetNextToken()
		if proxToken.TokenType == types.IDENTIFIER {
			currNode.AddNode(proxToken)
			proxToken = lex.GetNextToken()
			block(currNode, lex)
			return nil
		}
		return fmt.Errorf("expected identifier at %v", proxToken.Position)
	}
	return fmt.Errorf("expected 'main' at %v", proxToken.Position)
}

func block(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	appendDown(currNode, BLOCK)
	if proxToken.TokenType == types.KW_BEGIN {
		proxToken = lex.GetNextToken()
		currNode.AddNode(proxToken)
		decl(currNode, lex)
		cmdSeq(currNode, lex)
		if proxToken.TokenType == types.KW_END {
			currNode.AddNode(proxToken)
			return nil
		}
		return fmt.Errorf("expected 'end' at %v", proxToken.Position)
	}
	return fmt.Errorf("expected 'begin' at %v", proxToken.Position)
}

func decl(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdSeq(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func list(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmd(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdOrBlock(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdSel(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdSelSimp(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdSelComp(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdRep(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cmdAtr(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func exp(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func expPrime(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func term(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func termPrime(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func factor(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func factorPrime(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}

func cond(currNode *types.Node[types.GrammarSymbol], lex *lexer.Lexer) error {
	return nil
}
