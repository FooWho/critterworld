package main

import (
	"fmt"

	"github.com/FooWho/critterworld/internal/parser"
)

func main() {
	header, source := parser.ReadCritterSource("critter1.crtr")
	fmt.Println("======================================")
	fmt.Print(header)
	fmt.Println("======================================")
	fmt.Print(source)
	fmt.Println("======================================")

	lexer := parser.NewLexer(source)

	for {
		token := lexer.NextToken()
		if token == nil {
			break
		}
		fmt.Printf("Token Type: %-12s | Value: %s\n", token.Type, token.Lexeme)
	}
}
