package parser

import (
	"testing"
)

func TestRule_ReplaceChild(t *testing.T) {
	oldCond := &RelationalOperator{}
	cmd := &Update{}
	r := &Rule{condition: oldCond, commands: []Command{cmd}}

	newCond := &LogicalOperator{}
	err := r.ReplaceChild(oldCond, newCond)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if r.condition != newCond {
		t.Errorf("Condition was not replaced")
	}

	newCmd := &Update{}
	err = r.ReplaceChild(cmd, newCmd)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if r.commands[0] != newCmd {
		t.Errorf("Command was not replaced")
	}
}

func TestRule_InsertChild(t *testing.T) {
	r := &Rule{condition: &RelationalOperator{}, commands: []Command{}}

	cmd1 := &Update{}
	err := r.InsertChild(cmd1, 0)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	action := &Action{actionType: LexedToken{TokenType: tEat, Lexeme: "eat"}}
	err = r.InsertChild(action, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	err = r.InsertChild(&Update{}, 2)
	if err == nil {
		t.Errorf("Expected error inserting command after Action")
	}
}
