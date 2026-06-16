package parser

import (
	"fmt"
	"slices"
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

func (r *Rule) ReplaceChild(oldChild, newChild ASTNode) error {
	if r.condition == oldChild {
		newCond, ok := newChild.(BooleanOperator)
		if !ok {
			return fmt.Errorf("expected BooleanOperator, got %T", newChild)
		}
		r.condition = newCond
		return nil
	}

	for i, command := range r.commands {
		if command == oldChild {
			newCmd, ok := newChild.(Command)
			if !ok {
				return fmt.Errorf("expected Command, got %T", newChild)
			}

			if _, isAction := newCmd.(ActionInterface); isAction && i != len(r.commands)-1 {
				return fmt.Errorf("an Action can only be the final command in a rule")
			}

			r.commands[i] = newCmd
			return nil
		}
	}
	return fmt.Errorf("oldChild not found in Rule")
}

func (r *Rule) RemoveChild(child ASTNode) error {
	if r.condition == child {
		return fmt.Errorf("Rule %v cannot remove condition %v", r, r.condition)
	}
	if len(r.commands) == 1 {
		return fmt.Errorf("rule %v cannot remove only command")
	}
	for i, command := range r.commands {
		if command == child {
			r.commands = slices.Delete(r.commands, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("child %v not found in rule %v", child, r)
}

func (r *Rule) String() string {
	var str string
	str = fmt.Sprintf("%s --> \n", r.condition)
	for _, command := range r.commands {
		str += fmt.Sprintf("      %s\n", command)
	}
	return str
}

func (r *Rule) isASTNode() {
}

// Interface guard
var _ ASTNode = (*Rule)(nil)
