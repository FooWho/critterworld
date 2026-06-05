package main

import (
	"fmt"
	"log"

	"github.com/FooWho/critterworld/internal/parser"
)

func main() {
	header, source, err := parser.ReadCritterSource("critter1.crtr")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("======================================")
	fmt.Print(header)
	fmt.Println("======================================")
	fmt.Print(source)
	fmt.Println("======================================")

	lexer := parser.NewLexer(source)
	tokens, err := lexer.Tokenize()
	if err != nil {
		fmt.Printf("Error is: %v\n", err)
	}
	for _, token := range tokens {

		fmt.Printf("Token Type: %-12s | Value: %s\n", token.TokenType, token.Lexeme)
	}

	lexer = parser.NewLexer("mem[0] = 1")
	tokens, err = lexer.Tokenize()
	p := parser.NewParser(tokens)
	program := p.Parse()
	fmt.Println(program)

}
