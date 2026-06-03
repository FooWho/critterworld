// Package parser implements the lexical analysis and parsing pipeline
// for the Critterworld simulation.
//
// It reads raw source code, handles desugaring of memory addresses,
// and constructs the final Abstract Syntax Tree used by the
// execution engine.
//
// The parsing process is divided into two main phases:
//  1. Tokenization (handled by the Lexer)
//  2. Recursive descent parsing (handled by the Parser)
//
// The output of the parser is the abstract syntax tree representing
// the critter's genome.
package parser
