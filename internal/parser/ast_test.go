package parser

import (
	"testing"
)

func TestAbstractSyntaxTree_GetNodesOfType(t *testing.T) {
	source := "1 = 1 --> mem[1] := 2 eat;"

	lexer := NewLexer(source)
	tokens, err := lexer.Tokenize()
	if err != nil {
		t.Fatalf("Unable to tokenize source in TestAbstractSyntaxTree_GetNodesOfType(): %v", err)
	}
	parser := NewParser(tokens)
	program, err := parser.Parse()
	if err != nil {
		t.Fatalf("Unable to parse source in TestAbstractSyntaxTree_GetNodesOfType(): %v", err)
	}

	ast := NewAbstractSyntaxTree(program)
	numberNodes := GetNodesOfType[*Number](&ast)

	// Should have four *Number nodes: '1', '1', '1', and '2'
	if len(numberNodes) != 4 {
		t.Errorf("Expected 4 *Number nodes, got %d", len(numberNodes))
	}
	if numberNodes[0].value != 1 || numberNodes[1].value != 1 || numberNodes[2].value != 1 || numberNodes[3].value != 2 {
		t.Errorf("Expected 1, 1, 1, 2 values for *Number nodes, got %d, %d, %d, %d",
			numberNodes[0].value, numberNodes[1].value, numberNodes[2].value, numberNodes[3].value)
	}

	memNodes := GetNodesOfType[*MemNode](&ast)
	// Should have one
	if len(memNodes) != 1 {
		t.Errorf("Expected 1 *MemNode node, got %d", len(memNodes))
	}
	if memNodes[0].String() != "mem[1]" {
		t.Errorf("Expected node 'mem[1], got %s", memNodes[0].String())
	}

	booleanNodes := GetNodesOfType[BooleanOperator](&ast)
	if len(booleanNodes) != 1 {
		t.Errorf("Expected 1 BooleanOperator node, got %d", len(booleanNodes))
	}
	if booleanNodes[0].String() != "1 = 1" {
		t.Errorf("Expected BooleanOperator '1 = 1', got %s", booleanNodes[0].String())
	}

}
