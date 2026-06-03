package parser

import (
	"fmt"
)

type AbstractSyntaxTree struct {
	RootNode  *Program
	nodeCount int
}

type ASTNode interface {
	Type() string
	Children() []ASTNode
	Clone() ASTNode
}

type Program struct {
	Rules []*Rule
}

func (p *Program) addRule(rule *Rule) {
	p.Rules = append(p.Rules, rule)
}

func (p *Program) Type() string {
	return "Program"
}

func (p *Program) Children() []ASTNode {
	children := make([]ASTNode, len(p.Rules))
	for i, rule := range p.Rules {
		children[i] = rule
	}
	return children
}

func (p *Program) Clone() ASTNode {
	var cloneP = Program{}

	for _, rule := range p.Rules {
		cloneR, ok := rule.Clone().(*Rule)
		if !ok {
			panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (p *Program).Clone, got %T", cloneR))
		}
		cloneP.addRule(cloneR)
	}
	return &cloneP
}

type Rule struct {
}

func (r *Rule) Type() string {
	return ""
}

func (r *Rule) Children() []ASTNode {
	return nil
}

func (r *Rule) Clone() ASTNode {
	return nil
}

var test Program = Program{}

var _ ASTNode = &test
