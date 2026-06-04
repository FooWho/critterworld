package parser

import (
	"fmt"
)

type BooleanOperator interface {
	ASTNode
	Evaluate() bool
	IsBooleanOperator() bool
}

type LogicalOperator struct {
	operation    LexedToken
	leftOperand  Expression
	rightOperand Expression
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
	cloneLO := LogicalOperator{}

	clonedOperand := lo.leftOperand.Clone()
	cloneLO.leftOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (lo *LogicalOperator).Clone(), got %T", clonedOperand))
	}
	clonedOperand = lo.rightOperand.Clone()
	cloneLO.rightOperand = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (lo *LogicalOperator).Clone(), got %T", clonedOperand))
	}
	cloneLO.rightOperand, ok = clonedOperand.(Expression)

	return &cloneLO
}

func (lo *LogicalOperator) Evaluate() bool {
	return true
}

func (lo *LogicalOperator) IsBooleanOperator() bool {
	return true
}

type RelationalOperator struct {
	operation    LexedToken
	rightOperand Expression
	leftOperand  Expression
}

func (ro *RelationalOperator) NodeType() string {
	return "LogicalOperator"
}

func (ro *RelationalOperator) Children() []ASTNode {
	children := make([]ASTNode, 2)
	children[0] = ro.leftOperand
	children[1] = ro.rightOperand
	return children
}

func (ro *RelationalOperator) Clone() ASTNode {
	var ok bool
	cloneRO := LogicalOperator{}

	clonedOperand := ro.leftOperand.Clone()
	cloneRO.leftOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (ro *RelationalOperator).Clone(), got %T", clonedOperand))
	}
	clonedOperand = ro.rightOperand.Clone()
	cloneRO.rightOperand = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (ro *RelationalOperator).Clone(), got %T", clonedOperand))
	}
	cloneRO.rightOperand, ok = clonedOperand.(Expression)

	return &cloneRO
}

func (ro *RelationalOperator) Evaluate() bool {
	return true
}

func (ro *RelationalOperator) IsBooleanOperator() bool {
	return true
}
