package parser

import (
	"fmt"
	"testing"
)

type sequenceRNG struct {
	values []int
	index  int
}

func (s *sequenceRNG) Intn(n int) int {
	if n <= 0 {
		panic("test invariant violated: Intn called with n <= 0")
	}
	if len(s.values) == 0 {
		panic("test invariant violated: sequenceRNG has no values")
	}
	v := s.values[s.index%len(s.values)]
	s.index++
	return v % n
}

func TestProgramMutationSwap_NonProgramPanics(t *testing.T) {
	m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{makeCloneableRule(tEat, "eat")}}))

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected swap mutation to panic for non-program locus node")
		}
	}()

	_ = m.programMutationSwap(FaultLocus{parent: nil, node: &Rule{}})
}

func TestProgramMutationSwap_InsufficientRulesReturnsFalse(t *testing.T) {
	tests := []struct {
		name  string
		rules []*Rule
	}{
		{name: "zero rules", rules: []*Rule{}},
		{name: "one rule", rules: []*Rule{makeCloneableRule(tEat, "eat")}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			program := &Program{rules: tc.rules}
			m := NewMutator(NewAbstractSyntaxTree(program))

			ok := m.programMutationSwap(FaultLocus{node: program})
			if ok {
				t.Fatalf("expected swap mutation to fail with %d rules", len(tc.rules))
			}
			if len(program.rules) != len(tc.rules) {
				t.Fatalf("expected rules length to remain unchanged")
			}
		})
	}
}

func TestProgramMutationSwap_ReRollsSameIndexAndSwaps(t *testing.T) {
	rule1 := makeCloneableRule(tEat, "eat")
	rule2 := makeCloneableRule(tWait, "wait")
	rule3 := makeCloneableRule(tGrow, "grow")
	program := &Program{rules: []*Rule{rule1, rule2, rule3}}
	var rule1str = fmt.Sprint(rule1)
	fmt.Println(rule1str)

	// Force initial equal picks (0,0), then distinct picks (2,1).
	m := newMutatorWithRNG(NewAbstractSyntaxTree(program), &sequenceRNG{values: []int{0, 0, 2, 1}})

	ok := m.programMutationSwap(FaultLocus{node: program})
	if !ok {
		t.Fatalf("expected swap mutation to succeed with 3 rules")
	}

	if program.rules[0] != rule1 {
		t.Fatalf("expected first rule to remain unchanged")
	}
	if program.rules[1] != rule3 || program.rules[2] != rule2 {
		t.Fatalf("expected second and third rules to be swapped")
	}
}

func TestProgramMutationDuplicate_NonProgramPanics(t *testing.T) {
	m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{makeCloneableRule(tEat, "eat")}}))

	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expected duplicate mutation to panic for non-program locus node")
		}
	}()

	_ = m.programMutationDuplicate(FaultLocus{parent: nil, node: &Rule{}})
}

func TestProgramMutationDuplicate_NoRulesReturnsFalse(t *testing.T) {
	program := &Program{rules: []*Rule{}}
	m := NewMutator(NewAbstractSyntaxTree(program))

	ok := m.programMutationDuplicate(FaultLocus{node: program})
	if ok {
		t.Fatalf("expected duplicate mutation to fail with no rules")
	}
	if len(program.rules) != 0 {
		t.Fatalf("expected no rules to be added")
	}
}

func TestProgramMutationDuplicate_AppendsIndependentClone(t *testing.T) {
	rule1 := makeCloneableRule(tEat, "eat")
	rule2 := makeCloneableRule(tWait, "wait")
	program := &Program{rules: []*Rule{rule1, rule2}}

	// Pick index 1 so duplication source is rule2.
	m := newMutatorWithRNG(NewAbstractSyntaxTree(program), &sequenceRNG{values: []int{1}})

	ok := m.programMutationDuplicate(FaultLocus{node: program})
	if !ok {
		t.Fatalf("expected duplicate mutation to succeed")
	}

	if len(program.rules) != 3 {
		t.Fatalf("expected one cloned rule to be appended, got %d rules", len(program.rules))
	}

	cloned := program.rules[2]
	if cloned == rule2 {
		t.Fatalf("expected appended rule to be a clone, not the original")
	}

	clonedAction, ok := cloned.commands[0].(*Action)
	if !ok {
		t.Fatalf("expected cloned rule to contain an Action command")
	}
	if clonedAction.actionType.Lexeme != "wait" {
		t.Fatalf("expected cloned rule action lexeme to match source rule")
	}

	sourceAction := rule2.commands[0].(*Action)
	sourceAction.actionType = LexedToken{TokenType: tGrow, Lexeme: "grow"}
	if clonedAction.actionType.Lexeme != "wait" {
		t.Fatalf("expected clone to remain independent after source mutation")
	}
}
