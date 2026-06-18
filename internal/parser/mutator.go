package parser

import (
	"math/rand"
)

type Mutator struct {
	ast AbstractSyntaxTree
}

type FaultLocus struct {
	parent ASTNode
	node   ASTNode
}

type mutations int

const (
	mRemove mutations = iota
	mSwap
	mReplace
	mTransform
	mInsert
	mDuplicate
)

func NewMutator(ast AbstractSyntaxTree) *Mutator {
	return &Mutator{ast: ast}
}

func (m *Mutator) GetFaultLocus() FaultLocus {
	nodes := m.ast.GetNodes()
	randomIndex := rand.Intn(len(nodes))
	node := nodes[randomIndex]
	parent := m.ast.GetParentOf(node)

	return FaultLocus{parent: parent, node: node}
}

func (m *Mutator) ruleFaultInjector(locus FaultLocus) bool {
	// Remove, Swap, Replace
	ruleMutations := []mutations{mRemove, mSwap, mReplace}
	mutationType := ruleMutations[rand.Intn(len(ruleMutations))]
	switch mutationType {
	case mRemove:
		return m.ruleMutationRemove(locus)
	case mSwap:
		return m.ruleMutationSwap(locus)
	case mReplace:
		return m.ruleMutationReplace(locus)
	default:
		return false
	}
}

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
	return false
}

func (m *Mutator) ruleMutationReplace(locus FaultLocus) bool {
	return false
}
