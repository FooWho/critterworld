package parser

import (
	"fmt"
)

func (m *Mutator) updateMutationRemove(locus FaultLocus) bool {
	update, ok := locus.node.(*Update)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Update in (m *Mutator).updateMutationRemove(), got %T", locus.node))
	}
	rule, ok := locus.parent.(*Rule)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (m *Mutator).updateMutationRemove(), got %T", locus.parent))
	}
	if len(rule.commands) < 2 {
		return false
	}
	err := rule.RemoveChild(update)
	if err != nil {
		return false
	}
	return true
}

func (m *Mutator) updateMutationSwap(locus FaultLocus) bool {
	// Both sides must be MemNodes or this fails.
	update, ok := locus.node.(*Update)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Update in (m *Mutator).updateMutationSwap(), got %T", locus.node))
	}
	dest, src := update.destination, update.source
	src, ok = src.(*MemNode)
	if !ok {
		return false
	}
	err := update.SwapChildren(dest, src)
	if err != nil {
		return false
	}
	return true
}

func (m *Mutator) updateMutationReplace(locus FaultLocus) bool {
	node, ok := locus.node.(*Update)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Update in (m *Mutator).updateMutationReplace(), got %T", locus.node))
	}
	parent, ok := locus.parent.(*Rule)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Rule in (m *Mutator).updateMutationReplace(), got %T", locus.parent))
	}
	candidates := GetNodesByType[*Update](&m.ast)
	if len(candidates) < 2 {
		return false
	}
	for candidate := candidates[m.rng.Intn(len(candidates))]; candidate == node; {
		candidate = m.
	}
	
	return false
}
