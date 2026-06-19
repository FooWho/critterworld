package parser

import (
	"fmt"
)

func (m *Mutator) ruleMutationRemove(locus FaultLocus) bool {
	if len(m.ast.rootNode.rules) == 1 {
		return false
	} else {
		err := locus.parent.RemoveChild(locus.node)
		if err != nil {
			return false
		}
		return true
	}
}

func (m *Mutator) ruleMutationSwap(locus FaultLocus) bool {
	rule, ok := locus.node.(*Rule)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (m *Mutator).ruleMutationSwap(), got %T", locus.node))
	}

	swappableCount := len(rule.commands)
	if swappableCount > 0 {
		if _, isAction := rule.commands[swappableCount-1].(ActionInterface); isAction {
			swappableCount--
		}
	}

	if swappableCount < 2 {
		return false
	}

	loc1 := m.rng.Intn(swappableCount)
	loc2 := m.rng.Intn(swappableCount)
	for loc2 == loc1 {
		loc2 = m.rng.Intn(swappableCount)
	}

	rule.commands[loc1], rule.commands[loc2] = rule.commands[loc2], rule.commands[loc1]
	return true
}

func (m *Mutator) ruleMutationReplace(locus FaultLocus) bool {
	rule, ok := locus.node.(*Rule)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (m *Mutator).ruleMutationReplace(), got %T", locus.node))
	}
	program, ok := locus.parent.(*Program)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Program in (m *Mutator).ruleMutationReplace(), got %T", locus.parent))
	}
	replaceableCount := len(program.rules)
	if replaceableCount < 2 {
		return false
	}
	loc := m.rng.Intn(replaceableCount)
	for program.rules[loc] == rule {
		loc = m.rng.Intn(replaceableCount)
	}
	clonedRule := program.rules[loc].Clone()
	for i, rle := range program.rules {
		if rle == rule {
			*rule = Rule{}
			program.rules[i] = clonedRule.(*Rule)
			return true
		}
	}
	return false
}

func (m *Mutator) ruleMutationDuplicate(locus FaultLocus) bool {
	return false
}
