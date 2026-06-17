package parser

import (
	"fmt"
	"slices"
)

type Program struct {
	rules []*Rule
}

func (p *Program) AddRule(rule *Rule) {
	p.rules = append(p.rules, rule)
}

func (p *Program) NodeType() string {
	return "Program"
}

func (p *Program) Children() []ASTNode {
	children := make([]ASTNode, len(p.rules))
	for i, rule := range p.rules {
		children[i] = rule
	}
	return children
}

func (p *Program) Clone() ASTNode {
	var cloneP = Program{}

	for _, rule := range p.rules {
		clonedNode := rule.Clone()
		cloneR, ok := clonedNode.(*Rule)
		if !ok {
			panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (p *Program).Clone(), got %T", clonedNode))
		}
		cloneP.AddRule(cloneR)
	}
	return &cloneP
}

func (p *Program) String() string {
	var str string
	for _, rule := range p.rules {
		str += fmt.Sprintf("%s\n", rule)
	}
	return str
}

func (p *Program) isASTNode() {

}

func (p *Program) ReplaceChild(oldChild, newChild ASTNode) error {
	newRule, ok := newChild.(*Rule)
	if !ok {
		return fmt.Errorf("Program can only contain *Rule nodes, got %T", newChild)
	}

	for i, rule := range p.rules {
		if rule == oldChild {
			p.rules[i] = newRule
			return nil
		}
	}
	return fmt.Errorf("oldChild not found in Program")
}

func (p *Program) RemoveChild(child ASTNode) error {
	if len(p.rules) == 1 {
		return fmt.Errorf("Unable to remove %v for %v", child, *p)
	}
	for i, rule := range p.rules {
		if rule == child {
			p.rules = slices.Delete(p.rules, i, i+1)
			return nil
		}
	}
	return fmt.Errorf("rule was not located in program: %v", child)
}

func (p *Program) SwapChildren(firstChild, secondChild ASTNode) error {
	firstRule, firstOk := firstChild.(*Rule)
	if !firstOk {
		return fmt.Errorf("firstChild %v is of type %T", firstChild, firstChild)
	}
	secondRule, secondOk := secondChild.(*Rule)
	if !secondOk {
		return fmt.Errorf("secondChild %v is of type %T", secondChild, secondChild)
	}
	firstLocation := -1
	secondLocation := -1

	for i, rule := range p.rules {
		if rule == firstRule {
			firstLocation = i
			if secondLocation >= 0 {
				break
			}
		}
		if rule == secondRule {
			secondLocation = i
			if firstLocation >= 0 {
				break
			}
		}
	}
	if firstLocation >= 0 && secondLocation >= 0 {
		p.rules[firstLocation], p.rules[secondLocation] = p.rules[secondLocation], p.rules[firstLocation]
		return nil
	}
	return fmt.Errorf("unable to swap rule %v and rule %v", firstChild, secondChild)
}

func (p *Program) InsertChild(child ASTNode, location int) error {
	childRule, ok := child.(*Rule)
	if !ok {
		return fmt.Errorf("child needs to be type *Rule, got type %T", child)
	}
	if location > len(p.rules) {
		return fmt.Errorf("unable to insert %v at %d", childRule, location)
	}
	p.rules = slices.Insert(p.rules, location, childRule)
	return nil
}

func (p *Program) Transform(newValue any) error {
	return fmt.Errorf("Program cannot be transformed")
}

// Interface guard
var _ ASTNode = (*Program)(nil)
