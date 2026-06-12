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

func (p *Program) ReplaceChild(oldChild *Rule, newChild *Rule) error {
	for i, rule := range p.rules {
		if rule == oldChild {
			p.rules[i] = newChild
			return nil
		}
	}
	return fmt.Errorf("rule was not located in program: %v", *oldChild)
}

func (p *Program) RemoveChild(child *Rule) error {
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

func (p *Program) SwapChildren(first *Rule, second *Rule) error {
	firstLocation := -1
	secondLocation := -1

	for i, rule := range p.rules {
		if rule == first {
			firstLocation = i
			if secondLocation >= 0 {
				break
			}
		}
		if rule == second {
			secondLocation = i
			if firstLocation >= 0 {
				break
			}
		}
	}
	if firstLocation >= 0 && secondLocation >= 0 {
		if firstLocation < secondLocation {
			p.rules = slices.Delete(p.rules, firstLocation, firstLocation+1)
			p.rules = slices.Delete(p.rules, secondLocation-1, secondLocation)
			return nil
		} else {
			p.rules = slices.Delete(p.rules, secondLocation, secondLocation+1)
			p.rules = slices.Delete(p.rules, firstLocation-1, firstLocation)
			return nil
		}
	}
	return fmt.Errorf("unable to swap rule %v and rule %v", first, second)
}

func (p *Program) InsertChild(child *Rule, location int) error {
	if location > len(p.rules) {
		return fmt.Errorf("unable to insert %v at %d", child, location)
	}
	p.rules = slices.Insert(p.rules, location, child)
	return nil
}

// Interface guard
var _ ASTNode = (*Program)(nil)
