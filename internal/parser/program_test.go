package parser

import (
	"testing"
)

func TestProgram_SwapChildren(t *testing.T) {
	rule1 := &Rule{}
	rule2 := &Rule{}
	p := &Program{rules: []*Rule{rule1, rule2}}

	err := p.SwapChildren(rule1, rule2)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if p.rules[0] != rule2 || p.rules[1] != rule1 {
		t.Errorf("Rules were not swapped correctly")
	}
}

func TestProgram_RemoveChild(t *testing.T) {
	rule1 := &Rule{}
	rule2 := &Rule{}
	p := &Program{rules: []*Rule{rule1, rule2}}

	err := p.RemoveChild(rule1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if len(p.rules) != 1 || p.rules[0] != rule2 {
		t.Errorf("Rule was not removed correctly")
	}
}
