package parser

import "fmt"

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
