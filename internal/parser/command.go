package parser

import (
	"fmt"
)

type Command interface {
	ASTNode
	IsCommand() bool
}

type Update struct {
	destination *MemNode
	source      Expression
}

func (u *Update) NodeType() string {
	return "Update"
}

func (u *Update) Children() []ASTNode {
	return []ASTNode{u.destination, u.source}
}

func (u *Update) Clone() ASTNode {
	var ok bool
	uClone := Update{}

	clonedDestination := u.destination.Clone()
	uClone.destination, ok = clonedDestination.(*MemNode)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *MemNode in (u *Update).Clone(), got %T", clonedDestination))
	}

	clonedSource := u.source.Clone()
	uClone.source, ok = clonedSource.(Expression)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Expression in (u *Update).Clone(), got %T", clonedSource))
	}

	return &uClone
}

func (u *Update) IsCommand() bool {
	return true
}

func (u *Update) String() string {
	return fmt.Sprintf("%s := %s", u.destination, u.source)
}

// Interface guard
var _ Command = (*Update)(nil)
var _ ASTNode = (*Update)(nil)

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
	return &Action{actionType: act.actionType}
}

func (act *Action) IsCommand() bool {
	return true
}

func (act *Action) IsAction() bool {
	return true
}

func (act *Action) String() string {
	return fmt.Sprintf("%s", act.actionType.Lexeme)
}

// Interface guard
var _ Command = (*Action)(nil)
var _ ActionInterface = (*Action)(nil)
var _ ASTNode = (*Action)(nil)

type ServeAction struct {
	Action
	operand Expression
}

func (act *ServeAction) NodeType() string {
	return "ServeAction"
}

func (act *ServeAction) Children() []ASTNode {
	return []ASTNode{act.operand}
}

func (act *ServeAction) Clone() ASTNode {
	return &ServeAction{Action: Action{actionType: act.actionType}, operand: act.operand.Clone().(Expression)}
}

func (act *ServeAction) IsCommand() bool {
	return true
}

func (act *ServeAction) IsAction() bool {
	return true
}

func (act *ServeAction) IsExpression() bool {
	return true
}

func (act *ServeAction) String() string {
	return fmt.Sprintf("serve[%s]", act.operand)
}

// Interface guard
var _ Command = (*ServeAction)(nil)
var _ ActionInterface = (*ServeAction)(nil)
var _ ASTNode = (*ServeAction)(nil)
