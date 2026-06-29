package parser

import (
	"testing"
)

func TestUpdateMutationRemove_NonUpdatePanics(t *testing.T) {
	m := NewMutator(NewAbstractSyntaxTree(&Program{}))
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m.updateMutationRemove(FaultLocus{node: &Rule{}})
}

func TestUpdateMutationRemove_NonRuleParentPanics(t *testing.T) {
	m := NewMutator(NewAbstractSyntaxTree(&Program{}))
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	m.updateMutationRemove(FaultLocus{node: &Update{}, parent: &Program{}})
}

func TestUpdateMutationRemove(t *testing.T) {
	t.Run("Fails with only one command", func(t *testing.T) {
		update := &Update{}
		rule := makeRule(update)
		program := &Program{rules: []*Rule{rule}}
		m := NewMutator(NewAbstractSyntaxTree(program))

		ok := m.updateMutationRemove(FaultLocus{parent: rule, node: update})

		if ok {
			t.Error("Expected remove to fail with only one command")
		}
		if len(rule.commands) != 1 {
			t.Errorf("Expected rule to still have 1 command, got %d", len(rule.commands))
		}
	})

	t.Run("Succeeds with multiple commands", func(t *testing.T) {
		update1 := &Update{}
		update2 := &Update{}
		rule := makeRule(update1, update2)
		program := &Program{rules: []*Rule{rule}}
		m := NewMutator(NewAbstractSyntaxTree(program))

		ok := m.updateMutationRemove(FaultLocus{parent: rule, node: update1})

		if !ok {
			t.Error("Expected remove to succeed with multiple commands")
		}
		if len(rule.commands) != 1 {
			t.Errorf("Expected rule to have 1 command, got %d", len(rule.commands))
		}
		if rule.commands[0] != update2 {
			t.Error("Expected the correct command to be removed")
		}
	})
}

func TestUpdateMutationSwap(t *testing.T) {
	t.Run("Fails when source is not MemNode", func(t *testing.T) {
		update := &Update{
			destination: &MemNode{operand: &Number{value: 1}},
			source:      &Number{value: 2},
		}
		program := &Program{rules: []*Rule{makeRule(update)}}
		m := NewMutator(NewAbstractSyntaxTree(program))

		ok := m.updateMutationSwap(FaultLocus{node: update})

		if ok {
			t.Error("Expected swap to fail when source is not a MemNode")
		}
	})

	t.Run("Succeeds when source is MemNode", func(t *testing.T) {
		dest := &MemNode{operand: &Number{value: 1}}
		src := &MemNode{operand: &Number{value: 2}}
		update := &Update{destination: dest, source: src}
		program := &Program{rules: []*Rule{makeRule(update)}}
		m := NewMutator(NewAbstractSyntaxTree(program))

		ok := m.updateMutationSwap(FaultLocus{node: update})

		if !ok {
			t.Error("Expected swap to succeed")
		}
		if update.destination != src || update.source != dest {
			t.Error("Expected destination and source to be swapped")
		}
	})
}

func TestUpdateMutationReplace(t *testing.T) {
	t.Run("Fails with less than 2 updates in AST", func(t *testing.T) {
		update := makeCloneableUpdate()
		rule := makeRule(update)
		program := &Program{rules: []*Rule{rule}}
		m := NewMutator(NewAbstractSyntaxTree(program))

		ok := m.updateMutationReplace(FaultLocus{parent: rule, node: update})

		if ok {
			t.Error("Expected replace to fail with only one update in the AST")
		}
	})

	t.Run("Succeeds and replaces with a different update", func(t *testing.T) {
		update1 := &Update{destination: &MemNode{}, source: &Number{value: 1}}
		update2 := &Update{destination: &MemNode{}, source: &Number{value: 2}}
		rule1 := makeRule(update1)
		rule2 := makeRule(update2)
		program := &Program{rules: []*Rule{rule1, rule2}}
		// This will make GetNodesByType return [update1, update2]
		// The RNG will pick index 1, which is update2.
		m := newMutatorWithRNG(NewAbstractSyntaxTree(program), &sequenceRNG{values: []int{1}})

		ok := m.updateMutationReplace(FaultLocus{parent: rule1, node: update1})

		if !ok {
			t.Error("Expected replace to succeed")
		}
		if len(rule1.commands) != 1 {
			t.Fatalf("Expected rule to still have 1 command, got %d", len(rule1.commands))
		}

		newUpdate := rule1.commands[0]
		if newUpdate == update1 {
			t.Error("Expected update1 to be replaced")
		}

		// The replacement should be update2, as ReplaceChild doesn't clone.
		if newUpdate != update2 {
			t.Error("Expected update1 to be replaced by update2")
		}
	})
}
