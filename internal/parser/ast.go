package parser

import (
	"fmt"
)

type AbstractSyntaxTree struct {
	rootNode *Program
}

type ASTNode interface {
	fmt.Stringer
	Children() []ASTNode
	Clone() ASTNode
	RemoveChild(child ASTNode) error
	SwapChildren(firstChild, secondChild ASTNode) error
	ReplaceChild(oldChild, newChild ASTNode) error
	Transform(newValue any) error
	InsertChild(child ASTNode, location int) error
	isASTNode()
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

func GetNodesByType[T ASTNode](ast *AbstractSyntaxTree) []T {
	var matchedNodes []T
	for _, node := range ast.GetNodes() {
		if match, ok := node.(T); ok {
			matchedNodes = append(matchedNodes, match)
		}
	}
	return matchedNodes
}

func (ast *AbstractSyntaxTree) GetParentOf(target ASTNode) ASTNode {
	if target == ast.rootNode {
		return nil
	}

	var find func(current, parent ASTNode) ASTNode
	find = func(current, parent ASTNode) ASTNode {
		if current == target {
			return parent
		}
		for _, child := range current.Children() {
			if child != nil {
				if p := find(child, current); p != nil {
					return p
				}
			}
		}
		return nil
	}

	return find(ast.rootNode, nil)
}
