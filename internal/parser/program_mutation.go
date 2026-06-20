package parser

import (
	"fmt"
)

func (m *Mutator) programMutationSwap(locus FaultLocus) bool {
	program, ok := locus.node.(*Program)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Program in (m *Mutator).programMutationSwap(), got %T", locus.node))
	}
	candidateCount := len(program.rules)
	if candidateCount < 2 {
		return false
	}
	firstLoc := m.rng.Intn(candidateCount)
	secondLoc := m.rng.Intn(candidateCount)
	for firstLoc == secondLoc {
		firstLoc = m.rng.Intn(candidateCount)
		secondLoc = m.rng.Intn(candidateCount)
	}
	program.rules[firstLoc], program.rules[secondLoc] = program.rules[secondLoc], program.rules[firstLoc]
	return true
}

func (m *Mutator) programMutationDuplicate(locus FaultLocus) bool {
	program, ok := locus.node.(*Program)
	if !ok {
		panic(fmt.Sprintf("critterworld: invariant violation: expected *Program in (m *Mutator).programMutationDuplicate(), got %T", locus.node))
	}
	candidateCount := len(program.rules)
	if candidateCount < 1 {
		return false
	}
	loc := m.rng.Intn(candidateCount)
	clonedRule := program.rules[loc].Clone().(*Rule)
	program.rules = append(program.rules, clonedRule)
	return true
}
