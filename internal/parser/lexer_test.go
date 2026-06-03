package parser

import (
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `MEMSIZE >= 50 and DEFENSE = 10 --> eat; // This is a comment`

	expectedTokens := []Token{
		T_MEMSIZE,
		T_GEQU,
		T_NUMBER,
		T_AND,
		T_DEFENSE,
		T_EQU,
		T_NUMBER,
		T_COMM,
		T_EAT,
		T_SEMICOLON,
	}

	l := NewLexer(input)

	for i, expected := range expectedTokens {
		tokenRule := l.NextToken()

		if tokenRule.Type != expected {
			t.Fatalf("Test[%d] failed: expected token %v, got %v", i, expected, tokenRule.Type)
		}
	}
}
