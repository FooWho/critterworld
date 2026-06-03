package parser

import (
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// TokenRule holds a token identifier, the canonical lexeme for the token, and the
// regular expression used by the lexer to match the token.
type TokenRule struct {
	Type   Token
	Lexeme string
	Regex  *regexp.Regexp
}

// LexedToken is the output of the lexer. The lexer performs the lexical analysis
// on a string and emits the tokens that were identified.
type LexedToken struct {
	Type   Token
	Lexeme string
}

// Lexer is the representation of the lexer state. It contains the TokenRules, the input to
// be analyzed, and the current location of the cursor within the input string.
type Lexer struct {
	rules  []TokenRule
	input  string
	cursor int
}

// ReadCritterSource attempts to open and read from the file identified by the fileName
// parameter. If successful, it returns a string containing the header information from
// the critter source file and the critter source code. The header and the source are separated
// by a blank line in a critter source file. For example:
//
//	species: Proto-Critter
//	memsize: 7
//
//	ahead[1] <= -5 --> eat;
//
// Here, the header contains a species identifier and the size of memory available to the
// critter. The genome source is the rule following the blank line.
func ReadCritterSource(fileName string) (string, string, error) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}
	header, source := splitHeaderFromSource(string(content))
	return header, source, nil

}

func splitHeaderFromSource(src string) (string, string) {
	var header string
	var source string
	var splitSource []string
	if strings.Index(src, "species:") == 0 {
		splitSource = strings.SplitN(src, "\n\n", 2)
		header = splitSource[0]
		header += "\n"
		if len(splitSource) > 1 {
			source = splitSource[1]
			source += "\n"
		}
	} else {
		header = ""
		source = src
	}
	return header, source

}

// NewLexer initializes and returns a new Lexer configured for the given input string.
// It will immediately advance to the first character.
func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		rules: lexerRules,
	}
}

// Tokenize runs the lexical analysis over the input and emits the matched
// tokens as a slice of LexedTokens. It will return an error if the Lexer
// was not properly initialized or if encounters a T_MISMATCH token.
func (l *Lexer) Tokenize() ([]*LexedToken, error) {
	if l.input == "" || l.rules == nil {
		return make([]*LexedToken, 0), errors.New("Lexer not initialized.")
	}
	var tokens []*LexedToken
	for {
		token, err := l.nextToken()
		if err != nil {
			if token.Type != T_MISMATCH {
				return tokens, err
			} else {
				tokens = append(tokens, token)
				return tokens, err
			}
		}
		if token == nil {
			break
		}
		tokens = append(tokens, token)
	}
	return tokens, nil
}

func (l *Lexer) nextToken() (*LexedToken, error) {
	if l.cursor == len(l.input) {
		return nil, nil
	}
	if l.cursor > len(l.input) {
		return nil, errors.New("Attempt to read beyond end of stream.")
	}

	subInput := l.input[l.cursor:]

	for _, rule := range l.rules {
		loc := rule.Regex.FindStringIndex(subInput)
		if loc != nil && loc[0] == 0 { // Match starts at the beginning of subInput
			match := subInput[loc[0]:loc[1]]
			l.cursor += len(match)

			if rule.Type == T_WS || rule.Type == T_COMMENT {
				tok, err := l.nextToken()
				return tok, err
			}
			if rule.Type == T_MISMATCH {
				return &LexedToken{Type: rule.Type, Lexeme: match}, fmt.Errorf("Unknown token at index: %d: %q", l.cursor, subInput[0])
			}
			return &LexedToken{Type: rule.Type, Lexeme: match}, nil
		}
	}
	return nil, fmt.Errorf("This should never happen - Unexpected character at index %d: %q", l.cursor, subInput[0])
}
