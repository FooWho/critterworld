package parser

import (
	"fmt"
)

type BooleanOperator interface {
	ASTNode
	swapOperands()
	isBooleanOperator()
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

func (lo *LogicalOperator) isBooleanOperator() {

}

func (lo *LogicalOperator) isASTNode() {

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
	if op, ok := operand.(*LogicalOperator); ok &&
		lo.operator.TokenType == tAnd &&
		op.operator.TokenType == tOr {
		return true
	}
	return false
}

func (lo *LogicalOperator) SwapChildren(firstChild, secondChild ASTNode) error {
	lo.swapOperands()
	return nil
}

func (lo *LogicalOperator) swapOperands() {
	lo.leftOperand, lo.rightOperand = lo.rightOperand, lo.leftOperand
}

func (lo *LogicalOperator) Transform(newValue any) error {
	newOperator, ok := newValue.(LexedToken)
	if !ok {
		return fmt.Errorf("expected LexedToken, got %T", newValue)
	}
	lo.operator = newOperator
	return nil
}

func (lo *LogicalOperator) ReplaceChild(oldChild, newChild ASTNode) error {
	newBoolOp, isBoolOp := newChild.(BooleanOperator)
	if !isBoolOp {
		fmt.Errorf("newChild is type %T, does not implement BooleanOperator")
	}
	if lo.leftOperand == oldChild {
		lo.leftOperand = newBoolOp
		return nil
	} else if lo.rightOperand == oldChild {
		lo.rightOperand = newBoolOp
		return nil
	} else {
		return fmt.Errorf("oldChild: %v not found in lo: %v", oldChild, lo)
	}
}

func (lo *LogicalOperator) RemoveChild(child ASTNode) error {
	return fmt.Errorf("LogicalOperator %v cannot remove child %v", lo, child)
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

func (ro *RelationalOperator) swapOperands() {
	ro.leftOperand, ro.rightOperand = ro.rightOperand, ro.leftOperand
}

// Transform changes the operator of the RelationalOperator in-place.
// This is a non-structural, "value-only" transform.
func (ro *RelationalOperator) Transform(newValue any) error {
	newOperator, ok := newValue.(LexedToken)
	if !ok {
		return fmt.Errorf("expected LexedToken, got %T", newValue)
	}
	ro.operator = newOperator
	return nil
}

func (ro *RelationalOperator) SwapChildren(firstChild, secondChild ASTNode) error {
	ro.swapOperands()
	return nil
}

func (ro *RelationalOperator) ReplaceChild(oldChild, newChild ASTNode) error {
	newExp, isExp := newChild.(Expression)
	if !isExp {
		fmt.Errorf("newChild is type %T, does not implement Expression")
	}
	if ro.leftOperand == oldChild {
		ro.leftOperand = newExp
		return nil
	} else if ro.rightOperand == oldChild {
		ro.rightOperand = newExp
		return nil
	} else {
		return fmt.Errorf("oldChild: %v not found in ro: %v", oldChild, ro)
	}
}

func (ro *RelationalOperator) RemoveChild(child ASTNode) error {
	return fmt.Errorf("RelationalOperator %v cannot remove child %v", ro, child)
}

func (ro *RelationalOperator) isBooleanOperator() {

}

func (ro *RelationalOperator) isASTNode() {

}

// Interface guard
var _ BooleanOperator = (*RelationalOperator)(nil)
var _ ASTNode = (*RelationalOperator)(nil)
