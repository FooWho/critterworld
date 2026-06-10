package parser

import (
	"testing"
)

func TestExpression_Clone(t *testing.T) {
	numOriginal := &Number{value: 5}
	numClone, ok := numOriginal.Clone().(*Number)
	if !ok {
		t.Fatal("Clone of *Number did not return a *Number")
	}
	if numClone.value != numOriginal.value {
		t.Fatalf("numClone.value does not match numOriginal.value: Clone: %d - Original: %d", numClone.value, numOriginal.value)
	}
	numClone.value = numClone.value * 2
	if numClone.value == numOriginal.value {
		t.Fatalf("changing numClone.value also changed numOriginal: Clone: %d - Original: %d", numClone.value, numOriginal.value)
	}
}
