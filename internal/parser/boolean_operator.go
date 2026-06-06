package parser

import (
	"fmt"
)

type BooleanOperator interface {
	ASTNode
	SwapOperands()
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
	return []ASTNode{lo.leftOperand, lo.rightOperand}
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

func (lo *LogicalOperator) String() string {
	var str string
	if lo.breakingPrecedence(lo.leftOperand) {
		str += fmt.Sprintf("{%s}", lo.leftOperand)
	} else {
		str += fmt.Sprintf("%s", lo.leftOperand)
	}
	str += " " + lo.operator.Lexeme + " "
	if lo.breakingPrecedence(lo.rightOperand) {
		str += fmt.Sprintf("{%s}", lo.rightOperand)
	} else {
		str += fmt.Sprintf("%s", lo.rightOperand)
	}
	return str
}

func (lo *LogicalOperator) breakingPrecedence(operand BooleanOperator) bool {
	if operand.NodeType() == "LogicalOperator" &&
		lo.operator.TokenType == tAnd &&
		operand.(*LogicalOperator).operator.TokenType == tOr {
		return true
	}
	return false
}

func (lo *LogicalOperator) SwapOperands() {
	tmp := lo.leftOperand
	lo.leftOperand = lo.rightOperand
	lo.rightOperand = tmp
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
	return []ASTNode{ro.leftOperand, ro.rightOperand}
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

func (ro *RelationalOperator) String() string {
	return fmt.Sprintf("%s %s %s", ro.leftOperand, ro.operator, ro.rightOperand)
}

func (ro *RelationalOperator) IsBooleanOperator() bool {
	return true
}

func (ro *RelationalOperator) SwapOperands() {
	tmp := ro.leftOperand
	ro.leftOperand = ro.rightOperand
	ro.rightOperand = tmp
}

// Interface guard
var _ BooleanOperator = (*RelationalOperator)(nil)
var _ ASTNode = (*RelationalOperator)(nil)
