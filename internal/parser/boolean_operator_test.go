package parser

import (
	"testing"
)

func TestLogicalOperator_SwapChildren(t *testing.T) {
	left := &RelationalOperator{}
	right := &RelationalOperator{}
	lo := &LogicalOperator{leftOperand: left, rightOperand: right}

	err := lo.SwapChildren(left, right)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if lo.leftOperand != right || lo.rightOperand != left {
		t.Errorf("Operands were not swapped")
	}
}
