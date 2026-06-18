package parser

import (
	"fmt"
)

type Expression interface {
	ASTNode
	isExpression()
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
	return []ASTNode{bo.leftOperand, bo.rightOperand}
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
	clonedBO.rightOperand, ok = clonedOperand.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (bo *BinaryOperator).Clone(), got %T", clonedOperand))
	}

	return &clonedBO
}

func (bo *BinaryOperator) isExpression() {
}

func (bo *BinaryOperator) String() string {
	var str string
	if bo.breakingPrecedence(bo.leftOperand) {
		str += fmt.Sprintf("(%s)", bo.leftOperand)
	} else {
		str += fmt.Sprintf("%s", bo.leftOperand)
	}
	str += " " + bo.operator.Lexeme + " "
	if bo.breakingPrecedence(bo.rightOperand) {
		str += fmt.Sprintf("(%s)", bo.rightOperand)
	} else {
		str += fmt.Sprintf("%s", bo.rightOperand)
	}

	return str
}

func (bo *BinaryOperator) breakingPrecedence(operand Expression) bool {
	if op, ok := operand.(*BinaryOperator); ok &&
		(bo.operator.TokenType == tStar ||
			bo.operator.TokenType == tDiv ||
			bo.operator.TokenType == tMod) &&
		(op.operator.TokenType == tPlus ||
			op.operator.TokenType == tMinus) {
		return true
	}
	return false
}

func (bo *BinaryOperator) SwapChildren(firstChild, secondChild ASTNode) error {
	bo.SwapOperands()
	return nil
}

func (bo *BinaryOperator) SwapOperands() {
	bo.leftOperand, bo.rightOperand = bo.rightOperand, bo.leftOperand
}

func (bo *BinaryOperator) Transform(newValue any) error {
	newOperator, ok := newValue.(LexedToken)
	if !ok {
		return fmt.Errorf("expected LexedToken, got %T", newValue)
	}
	bo.operator = newOperator
	return nil
}

func (bo *BinaryOperator) ReplaceChild(oldChild, newChild ASTNode) error {
	newExp, isExp := newChild.(Expression)
	if !isExp {
		return fmt.Errorf("newChild does not implement Expression")
	}
	if bo.leftOperand == oldChild {
		bo.leftOperand = newExp
		return nil
	} else if bo.rightOperand == oldChild {
		bo.rightOperand = newExp
		return nil
	} else {
		return fmt.Errorf("oldChild %v not found in bo %v", oldChild, bo)
	}
}

func (bo *BinaryOperator) RemoveChild(child ASTNode) error {
	return fmt.Errorf("bo %v cannot remove child %v", bo, child)
}

func (bo *BinaryOperator) InsertChild(child ASTNode, location int) error {
	return fmt.Errorf("BinaryOperator cannot perform insert child operation")
}

