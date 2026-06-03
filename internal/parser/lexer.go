package parser

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type TokenRule struct {
	Type   Token
	Lexeme string
	Regex  *regexp.Regexp
}

type Lexer struct {
	rules  []TokenRule
	input  string
	cursor int
	line   int
}

func ReadCritterSource(fileName string) (string, string) {
	content, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return splitHeaderFromSource(string(content))
}

func splitHeaderFromSource(src string) (string, string) {
	// If the file begins with the species tag, we will assume it is fully formatted source with the header
	// If not, we will assume it source with no header.
	var header string
	var source string
	var splitSource []string
	if strings.Index(src, "species:") == 0 {
		splitSource = strings.Split(src, "\n\n")
		header = splitSource[0]
		header += "\n"
		source = splitSource[1]
		source += "\n"
	} else {
		header = ""
		source = src
	}
	return header, source

}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		rules: lexerRules,
	}
}

func (l *Lexer) NextToken() *LexedToken {
	if l.cursor >= len(l.input) {
		return nil
	}

	subInput := l.input[l.cursor:]

	for _, rule := range l.rules {
		loc := rule.Regex.FindStringIndex(subInput)
		if loc != nil && loc[0] == 0 { // Match starts at the beginning of subInput
			match := subInput[loc[0]:loc[1]]
			l.cursor += len(match)

			if rule.Type == T_WS || rule.Type == T_COMMENT {
				return l.NextToken()
			}
			return &LexedToken{Type: rule.Type, Lexeme: match}
		}
	}

	panic(fmt.Sprintf("This should never happen - Unexpected character at index %d: %q", l.cursor, subInput[0]))
}
