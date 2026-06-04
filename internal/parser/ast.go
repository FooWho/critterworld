package parser

type AbstractSyntaxTree struct {
	rootNode  *Program
	nodeCount int
}

type ASTNode interface {
	NodeType() string
	Children() []ASTNode
	Clone() ASTNode
}