func (bo *BinaryOperator) isASTNode() {

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
	return []ASTNode{uo.operand}
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

func (uo *UnaryOperator) ReplaceChild(oldChild, newChild ASTNode) error {
	newExp, isExp := newChild.(Expression)
	if !isExp {
		return fmt.Errorf("newChild is %T, does not implement Expression", newChild)
	}
	if uo.operand == oldChild {
		uo.operand = newExp
		return nil
	} else {
		return fmt.Errorf("oldChild %v not found in uo: %v", oldChild, uo)
	}
}

func (uo *UnaryOperator) RemoveChild(child ASTNode) error {
	return fmt.Errorf("UnaryOperator %v cannot remove child %v", uo, child)
}

func (uo *UnaryOperator) SwapChildren(firstChild, secondChild ASTNode) error {
	return fmt.Errorf("UnaryOperator cannot swap children")
}

func (uo *UnaryOperator) Transform(newValue any) error {
	return fmt.Errorf("UnaryOperator cannot be transformed")
}

func (uo *UnaryOperator) InsertChild(child ASTNode, location int) error {
	return fmt.Errorf("UnararyOperator cannot perform insert child operation")
}

func (uo *UnaryOperator) isExpression() {
}

func (uo *UnaryOperator) isASTNode() {
}

func (uo *UnaryOperator) String() string {
	return fmt.Sprintf("%s%s", uo.operator, uo.operand)
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
	return []ASTNode{mn.operand}
}

func (mn *MemNode) Clone() ASTNode {
	return &MemNode{operand: mn.operand.Clone().(Expression)}
}

func (mn *MemNode) ReplaceChild(oldChild, newChild ASTNode) error {
	return nil
}

func (mn *MemNode) RemoveChild(child ASTNode) error {
	return fmt.Errorf("MemNode %v cannot remove operand %v", mn, child)
}

func (mn *MemNode) SwapChildren(firstChild, secondChild ASTNode) error {
	return fmt.Errorf("MemNode cannot swap children")
}

func (mn *MemNode) Transform(newValue any) error {
	return fmt.Errorf("MemNode cannot be transformed")
}

func (mn *MemNode) InsertChild(child ASTNode, location int) error {
	return fmt.Errorf("MemNode cannot perform insert child operation")
}

func (mn *MemNode) isExpression() {

}

func (mn *MemNode) isASTNode() {

}

func (mn *MemNode) String() string {
	return fmt.Sprintf("mem[%s]", mn.operand)
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

func (n *Number) ReplaceChild(oldChild, newChild ASTNode) error {
	return fmt.Errorf("Number nodes do not have children")
}

func (n *Number) RemoveChild(child ASTNode) error {
	return fmt.Errorf("Number %v cannot remove child %v", n, child)
}

func (n *Number) SwapChildren(firstChild, secondChild ASTNode) error {
	return fmt.Errorf("Number cannot swap children")
}

func (n *Number) Transform(newValue any) error {
	newVal, ok := newValue.(int)
	if !ok {
		return fmt.Errorf("expected int, got %T", newValue)
	}
	n.value = newVal
	return nil
}

func (n *Number) InsertChild(child ASTNode, location int) error {
	return fmt.Errorf("Number cannot perform insert child operation")
}

func (n *Number) isExpression() {

}

func (n *Number) isASTNode() {

}

func (n *Number) String() string {
	return fmt.Sprintf("%d", n.value)
}

// Interface guard
var _ Expression = (*Number)(nil)
var _ ASTNode = (*Number)(nil)

type Sensor struct {
	sensorType string
}

func (s *Sensor) NodeType() string {
	return "Sensor"
}

func (s *Sensor) Children() []ASTNode {
	return nil
}

func (s *Sensor) Clone() ASTNode {
	return &Sensor{sensorType: s.sensorType}
}

func (s *Sensor) ReplaceChild(oldChild, newChild ASTNode) error {
	return fmt.Errorf("Sensor node does not have children")
}

func (s *Sensor) RemoveChild(child ASTNode) error {
	return fmt.Errorf("Sensor %v cannot remove child %v", s, child)
}

func (s *Sensor) SwapChildren(firstChild, secondChild ASTNode) error {
	return fmt.Errorf("Sensor cannot swap children")
}

func (s *Sensor) Transform(newValue any) error {
	return fmt.Errorf("Sensor cannot be transformed")
}

func (s *Sensor) InsertChild(child ASTNode, location int) error {
	return fmt.Errorf("Sensor cannot perform insert child operation")
}

func (s *Sensor) isExpression() {

}

func (s *Sensor) isSensor() {

}

func (s *Sensor) isASTNode() {

}

func (s *Sensor) String() string {
	return fmt.Sprintf("%s", s.sensorType)
}

// Interface guard
var _ SensorInterface = (*Sensor)(nil)
var _ Expression = (*Sensor)(nil)
var _ ASTNode = (*Sensor)(nil)

type DirectedSensor struct {
	Sensor
	operand Expression
}

func (ds *DirectedSensor) NodeType() string {
	return "DirectedSensor"
}

func (ds *DirectedSensor) Children() []ASTNode {
	return []ASTNode{ds.operand}
}

func (ds *DirectedSensor) Clone() ASTNode {
	return &DirectedSensor{Sensor: ds.Sensor, operand: ds.operand.Clone().(Expression)}
}

func (ds *DirectedSensor) ReplaceChild(oldChild, newChild ASTNode) error {
	newExp, isExp := newChild.(Expression)
	if !isExp {
		return fmt.Errorf("newChild is %T, does not implement Expression", newChild)
	}

	if ds.operand == oldChild {
		ds.operand = newExp
		return nil
	} else {
		return fmt.Errorf("oldChild %v not found in ds %v", oldChild, ds)
	}
}

func (ds *DirectedSensor) Transform(newValue any) error {
	newToken, ok := newValue.(LexedToken)
	if !ok {
		return fmt.Errorf("DirectedSensor node requires newValue of type LexedToken, got %T", newValue)
	}
	if newToken.TokenType == tAhead || newToken.TokenType == tNearby || newToken.TokenType == tRandom {
		ds.sensorType = newToken.Lexeme
		return nil
	}
	return fmt.Errorf("invalid token for DirectedSensor: %v", newToken.TokenType)
}
func (ds *DirectedSensor) isExpression() {

}

func (ds *DirectedSensor) isSensor() {

}

func (ds *DirectedSensor) String() string {
	return fmt.Sprintf("%s[%s]", ds.sensorType, ds.operand)
}

type SensorInterface interface {
	Expression
	isSensor()
}

// Interface guard
var _ SensorInterface = (*DirectedSensor)(nil)
var _ Expression = (*DirectedSensor)(nil)
var _ ASTNode = (*DirectedSensor)(nil)
