package parser

import (
	"fmt"
)

type AbstractSyntaxTree struct {
	rootNode  *Program
	nodeCount int
}

type ASTNode interface {
	fmt.Stringer
	NodeType() string
	Children() []ASTNode
	Clone() ASTNode
}
