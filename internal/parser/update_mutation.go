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
	return false
}

func (m *Mutator) updateMutationReplace(locus FaultLocus) bool {
	return false
}
