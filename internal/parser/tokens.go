package parser

import (
	"regexp"
)

//go:generate stringer -type=Token

// The Type field a LexedToken will be one of the following Token constants.
type Token int

const (
	T_MEMSIZE Token = iota
	T_DEFENSE
	T_OFFENSE
	T_SIZE
	T_ENERGY
	T_PASS
	T_POSTURE
	T_COMMENT
	T_COMM
	T_ASSIGN
	T_LEQU
	T_GEQU
	T_NEQU
	T_MEM
	T_WAIT
	T_FORWARD
	T_BACKWARD
	T_LEFT
	T_RIGHT
	T_EAT
	T_ATTACK
	T_GROW
	T_BUD
	T_SERVE
	T_NEARBY
	T_AHEAD
	T_RANDOM
	T_SMELL
	T_AND
	T_OR
	T_MOD
	T_STAR
	T_DIV
	T_PLUS
	T_MINUS
	T_LESS
	T_GREAT
	T_EQU
	T_L_PAREN
	T_R_PAREN
	T_L_BRACKET
	T_R_BRACKET
	T_L_BRACE
	T_R_BRACE
	T_SEMICOLON
	T_NUMBER
	T_WS
	T_MISMATCH
	T_NONE
)

var lexerRules = []TokenRule{
	{TokenType: T_MEMSIZE, Lexeme: "MEMSIZE", Regex: regexp.MustCompile(`^\bMEMSIZE\b`)},
	{TokenType: T_DEFENSE, Lexeme: "DEFENSE", Regex: regexp.MustCompile(`^\bDEFENSE\b`)},
	{TokenType: T_OFFENSE, Lexeme: "OFFENSE", Regex: regexp.MustCompile(`^\bOFFENSE\b`)},
	{TokenType: T_SIZE, Lexeme: "SIZE", Regex: regexp.MustCompile(`^\bSIZE\b`)},
	{TokenType: T_ENERGY, Lexeme: "ENERGY", Regex: regexp.MustCompile(`^\bENERGY\b`)},
	{TokenType: T_PASS, Lexeme: "PASS", Regex: regexp.MustCompile(`^\bPASS\b`)},
	{TokenType: T_POSTURE, Lexeme: "POSTURE", Regex: regexp.MustCompile(`^\bPOSTURE\b`)},
	{TokenType: T_COMMENT, Lexeme: "//", Regex: regexp.MustCompile(`^//.*`)},
	{TokenType: T_COMM, Lexeme: "-->", Regex: regexp.MustCompile(`^-->`)},
	{TokenType: T_ASSIGN, Lexeme: ":=", Regex: regexp.MustCompile(`^:=`)},
	{TokenType: T_LEQU, Lexeme: "<=", Regex: regexp.MustCompile(`^<=`)},
	{TokenType: T_GEQU, Lexeme: ">=", Regex: regexp.MustCompile(`^>=`)},
	{TokenType: T_NEQU, Lexeme: "!=", Regex: regexp.MustCompile(`^!=`)},
	{TokenType: T_MEM, Lexeme: "mem", Regex: regexp.MustCompile(`^\bmem\b`)},
	{TokenType: T_WAIT, Lexeme: "wait", Regex: regexp.MustCompile(`^\bwait\b`)},
	{TokenType: T_FORWARD, Lexeme: "forward", Regex: regexp.MustCompile(`^\bforward\b`)},
	{TokenType: T_BACKWARD, Lexeme: "backward", Regex: regexp.MustCompile(`^\bbackward\b`)},
	{TokenType: T_LEFT, Lexeme: "left", Regex: regexp.MustCompile(`^\bleft\b`)},
	{TokenType: T_RIGHT, Lexeme: "right", Regex: regexp.MustCompile(`^\bright\b`)},
	{TokenType: T_EAT, Lexeme: "eat", Regex: regexp.MustCompile(`^\beat\b`)},
	{TokenType: T_ATTACK, Lexeme: "attack", Regex: regexp.MustCompile(`^\battack\b`)},
	{TokenType: T_GROW, Lexeme: "grow", Regex: regexp.MustCompile(`^\bgrow\b`)},
	{TokenType: T_BUD, Lexeme: "bud", Regex: regexp.MustCompile(`^\bbud\b`)},
	{TokenType: T_SERVE, Lexeme: "serve", Regex: regexp.MustCompile(`^\bserve\b`)},
	{TokenType: T_NEARBY, Lexeme: "nearby", Regex: regexp.MustCompile(`^\bnearby\b`)},
	{TokenType: T_AHEAD, Lexeme: "ahead", Regex: regexp.MustCompile(`^\bahead\b`)},
	{TokenType: T_RANDOM, Lexeme: "random", Regex: regexp.MustCompile(`^\brandom\b`)},
	{TokenType: T_SMELL, Lexeme: "smell", Regex: regexp.MustCompile(`^\bsmell\b`)},
	{TokenType: T_AND, Lexeme: "and", Regex: regexp.MustCompile(`^\band\b`)},
	{TokenType: T_OR, Lexeme: "or", Regex: regexp.MustCompile(`^\bor\b`)},
	{TokenType: T_MOD, Lexeme: "mod", Regex: regexp.MustCompile(`^\bmod\b`)},
	{TokenType: T_STAR, Lexeme: "*", Regex: regexp.MustCompile(`^\*`)},
	{TokenType: T_DIV, Lexeme: "/", Regex: regexp.MustCompile(`^/`)},
	{TokenType: T_PLUS, Lexeme: "+", Regex: regexp.MustCompile(`^\+`)},
	{TokenType: T_MINUS, Lexeme: "-", Regex: regexp.MustCompile(`^-`)},
	{TokenType: T_LESS, Lexeme: "<", Regex: regexp.MustCompile(`^<`)},
	{TokenType: T_GREAT, Lexeme: ">", Regex: regexp.MustCompile(`^>`)},
	{TokenType: T_EQU, Lexeme: "=", Regex: regexp.MustCompile(`^=`)},
	{TokenType: T_L_PAREN, Lexeme: "(", Regex: regexp.MustCompile(`^\(`)},
	{TokenType: T_R_PAREN, Lexeme: ")", Regex: regexp.MustCompile(`^\)`)},
	{TokenType: T_L_BRACKET, Lexeme: "[", Regex: regexp.MustCompile(`^\[`)},
	{TokenType: T_R_BRACKET, Lexeme: "]", Regex: regexp.MustCompile(`^\]`)},
	{TokenType: T_L_BRACE, Lexeme: "{", Regex: regexp.MustCompile(`^\{`)},
	{TokenType: T_R_BRACE, Lexeme: "}", Regex: regexp.MustCompile(`^\}`)},
	{TokenType: T_SEMICOLON, Lexeme: ";", Regex: regexp.MustCompile(`^;`)},
	{TokenType: T_NUMBER, Lexeme: "NUMBER", Regex: regexp.MustCompile(`^[0-9]+`)},
	{TokenType: T_WS, Lexeme: "WS", Regex: regexp.MustCompile(`^[\s\n\r]+`)},
	{TokenType: T_MISMATCH, Lexeme: "MISMATCH", Regex: regexp.MustCompile(`^.*`)},
	{TokenType: T_NONE, Lexeme: "NONE", Regex: regexp.MustCompile(``)},
}
