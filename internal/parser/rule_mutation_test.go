package parser

import (
	"math/rand"
	"testing"
)

type stubNode struct {
	children []ASTNode
}

func (s *stubNode) String() string                                 { return "stub" }
func (s *stubNode) Children() []ASTNode                             { return s.children }
func (s *stubNode) Clone() ASTNode                                  { return &stubNode{children: s.children} }
func (s *stubNode) RemoveChild(child ASTNode) error                 { return nil }
func (s *stubNode) SwapChildren(firstChild, secondChild ASTNode) error { return nil }
func (s *stubNode) ReplaceChild(oldChild, newChild ASTNode) error   { return nil }
func (s *stubNode) Transform(newValue any) error                    { return nil }
func (s *stubNode) InsertChild(child ASTNode, location int) error   { return nil }
func (s *stubNode) isASTNode()                                       {}

func makeRule(commands ...Command) *Rule {
	return &Rule{condition: &RelationalOperator{}, commands: commands}
}

func makeAction() *Action {
	return &Action{actionType: LexedToken{TokenType: tEat, Lexeme: "eat"}}
}

func TestRuleMutationRemove_SingleRuleReturnsFalse(t *testing.T) {
	rule := makeRule(&Update{})
	program := &Program{rules: []*Rule{rule}}
	m := NewMutator(NewAbstractSyntaxTree(program))

	ok := m.ruleMutationRemove(FaultLocus{parent: program, node: rule})
	if ok {
		t.Fatalf("expected remove mutation to fail when there is only one rule")
	}
	if len(program.rules) != 1 {
		t.Fatalf("expected program to still have one rule, got %d", len(program.rules))
	}
}

func TestRuleMutationRemove_MultipleRulesRemovesTarget(t *testing.T) {
	rule1 := makeRule(&Update{})
	rule2 := makeRule(&Update{})
	program := &Program{rules: []*Rule{rule1, rule2}}
	m := NewMutator(NewAbstractSyntaxTree(program))

	ok := m.ruleMutationRemove(FaultLocus{parent: program, node: rule1})
	if !ok {
		t.Fatalf("expected remove mutation to succeed with multiple rules")
	}
	if len(program.rules) != 1 {
		t.Fatalf("expected one rule remaining, got %d", len(program.rules))
	}
	if program.rules[0] != rule2 {
		t.Fatalf("expected remaining rule to be rule2")
	}
}

func TestRuleMutationRemove_ParentRejectsNodeReturnsFalse(t *testing.T) {
	rule1 := makeRule(&Update{})
	rule2 := makeRule(&Update{})
	program := &Program{rules: []*Rule{rule1, rule2}}
	m := NewMutator(NewAbstractSyntaxTree(program))

	notAProgramChild := &Update{}
	ok := m.ruleMutationRemove(FaultLocus{parent: program, node: notAProgramChild})
	if ok {
		t.Fatalf("expected remove mutation to fail when parent cannot remove node")
	}
	if len(program.rules) != 2 {
		t.Fatalf("expected program rules to remain unchanged, got %d", len(program.rules))
	}
}

func TestRuleMutationSwap_NonRuleReturnsFalse(t *testing.T) {
	m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{makeRule(&Update{})}}))

	ok := m.ruleMutationSwap(FaultLocus{parent: nil, node: &Update{}})
	if ok {
		t.Fatalf("expected swap mutation to fail for non-rule locus node")
	}
}

func TestRuleMutationSwap_InsufficientSwappableCommandsReturnsFalse(t *testing.T) {
	tests := []struct {
		name     string
		commands []Command
	}{
		{name: "no commands", commands: []Command{}},
		{name: "single update", commands: []Command{&Update{}}},
		{name: "single action", commands: []Command{makeAction()}},
		{name: "update then action", commands: []Command{&Update{}, makeAction()}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			rule := makeRule(tc.commands...)
			m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{rule, makeRule(&Update{})}}))

			ok := m.ruleMutationSwap(FaultLocus{node: rule})
			if ok {
				t.Fatalf("expected swap mutation to fail for %s", tc.name)
			}
		})
	}
}

func TestRuleMutationSwap_TwoUpdatesAlwaysSwap(t *testing.T) {
	rule := makeRule(&Update{}, &Update{})
	firstBefore := rule.commands[0]
	secondBefore := rule.commands[1]
	m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{rule, makeRule(&Update{})}}))

	rand.Seed(1)
	ok := m.ruleMutationSwap(FaultLocus{node: rule})
	if !ok {
		t.Fatalf("expected swap mutation to succeed")
	}
	if rule.commands[0] != secondBefore || rule.commands[1] != firstBefore {
		t.Fatalf("expected two-command rule to swap positions")
	}
}

func TestRuleMutationSwap_ActionRemainsFinalAndCommandSetPreserved(t *testing.T) {
	u1 := &Update{}
	u2 := &Update{}
	u3 := &Update{}
	action := makeAction()
	rule := makeRule(u1, u2, u3, action)
	m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{rule, makeRule(&Update{})}}))

	seenBefore := map[Command]int{}
	for _, cmd := range rule.commands {
		seenBefore[cmd]++
	}

	rand.Seed(2)
	ok := m.ruleMutationSwap(FaultLocus{node: rule})
	if !ok {
		t.Fatalf("expected swap mutation to succeed")
	}

	if _, isAction := rule.commands[len(rule.commands)-1].(ActionInterface); !isAction {
		t.Fatalf("expected trailing command to remain an action")
	}

	seenAfter := map[Command]int{}
	for _, cmd := range rule.commands {
		seenAfter[cmd]++
	}

	if len(seenBefore) != len(seenAfter) {
		t.Fatalf("expected same number of command instances before and after swap")
	}
	for cmd, count := range seenBefore {
		if seenAfter[cmd] != count {
			t.Fatalf("expected command instance multiplicity to be preserved")
		}
	}
}

func TestRuleMutationReplaceAndDuplicate_CurrentlyReturnFalse(t *testing.T) {
	rule := makeRule(&Update{}, &Update{})
	m := NewMutator(NewAbstractSyntaxTree(&Program{rules: []*Rule{rule, makeRule(&Update{})}}))
	locus := FaultLocus{node: rule}

	if m.ruleMutationReplace(locus) {
		t.Fatalf("expected replace mutation to return false while unimplemented")
	}
	if m.ruleMutationDuplicate(locus) {
		t.Fatalf("expected duplicate mutation to return false while unimplemented")
	}
}

func TestRuleFaultInjector_DoesNotPanicAndMaintainsNonEmptyProgram(t *testing.T) {
	rule1 := makeRule(&Update{}, &Update{}, makeAction())
	rule2 := makeRule(&Update{}, &Update{})
	program := &Program{rules: []*Rule{rule1, rule2}}
	m := NewMutator(NewAbstractSyntaxTree(program))

	rand.Seed(3)
	for i := 0; i < 50; i++ {
		_ = m.ruleFaultInjector(FaultLocus{parent: program, node: program.rules[0]})
		if len(program.rules) == 0 {
			t.Fatalf("expected program to keep at least one rule")
		}
		for _, r := range program.rules {
			if len(r.commands) == 0 {
				t.Fatalf("expected each rule to retain at least one command")
			}
			for idx, cmd := range r.commands {
				if _, isAction := cmd.(ActionInterface); isAction && idx != len(r.commands)-1 {
					t.Fatalf("expected action to remain final command")
				}
			}
		}
	}
}
