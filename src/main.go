package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/heitorfreitasferreira/compiler/lexer"
	"github.com/heitorfreitasferreira/compiler/myBufferedByteReader"
	simboltable "github.com/heitorfreitasferreira/compiler/simbol_table"
	"github.com/heitorfreitasferreira/compiler/sintatical"
	statemachine "github.com/heitorfreitasferreira/compiler/state_machine"
	"github.com/heitorfreitasferreira/compiler/types"
)

func main() {
	sourceFilePath := os.Args[1:][0]

	sourceFilePath = strings.TrimSpace(sourceFilePath)
	if sourceFilePath == "" {
		fmt.Println("Please provide a file path as the first argument")
		os.Exit(1)
	}

	// Open file
	file, err := os.Open(sourceFilePath)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	defer file.Close()

	st := simboltable.NewSymbolTable()
	myBuffReader := &myBufferedByteReader.BufferedByteReader{}
	myBufferedByteReader.InitBufferedByteReader(myBuffReader, file)

	l := lexer.NewLexer(myBuffReader, st, statemachine.DefaultDFA)
	onlyLexer := false
	if onlyLexer {
		fmt.Println("Tokens:")
		for {
			token := l.GetNextToken()
			fmt.Printf("%v  ", token)
			if token.TokenType == types.EOF {
				break
			}
		}
		return
	}

	analyzer := sintatical.Sintatical{Lexer: l}
	tree, err := analyzer.Analize()
	if err != nil {
		fmt.Printf("%v\n", err)
		os.Exit(0)
	}
	fmt.Println(tree)

	// fmt.Println("\nSymbol Table:", st)
}
