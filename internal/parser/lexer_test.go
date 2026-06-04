package parser

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestTokenize(t *testing.T) {
	type expectedToken struct {
		Type   Token
		Lexeme string
	}

	tests := []struct {
		name           string
		input          string
		expectedTokens []expectedToken
		expectedErr    string // If an error is expected, this string should be contained in the error message
	}{
		{
			name:  "Valid statement with comment",
			input: `MEMSIZE >= 50 and DEFENSE = 10 --> eat; // This is a comment`,
			expectedTokens: []expectedToken{
				{T_MEMSIZE, "MEMSIZE"},
				{T_GEQU, ">="},
				{T_NUMBER, "50"},
				{T_AND, "and"},
				{T_DEFENSE, "DEFENSE"},
				{T_EQU, "="},
				{T_NUMBER, "10"},
				{T_COMM, "-->"},
				{T_EAT, "eat"},
				{T_SEMICOLON, ";"},
			},
		},
		{
			name:  "Valid statement without comment",
			input: `ENERGY <= 50 and ahead[1] = -10 --> eat;`,
			expectedTokens: []expectedToken{
				{T_ENERGY, "ENERGY"},
				{T_LEQU, "<="},
				{T_NUMBER, "50"},
				{T_AND, "and"},
				{T_AHEAD, "ahead"},
				{T_L_BRACKET, "["},
				{T_NUMBER, "1"},
				{T_R_BRACKET, "]"},
				{T_EQU, "="},
				{T_MINUS, "-"},
				{T_NUMBER, "10"},
				{T_COMM, "-->"},
				{T_EAT, "eat"},
				{T_SEMICOLON, ";"},
			},
		},
		{
			name:  "Unknown token mismatch",
			input: `MEMSIZE >= 50 and  | DEFENSE = 10 --> eat; // This is a comment`,
			expectedTokens: []expectedToken{
				{T_MEMSIZE, "MEMSIZE"},
				{T_GEQU, ">="},
				{T_NUMBER, "50"},
				{T_AND, "and"},
				{T_MISMATCH, "| DEFENSE = 10 --> eat; // This is a comment"},
			},
			expectedErr: "Unknown token",
		},
		{
			name:           "Whitespace and comments only",
			input:          "   \n\t // just comments \n",
			expectedTokens: []expectedToken{},
		},
		{
			name:           "Empty input error",
			input:          "",
			expectedTokens: []expectedToken{},
			expectedErr:    "Lexer not initialized.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := NewLexer(tt.input)
			tokens, err := l.Tokenize()

			// Verify Error State
			if tt.expectedErr != "" {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.expectedErr)
				}
				if !strings.Contains(err.Error(), tt.expectedErr) {
					t.Fatalf("expected error containing %q, got: %v", tt.expectedErr, err)
				}
			} else if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			// Verify Token Length
			if len(tokens) != len(tt.expectedTokens) {
				t.Fatalf("expected %d tokens, got %d", len(tt.expectedTokens), len(tokens))
			}

			// Verify Individual Tokens (Type and Lexeme)
			for i, expected := range tt.expectedTokens {
				if tokens[i].TokenType != expected.Type {
					t.Errorf("token[%d]: expected type %v, got %v", i, expected.Type, tokens[i].TokenType)
				}
				if tokens[i].Lexeme != expected.Lexeme {
					t.Errorf("token[%d]: expected lexeme %q, got %q", i, expected.Lexeme, tokens[i].Lexeme)
				}
			}
		})
	}
}

func TestSplitHeaderFromSource(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedHeader string
		expectedSource string
	}{
		{
			name:           "With header and source",
			input:          "species: Critter 17\nmemsize: 7\n\nPOSTURE != 17 --> eat;",
			expectedHeader: "species: Critter 17\nmemsize: 7\n",
			expectedSource: "POSTURE != 17 --> eat;\n",
		},
		{
			name:           "Without header",
			input:          "POSTURE != 17 --> eat;",
			expectedHeader: "",
			expectedSource: "POSTURE != 17 --> eat;",
		},
		{
			name:           "Header only (no double newline)",
			input:          "species: Critter 17\nmemsize: 7",
			expectedHeader: "species: Critter 17\nmemsize: 7\n",
			expectedSource: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			header, source := splitHeaderFromSource(tt.input)
			if header != tt.expectedHeader {
				t.Errorf("expected header %q, got %q", tt.expectedHeader, header)
			}
			if source != tt.expectedSource {
				t.Errorf("expected source %q, got %q", tt.expectedSource, source)
			}
		})
	}
}

func TestReadCritterSource(t *testing.T) {
	// Create a temporary directory for safe, isolated file I/O testing
	tempDir := t.TempDir()
	tempFile := filepath.Join(tempDir, "test_critter.crtr")

	content := "species: Test\n\n1 = 1 --> eat;"
	err := os.WriteFile(tempFile, []byte(content), 0644)
	if err != nil {
		t.Fatalf("Failed to write temp file: %v", err)
	}

	t.Run("Valid File", func(t *testing.T) {
		header, source, err := ReadCritterSource(tempFile)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if header != "species: Test\n" {
			t.Errorf("Expected header %q, got %q", "species: Test\n", header)
		}
		if source != "1 = 1 --> eat;\n" {
			t.Errorf("Expected source %q, got %q", "eat;\n", source)
		}
	})

	t.Run("Missing File", func(t *testing.T) {
		_, _, err := ReadCritterSource(filepath.Join(tempDir, "missing.crtr"))
		if err == nil {
			t.Fatal("Expected error for missing file, got nil")
		}
	})
}
