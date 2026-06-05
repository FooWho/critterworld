package parser

import (
	"regexp"
)

//go:generate stringer -type=Token

// The Type field a LexedToken will be one of the following Token constants.
type Token int

const (
	tMemSize Token = iota
	tDefense
	TOffense
	tSize
	tEnergy
	tPass
	tPosture
	tComment
	tComm
	tAssign
	tLequ
	tGequ
	tNequ
	tMem
	tWait
	tForward
	tBackward
	tLeft
	tRight
	tEat
	tAttack
	tGrow
	tBud
	tServe
	tNearby
	tAhead
	tRandom
	tSmell
	tAnd
	tOr
	tMod
	tStar
	tDiv
	tPlus
	tMinus
	tLess
	tGreat
	tEqu
	tLParen
	tRParen
	tLBracket
	tRBracket
	tLBrace
	tRBrace
	tSemicolon
	tNumber
	tWS
	tMismatch
	tNone
)

var lexerRules = []TokenRule{
	{TokenType: tMemSize, Lexeme: "MEMSIZE", Regex: regexp.MustCompile(`^\bMEMSIZE\b`)},
	{TokenType: tDefense, Lexeme: "DEFENSE", Regex: regexp.MustCompile(`^\bDEFENSE\b`)},
	{TokenType: TOffense, Lexeme: "OFFENSE", Regex: regexp.MustCompile(`^\bOFFENSE\b`)},
	{TokenType: tSize, Lexeme: "SIZE", Regex: regexp.MustCompile(`^\bSIZE\b`)},
	{TokenType: tEnergy, Lexeme: "ENERGY", Regex: regexp.MustCompile(`^\bENERGY\b`)},
	{TokenType: tPass, Lexeme: "PASS", Regex: regexp.MustCompile(`^\bPASS\b`)},
	{TokenType: tPosture, Lexeme: "POSTURE", Regex: regexp.MustCompile(`^\bPOSTURE\b`)},
	{TokenType: tComment, Lexeme: "//", Regex: regexp.MustCompile(`^//.*`)},
	{TokenType: tComm, Lexeme: "-->", Regex: regexp.MustCompile(`^-->`)},
	{TokenType: tAssign, Lexeme: ":=", Regex: regexp.MustCompile(`^:=`)},
	{TokenType: tLequ, Lexeme: "<=", Regex: regexp.MustCompile(`^<=`)},
	{TokenType: tGequ, Lexeme: ">=", Regex: regexp.MustCompile(`^>=`)},
	{TokenType: tNequ, Lexeme: "!=", Regex: regexp.MustCompile(`^!=`)},
	{TokenType: tMem, Lexeme: "mem", Regex: regexp.MustCompile(`^\bmem\b`)},
	{TokenType: tWait, Lexeme: "wait", Regex: regexp.MustCompile(`^\bwait\b`)},
	{TokenType: tForward, Lexeme: "forward", Regex: regexp.MustCompile(`^\bforward\b`)},
	{TokenType: tBackward, Lexeme: "backward", Regex: regexp.MustCompile(`^\bbackward\b`)},
	{TokenType: tLeft, Lexeme: "left", Regex: regexp.MustCompile(`^\bleft\b`)},
	{TokenType: tRight, Lexeme: "right", Regex: regexp.MustCompile(`^\bright\b`)},
	{TokenType: tEat, Lexeme: "eat", Regex: regexp.MustCompile(`^\beat\b`)},
	{TokenType: tAttack, Lexeme: "attack", Regex: regexp.MustCompile(`^\battack\b`)},
	{TokenType: tGrow, Lexeme: "grow", Regex: regexp.MustCompile(`^\bgrow\b`)},
	{TokenType: tBud, Lexeme: "bud", Regex: regexp.MustCompile(`^\bbud\b`)},
	{TokenType: tServe, Lexeme: "serve", Regex: regexp.MustCompile(`^\bserve\b`)},
	{TokenType: tNearby, Lexeme: "nearby", Regex: regexp.MustCompile(`^\bnearby\b`)},
	{TokenType: tAhead, Lexeme: "ahead", Regex: regexp.MustCompile(`^\bahead\b`)},
	{TokenType: tRandom, Lexeme: "random", Regex: regexp.MustCompile(`^\brandom\b`)},
	{TokenType: tSmell, Lexeme: "smell", Regex: regexp.MustCompile(`^\bsmell\b`)},
	{TokenType: tAnd, Lexeme: "and", Regex: regexp.MustCompile(`^\band\b`)},
	{TokenType: tOr, Lexeme: "or", Regex: regexp.MustCompile(`^\bor\b`)},
	{TokenType: tMod, Lexeme: "mod", Regex: regexp.MustCompile(`^\bmod\b`)},
	{TokenType: tStar, Lexeme: "*", Regex: regexp.MustCompile(`^\*`)},
	{TokenType: tDiv, Lexeme: "/", Regex: regexp.MustCompile(`^/`)},
	{TokenType: tPlus, Lexeme: "+", Regex: regexp.MustCompile(`^\+`)},
	{TokenType: tMinus, Lexeme: "-", Regex: regexp.MustCompile(`^-`)},
	{TokenType: tLess, Lexeme: "<", Regex: regexp.MustCompile(`^<`)},
	{TokenType: tGreat, Lexeme: ">", Regex: regexp.MustCompile(`^>`)},
	{TokenType: tEqu, Lexeme: "=", Regex: regexp.MustCompile(`^=`)},
	{TokenType: tLParen, Lexeme: "(", Regex: regexp.MustCompile(`^\(`)},
	{TokenType: tRParen, Lexeme: ")", Regex: regexp.MustCompile(`^\)`)},
	{TokenType: tLBracket, Lexeme: "[", Regex: regexp.MustCompile(`^\[`)},
	{TokenType: tRBracket, Lexeme: "]", Regex: regexp.MustCompile(`^\]`)},
	{TokenType: tLBrace, Lexeme: "{", Regex: regexp.MustCompile(`^\{`)},
	{TokenType: tRBrace, Lexeme: "}", Regex: regexp.MustCompile(`^\}`)},
	{TokenType: tSemicolon, Lexeme: ";", Regex: regexp.MustCompile(`^;`)},
	{TokenType: tNumber, Lexeme: "NUMBER", Regex: regexp.MustCompile(`^[0-9]+`)},
	{TokenType: tWS, Lexeme: "WS", Regex: regexp.MustCompile(`^[\s\n\r]+`)},
	{TokenType: tMismatch, Lexeme: "MISMATCH", Regex: regexp.MustCompile(`^.*`)},
	{TokenType: tNone, Lexeme: "NONE", Regex: regexp.MustCompile(``)},
}
