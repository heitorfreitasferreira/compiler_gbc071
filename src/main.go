package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/heitorfreitasferreira/compiler/lexer"
	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

func main() {
	sourceFilePath := os.Args[1:][0]

	if sourceFilePath == "" {
		panic("Please provide a file path as the first argument")
	}

	// Open file
	file, err := os.Open(sourceFilePath)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	st := simboltable.NewSymbolTable()
	l := lexer.NewLexer(bufio.NewReader(file), st, statemachine.DefaultDFA)
	fmt.Println("Tokens:")
	for {
		token := l.GetNextToken()
		if token.TokenType == types.EOF {
			break
		}
		fmt.Println(token)
	}

	fmt.Println("\nSymbol Table:")
	fmt.Println(st)
}
