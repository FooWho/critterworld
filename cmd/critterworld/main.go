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
	p := parser.NewParser(tokens)
	program, err := p.Parse()
	if err != nil {
		fmt.Printf("Got error: %v", err)
	}

	ast := parser.NewAbstractSyntaxTree(program)
	nodes := ast.GetNodes()
	for i := 0; i < len(nodes); i++ {
		fmt.Printf("%t\n", nodes[i])
	}

	//lexer = parser.NewLexer("{1 = 1 or 2 = 2} and 5 = 3 --> mem[0] := 3 * (5 + 3);")
	//tokens, err = lexer.Tokenize()
	//p = parser.NewParser(tokens)
	//program, err = p.Parse()
	//if err != nil {
	//		fmt.Printf("Got error: %v", err)
	//}
	//if program != nil {
	//		fmt.Print(program)
	//}

}
