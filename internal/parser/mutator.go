package parser

import (
	"math/rand"
	"time"
)

type randomSource interface {
	Intn(int) int
}

type Mutator struct {
	ast AbstractSyntaxTree
	rng randomSource
}

type FaultLocus struct {
	parent ASTNode
	node   ASTNode
}

type mutations int

const (
	mRemove mutations = iota // The node, along with all its descendants, is removed. If the parent of the node being removed needs
	// a replacement child, one of the node’s direct children of the correct kind is randomly selected.

	mSwap // The order of two children of the node is switched.

	mReplace // The node and its descendants are replaced with a randomly selected subtree of the right kind.

	mTransform // The node is replaced with a random, newly created node of the same kind (for example,
	// replacing attack with eat, or + with *), but its children remain the same.

	mInsert // A newly created node is inserted as the parent of the mutated node. The old parent of the mutated
	// node becomes the parent of the inserted node, and the mutated node becomes a child of the
	// inserted node. If the inserted node requires more than one child, the children that are not
	// the original node are copies of randomly chosen nodes of the right kind from the entire rule set.

	mDuplicate // For nodes with a variable number of children, a randomly selected subtree of the right type
	// (as in Replace mutations) is appended to the end of the list of children. This applies to the
	// root node, where a new rule can be added, and also to command nodes, where the sequence of updates
	// can be extended with another update.
)

func NewMutator(ast AbstractSyntaxTree) *Mutator {
	return &Mutator{ast: ast, rng: rand.New(rand.NewSource(time.Now().UnixNano()))}
}

func newMutatorWithRNG(ast AbstractSyntaxTree, rng randomSource) *Mutator {
	if rng == nil {
		panic("critterworld: invariant violation: rng cannot be nil in newMutatorWithRNG")
	}
	return &Mutator{ast: ast, rng: rng}
}

func (m *Mutator) GetFaultLocus() FaultLocus {
	nodes := m.ast.GetNodes()
	randomIndex := m.rng.Intn(len(nodes))
	node := nodes[randomIndex]
	parent := m.ast.GetParentOf(node)

	return FaultLocus{parent: parent, node: node}
}

func (m *Mutator) programFaultInjector(locus FaultLocus) bool {
	// Swap, Duplicate
	programMutations := []mutations{mSwap, mDuplicate}
	mutatationType := programMutations[m.rng.Intn(len(programMutations))]
	switch mutatationType {
	case mSwap:
		return m.programMutationSwap(locus)
	case mDuplicate:
		return m.programMutationDuplicate(locus)
	default:
		return false
	}
}

func (m *Mutator) ruleFaultInjector(locus FaultLocus) bool {
	// Remove, Swap, Replace, Duplicate
	ruleMutations := []mutations{mRemove, mSwap, mReplace, mDuplicate}
	mutationType := ruleMutations[m.rng.Intn(len(ruleMutations))]
	switch mutationType {
	case mRemove:
		return m.ruleMutationRemove(locus)
	case mSwap:
		return m.ruleMutationSwap(locus)
	case mReplace:
		return m.ruleMutationReplace(locus)
	case mDuplicate:
		return m.ruleMutationDuplicate(locus)
	default:
		return false
	}
}
