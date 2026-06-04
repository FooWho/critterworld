package parser

import (
	"fmt"
)

type Expression interface {
	ASTNode
	Evaluate() int
	isExpression() bool
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
	cloneBO := BinaryOperator{operator: bo.operator}

	clonedOperand := bo.leftOperand.Clone()
	cloneBO.leftOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (bo *BinaryOperator).Clone(), got %T", clonedOperand))
	}

	clonedOperand = bo.rightOperand.Clone()
	cloneBO.rightOperand = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (bo *BinaryOperator).Clone(), got %T", clonedOperand))
	}

	return &cloneBO
}

func (bo *BinaryOperator) Evaluate() int {
	return 0
}

func (bo *BinaryOperator) IsBooleanOperator() bool {
	return true
}
