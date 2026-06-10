package parser

import (
	"testing"
)

func TestUpdate_Clone(t *testing.T) {
	original := &Update{
		destination: &MemNode{operand: &Number{value: 0}},
		source:      &Number{value: 42},
	}

	clonedNode := original.Clone()
	clone, ok := clonedNode.(*Update)
	if !ok {
		t.Fatalf("Clone() did not return an *Update")
	}

	// Modify the clone's source and destination to verify deep copy
	clone.source.(*Number).value = 99
	clone.destination.operand.(*Number).value = 1

	// The original should remain unchanged
	if original.source.(*Number).value != 42 {
		t.Errorf("Modifying clone's source affected the original node!")
	}
	if original.destination.operand.(*Number).value != 0 {
		t.Errorf("Modifying clone's destination affected the original node!")
	}
}

func TestServeAction_Clone(t *testing.T) {
	original := &ServeAction{
		Action:  Action{actionType: LexedToken{TokenType: tServe, Lexeme: "serve"}},
		operand: &Number{value: 15},
	}

	clonedNode := original.Clone()
	clone, ok := clonedNode.(*ServeAction)
	if !ok {
		t.Fatalf("Clone() did not return a *ServeAction")
	}

	// Modify the clone
	clone.operand.(*Number).value = 100

	// Verify original is untouched
	if original.operand.(*Number).value != 15 {
		t.Errorf("Modifying clone affected the original node!")
	}
}

func TestCommand_Children(t *testing.T) {
	tests := []struct {
		name          string
		command       Command
		expectedCount int
	}{
		{
			name: "Update children",
			command: &Update{
				destination: &MemNode{operand: &Number{value: 1}},
				source:      &Number{value: 5},
			},
			expectedCount: 2,
		},
		{
			name: "ServeAction children",
			command: &ServeAction{
				Action:  Action{actionType: LexedToken{TokenType: tServe, Lexeme: "serve"}},
				operand: &Number{value: 5},
			},
			expectedCount: 1,
		},
		{
			name:          "Action children",
			command:       &Action{actionType: LexedToken{TokenType: tEat, Lexeme: "eat"}},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			children := tt.command.Children()
			if len(children) != tt.expectedCount {
				t.Errorf("Expected %d children, got %d", tt.expectedCount, len(children))
			}
		})
	}
}
