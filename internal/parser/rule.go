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

func (r *Rule) SwapChildren(firstChild, secondChild ASTNode) error {
	firstCommand, commandOk := firstChild.(Command)
	if !commandOk {
		return fmt.Errorf("children for swap must be commands, got %T", firstChild)
	}
	secondCommand, commandOk := secondChild.(Command)
	if !commandOk {
		return fmt.Errorf("children for swap must be commands, got %T", secondChild)
	}
	if _, ok := firstChild.(ActionInterface); ok {
		return fmt.Errorf("an Action can only be the final command in a rule")
	}
	if _, ok := secondChild.(ActionInterface); ok {
		return fmt.Errorf("an Action can only be the final command in a rule")
	}
	firstLocation, secondLocation := -1, -1
	for i, command := range r.commands {
		if command == firstCommand {
			firstLocation = i
			if secondLocation >= 0 {
				break
			}
		} else if command == secondCommand {
			secondLocation = i
			if firstLocation >= 0 {
				break
			}
		}
	}
	if firstLocation >= 0 && secondLocation >= 0 {
		r.commands[firstLocation], r.commands[secondLocation] = r.commands[secondLocation], r.commands[firstLocation]
		return nil
	}
	return fmt.Errorf("children not located for swap")
}

func (r *Rule) Transform(newValue any) error {
	return fmt.Errorf("Rule cannot be transformed")
}

func (r *Rule) InsertChild(child ASTNode, location int) error {
	command, ok := child.(Command)
	if !ok {
		return fmt.Errorf("child to insert must implement Command, got type %T", child)
	}

	if location < 0 || location > len(r.commands) {
		return fmt.Errorf("location %d is out of bounds", location)
	}

	hasExistingAction := false
	if len(r.commands) > 0 {
		if _, ok := r.commands[len(r.commands)-1].(ActionInterface); ok {
			hasExistingAction = true
		}
	}

	if _, isNewAction := command.(ActionInterface); isNewAction {
		if location != len(r.commands) {
			return fmt.Errorf("Action must be final command")
		}
		if hasExistingAction {
			return fmt.Errorf("Command can have only one Action")
		}
	} else if hasExistingAction && location == len(r.commands) {
		return fmt.Errorf("cannot insert a new command after final Action")
	}

	r.commands = slices.Insert(r.commands, location, command)
	return nil
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
