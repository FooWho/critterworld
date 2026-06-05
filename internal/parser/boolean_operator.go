package parser

import (
	"fmt"
)

type BooleanOperator interface {
	ASTNode
	IsBooleanOperator() bool
}

type LogicalOperator struct {
	operator     LexedToken
	leftOperand  BooleanOperator
	rightOperand BooleanOperator
}

func (lo *LogicalOperator) NodeType() string {
	return "LogicalOperator"
}

func (lo *LogicalOperator) Children() []ASTNode {
	children := make([]ASTNode, 2)
	children[0] = lo.leftOperand
	children[1] = lo.rightOperand
	return children
}

func (lo *LogicalOperator) Clone() ASTNode {
	var ok bool
	cloneLO := LogicalOperator{operator: lo.operator}

	clonedOperand := lo.leftOperand.Clone()
	cloneLO.leftOperand, ok = clonedOperand.(BooleanOperator)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected BooleanOperator in (lo *LogicalOperator).Clone(), got %T", clonedOperand))
	}
	clonedOperand = lo.rightOperand.Clone()
	cloneLO.rightOperand, ok = clonedOperand.(BooleanOperator)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected BooleanOperator in (lo *LogicalOperator).Clone(), got %T", clonedOperand))
	}

	return &cloneLO
}

func (lo *LogicalOperator) IsBooleanOperator() bool {
	return true
}

// Interface guard
var _ BooleanOperator = (*LogicalOperator)(nil)
var _ ASTNode = (*LogicalOperator)(nil)

type RelationalOperator struct {
	operator     LexedToken
	rightOperand Expression
	leftOperand  Expression
}

func (ro *RelationalOperator) NodeType() string {
	return "RelationalOperator"
}

func (ro *RelationalOperator) Children() []ASTNode {
	children := make([]ASTNode, 2)
	children[0] = ro.leftOperand
	children[1] = ro.rightOperand
	return children
}

func (ro *RelationalOperator) Clone() ASTNode {
	var ok bool
	cloneRO := RelationalOperator{operator: ro.operator}

	clonedOperand := ro.leftOperand.Clone()
	cloneRO.leftOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (ro *RelationalOperator).Clone(), got %T", clonedOperand))
	}
	clonedOperand = ro.rightOperand.Clone()
	cloneRO.rightOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (ro *RelationalOperator).Clone(), got %T", clonedOperand))
	}

	return &cloneRO
}

func (ro *RelationalOperator) IsBooleanOperator() bool {
	return true
}

// Interface guard
var _ BooleanOperator = (*RelationalOperator)(nil)
var _ ASTNode = (*RelationalOperator)(nil)
