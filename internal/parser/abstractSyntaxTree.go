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
		clonedNode := rule.Clone()
		cloneR, ok := clonedNode.(*Rule)
		if !ok {
			panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (p *Program).Clone(), got %T", clonedNode))
		}
		cloneP.addRule(cloneR)
	}
	return &cloneP
}

type Rule struct {
	Condition BooleanOperator
	Commands  []Command
}

func (r *Rule) Type() string {
	return "Rule"
}

func (r *Rule) Children() []ASTNode {
	children := make([]ASTNode, len(r.Commands)+1)
	children[0] = r.Condition
	for i, command := range r.Commands {
		children[i+1] = command
	}
	return children
}

func (r *Rule) Clone() ASTNode {
	return nil
}

type BooleanOperator interface {
	ASTNode
	Evaluate() bool
}

type LogicalOperator struct {
}

func (lo *LogicalOperator) Type() string {
	return ""
}

func (lo *LogicalOperator) Children() []ASTNode {
	return nil
}

func (lo *LogicalOperator) Clone() ASTNode {
	return nil
}

func (lo *LogicalOperator) Evaluate() bool {
	return true
}

type Command interface {
	ASTNode
	IsCommand() bool
}

var test Program = Program{}

var _ ASTNode = &test

var test2 = LogicalOperator{}
var _ BooleanOperator = &test2
var _ ASTNode = &test2
