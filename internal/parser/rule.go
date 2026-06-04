package parser

import (
	"fmt"
)

type Rule struct {
	condition BooleanOperator
	commands  []Command
}

func (r *Rule) NodeType() string {
	return "Rule"
}

func (r *Rule) Children() []ASTNode {
	children := make([]ASTNode, len(r.commands)+1)
	children[0] = r.condition
	for i, command := range r.commands {
		children[i+1] = command
	}
	return children
}

func (r *Rule) Clone() ASTNode {
	var rClone = &Rule{}
	var ok bool
	clonedNode := r.condition.Clone()
	rClone.condition, ok = clonedNode.(BooleanOperator)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected Boolean in (r *Rule).Clone(), got %T", clonedNode))
	}
	rClone.commands = make([]Command, len(r.commands))
	for i, command := range r.commands {
		clonedCmd := command.Clone()
		rClone.commands[i], ok = clonedCmd.(Command)
		if !ok {
			panic(fmt.Sprintf("critterworld: invariant violation: expected Command in (r *Rule).Clone(), got %T", clonedCmd))
		}
	}
	return rClone
}

// Interface guard
var _ ASTNode = (*Rule)(nil)
