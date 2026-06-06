package parser

import (
	"strings"
	"testing"
)

func TestParser(t *testing.T) {
	tests := []struct {
		name              string
		input             string
		expectedErr       string
		expectedRuleCount int
	}{
		{
			name:              "Basic valid rule",
			input:             "1 = 1 --> eat;",
			expectedErr:       "",
			expectedRuleCount: 1,
		},
		{
			name:              "Multiple rules",
			input:             "1 = 1 --> wait; 2 = 2 --> forward;",
			expectedErr:       "",
			expectedRuleCount: 2,
		},
		{
			name:              "Complex expressions and conditions",
			input:             "(mem[1] + 2) * 3 > 4 and ahead[1] != 0 --> mem[2] := 5 attack;",
			expectedErr:       "",
			expectedRuleCount: 1,
		},
		{
			name:              "Serve action",
			input:             "1 = 1 --> serve[mem[3]];",
			expectedErr:       "",
			expectedRuleCount: 1,
		},
		{
			name:              "Error: Missing arrow",
			input:             "1 = 1 eat;",
			expectedErr:       "expected '-->'",
			expectedRuleCount: 0,
		},
		{
			name:              "Error: Action not at the end of rule",
			input:             "1 = 1 --> eat wait;",
			expectedErr:       "Commands continue after Action",
			expectedRuleCount: 0,
		},
		{
			name:              "Error: Invalid syntax missing expression",
			input:             "--> eat;",
			expectedErr:       "unexpected symbol: -->",
			expectedRuleCount: 0,
		},
		{
			name:              "Error: Missing semicolon at EOF",
			input:             "1 = 1 --> eat",
			expectedErr:       "Commands continue after Action",
			expectedRuleCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			tokens, err := l.Tokenize()
			if err != nil {
				t.Fatalf("unexpected lexer error: %v", err)
			}

			p := NewParser(tokens)
			program, err := p.Parse()

			if tt.expectedErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.expectedErr)
				}
				if !strings.Contains(err.Error(), tt.expectedErr) {
					t.Fatalf("expected error containing %q, got: %v", tt.expectedErr, err)
				}
			} else {
				if err != nil {
					t.Fatalf("unexpected parser error: %v", err)
				}
				if program == nil {
					t.Fatalf("expected program, got nil")
				}
				if len(program.rules) != tt.expectedRuleCount {
					t.Fatalf("expected %d rules, got %d", tt.expectedRuleCount, len(program.rules))
				}
			}
		})
	}
}
