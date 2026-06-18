package parser

import (
	"math/rand"
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
		return false
	}

	// Determine how many commands are actually swappable.
	// Actions can only be at the end, so we exclude it from our pool.
	swappableCount := len(rule.commands)
	if swappableCount > 0 {
		if _, isAction := rule.commands[swappableCount-1].(ActionInterface); isAction {
			swappableCount--
		}
	}

	// We need at least 2 swappable commands to perform a swap
	if swappableCount < 2 {
		return false
	}

	loc1 := rand.Intn(swappableCount)
	loc2 := rand.Intn(swappableCount)
	for loc2 == loc1 {
		loc2 = rand.Intn(swappableCount)
	}

	rule.commands[loc1], rule.commands[loc2] = rule.commands[loc2], rule.commands[loc1]
	return true
}

func (m *Mutator) ruleMutationReplace(locus FaultLocus) bool {
	return false
}

func (m *Mutator) ruleMutationDuplicate(locus FaultLocus) bool {
	return false
}
