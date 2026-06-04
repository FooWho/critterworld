package parser

import (
	"fmt"
)

type Command interface {
	ASTNode
	IsCommand() bool
}

type Update struct {
	destination ASTNode
	source      ASTNode
}

func (u *Update) NodeType() string {
	return "Update"
}

func (u *Update) Children() []ASTNode {
	children := make([]ASTNode, 2)
	children[0] = u.destination
	children[1] = u.source
	return children
}

func (u *Update) Clone() ASTNode {
	var ok bool
	uClone := Update{}

	clonedDestination := u.destination.Clone()
	uClone.destination, ok = clonedDestination.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (ro *RelationalOperator).Clone(), got %T", clonedDestination))
	}

	clonedSource := u.source.Clone()
	uClone.source, ok = clonedSource.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (ro *RelationalOperator).Clone(), got %T", clonedSource))
	}

	return &uClone
}

func (u *Update) IsCommand() bool {
	return true
}

type ActionInterface interface {
	Command
	IsAction() bool
}

type Action struct {
	actionType LexedToken
}

func (act *Action) NodeType() string {
	return "Action"
}

func (act *Action) Children() []ASTNode {
	return nil
}

func (act *Action) Clone() ASTNode {
	return nil
}

func (act *Action) IsCommand() bool {
	return true
}

func (act *Action) IsAction() bool {
	return true
}

type ServeAction struct {
	Action
	operand Expression
}

func (act *ServeAction) NodeType() string {
	return "ServeAction"
}

func (act *ServeAction) Children() []ASTNode {
	return nil
}

func (act *ServeAction) Clone() ASTNode {
	return nil
}

func (act *ServeAction) IsCommand() bool {
	return true
}

func (act *ServeAction) IsAction() bool {
	return true
}

func (act *ServeAction) isExpression() bool {
	return true
}
