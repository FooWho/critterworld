package parser

import (
	"fmt"
)

type Expression interface {
	ASTNode
	IsExpression() bool
}

type BinaryOperator struct {
	operator     LexedToken
	leftOperand  Expression
	rightOperand Expression
}

func (bo *BinaryOperator) NodeType() string {
	return "BinaryOperator"
}

func (bo *BinaryOperator) Children() []ASTNode {
	children := make([]ASTNode, 2)
	children[0] = bo.leftOperand
	children[1] = bo.rightOperand
	return children
}

func (bo *BinaryOperator) Clone() ASTNode {
	var ok bool
	clonedBO := BinaryOperator{operator: bo.operator}

	clonedOperand := bo.leftOperand.Clone()
	clonedBO.leftOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (bo *BinaryOperator).Clone(), got %T", clonedOperand))
	}

	clonedOperand = bo.rightOperand.Clone()
	clonedBO.rightOperand = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (bo *BinaryOperator).Clone(), got %T", clonedOperand))
	}

	return &clonedBO
}

func (bo *BinaryOperator) IsExpression() bool {
	return true
}

// Interface guard
var _ Expression = (*BinaryOperator)(nil)
var _ ASTNode = (*BinaryOperator)(nil)

type UnaryOperator struct {
	operator LexedToken
	operand  Expression
}

func (uo *UnaryOperator) NodeType() string {
	return "UnaryOperator"
}

func (uo *UnaryOperator) Children() []ASTNode {
	children := make([]ASTNode, 1)
	children[0] = uo.operand
	return children
}

func (uo *UnaryOperator) Clone() ASTNode {
	var ok bool
	clonedUO := UnaryOperator{operator: uo.operator}

	clonedOperand := uo.operand.Clone()
	clonedUO.operand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (uo *UnaryOperator).Clone(), got %T", clonedOperand))
	}

	return &clonedUO
}

func (uo *UnaryOperator) IsExpression() bool {
	return true
}

// Interface guard
var _ Expression = (*UnaryOperator)(nil)
var _ ASTNode = (*UnaryOperator)(nil)

type MemNode struct {
	operand Expression
}

func (mn *MemNode) NodeType() string {
	return "MemNode"
}

func (mn *MemNode) Children() []ASTNode {
	children := make([]ASTNode, 1)
	children[0] = mn.operand
	return children
}

func (mn *MemNode) Clone() ASTNode {
	var ok bool
	clonedMN := MemNode{}

	clonedOperand := mn.operand.Clone()
	clonedMN.operand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (mn *MemNode).Clone(), got %T", clonedOperand))
	}

	return &clonedMN
}

func (mn *MemNode) IsExpression() bool {
	return true
}

// Interface guard
var _ Expression = (*MemNode)(nil)
var _ ASTNode = (*MemNode)(nil)

type Number struct {
	value int
}

func (n *Number) NodeType() string {
	return "Number"
}

func (n *Number) Children() []ASTNode {
	return nil
}

func (n *Number) Clone() ASTNode {
	clonedN := Number{value: n.value}

	return &clonedN
}

func (n *Number) IsExpression() bool {
	return true
}

// Interface guard
var _ Expression = (*Number)(nil)
var _ ASTNode = (*Number)(nil)
