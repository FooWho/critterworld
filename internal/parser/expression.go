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

func (bo *BinaryOperator) IsExpression() bool {
	return true
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
	if operand.NodeType() == "BinaryOperator" &&
		(bo.operator.TokenType == tStar ||
			bo.operator.TokenType == tDiv ||
			bo.operator.TokenType == tMod) &&
		(operand.(*BinaryOperator).operator.TokenType == tPlus ||
			operand.(*BinaryOperator).operator.TokenType == tMinus) {
		return true
	}
	return false
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

func (uo *UnaryOperator) IsExpression() bool {
	return true
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

func (mn *MemNode) IsExpression() bool {
	return true
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

func (n *Number) IsExpression() bool {
	return true
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

func (s *Sensor) IsExpression() bool {
	return true
}

func (s *Sensor) IsSensor() bool {
	return true
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

func (ds *DirectedSensor) IsExpression() bool {
	return true
}

func (ds *DirectedSensor) IsSensor() bool {
	return true
}

func (ds *DirectedSensor) String() string {
	return fmt.Sprintf("%s[%s]", ds.sensorType, ds.operand)
}

type SensorInterface interface {
	Expression
	IsSensor() bool
}

// Interface guard
var _ SensorInterface = (*DirectedSensor)(nil)
var _ Expression = (*DirectedSensor)(nil)
var _ ASTNode = (*DirectedSensor)(nil)
