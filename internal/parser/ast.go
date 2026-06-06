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

func NewAbstractSyntaxTree(root *Program) AbstractSyntaxTree {
	ast := AbstractSyntaxTree{rootNode: root}
	return ast
}

func (ast *AbstractSyntaxTree) GetNodes() []ASTNode {
	var nodes []ASTNode
	var toVisit []ASTNode

	toVisit = append(toVisit, ast.rootNode)
	for len(toVisit) > 0 {
		n := len(toVisit) - 1
		currentNode := toVisit[n]
		toVisit = toVisit[:n]

		if currentNode == nil {
			continue
		}

		nodes = append(nodes, currentNode)
		children := currentNode.Children()
		for i := len(children) - 1; i >= 0; i-- {
			toVisit = append(toVisit, children[i])
		}
	}
	return nodes
}

func (ast *AbstractSyntaxTree) GetNodesByInterface() []ASTNode {
	candidates := ast.GetNodes()
	for _, node := range candidates {
		if 
	}
}

