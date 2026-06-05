package parser

import (
	"fmt"
	"strconv"
)

type Parser struct {
	tokens     []*LexedToken
	cursor     int
	tokenCount int
}

func NewParser(input []*LexedToken) *Parser {
	return &Parser{
		tokens:     input,
		cursor:     0,
		tokenCount: len(input),
	}
}

func (p *Parser) peek() *LexedToken {
	if p.cursor < p.tokenCount {
		return p.tokens[p.cursor]
	}

	return &LexedToken{tNone, tNone.String()}

}

func (p *Parser) next() *LexedToken {
	if p.cursor < p.tokenCount {
		token := p.tokens[p.cursor]
		p.cursor++
		return token
	}

	return &LexedToken{tNone, tNone.String()}
}

func (p *Parser) Parse() (*Program, error) {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("parsing failed: %v", r)
		}
	}()
	var program Program
	for p.peek().TokenType != tNone {
		program.AddRule(p.ParseRule())
	}
	return &program, err
}

func (p *Parser) ParseRule() *Rule {
	return &Rule{condition: p.ParseCondition(), commands: p.ParseCommands()}
}

func (p *Parser) ParseCommands() []Command {
	if p.next().TokenType != tComm {
		panic(fmt.Sprintf("critterworld: parse error: expected '-->' in (p *Parser).ParseCommands(), got %s", p.peek().Lexeme))
	}
	var commands []Command
	command := p.ParseCommand()
	commands = append(commands, command)
	for p.peek().TokenType != tSemicolon {
		if command.NodeType() == "Action" || command.NodeType() == "ServeAction" {
			panic("critterworld: parse error: in (p *Parser).ParseCommands(), got Commands continue after Action")
		}
		command = p.ParseCommand()
		commands = append(commands, command)
	}
	_ = p.next()
	return commands
}

func (p *Parser) ParseCommand() Command {
	switch p.peek().TokenType {
	case tMem, tMemSize, tDefense, tOffense, tEnergy, tPass, tPosture:
		return p.ParseUpdate()
	case tForward, tBackward, tLeft, tRight, tEat, tAttack, tGrow, tBud, tServe:
		return p.ParseAction()
	default:
		panic(fmt.Sprintf("critterworld: parse error: in (p *Parser).ParseCommand(), got %s but expected a Command", p.peek().TokenType))
	}

}

func (p *Parser) ParseCondition() BooleanOperator {
	conjunction := p.ParseConjunction()

	for p.peek().TokenType == tOr {
		op := p.next()
		logOp := &LogicalOperator{}
		logOp.leftOperand = conjunction
		logOp.operator = *op
		logOp.rightOperand = p.ParseConjunction()
		conjunction = logOp
	}

	return conjunction
}

func (p *Parser) ParseConjunction() BooleanOperator {
	relation := p.ParseRelation()
	for p.peek().TokenType == tAnd {
		op := p.next()
		logOp := &LogicalOperator{}
		logOp.leftOperand = relation
		logOp.operator = *op
		logOp.rightOperand = p.ParseRelation()
		relation = logOp
	}
	return relation
}

func (p *Parser) ParseRelation() BooleanOperator {
	if p.peek().TokenType == tLBrace {
		_ = p.next()
		condition := p.ParseCondition()
		if p.next().TokenType != tRBrace {
			panic(fmt.Sprintf("critterworld: parse error: expected '}' in (p *Parser).ParseRelation(), got %s", p.peek().Lexeme))
		}
		return condition
	}

	leftOperand := p.ParseExpression()
	operator := p.next()
	if operator.TokenType != tGreat && operator.TokenType != tGequ && operator.TokenType != tEqu &&
		operator.TokenType != tLequ && operator.TokenType != tLess && operator.TokenType != tNequ {
		panic(fmt.Sprintf("critterworld: parse error: expected RelationalOperator in (p *Parser).ParseRelation(), got %s", p.peek().Lexeme))
	}
	rightOperand := p.ParseExpression()
	return &RelationalOperator{operator: *operator, leftOperand: leftOperand, rightOperand: rightOperand}
}

func (p *Parser) ParseExpression() Expression {
	expression := p.ParseTerm()
	for p.peek().TokenType == tPlus || p.peek().TokenType == tMinus {
		operator := p.next()
		binOp := &BinaryOperator{
			operator:     *operator,
			leftOperand:  expression,
			rightOperand: p.ParseTerm(),
		}
		expression = binOp
	}
	return expression
}

func (p *Parser) ParseTerm() Expression {
	term := p.ParseFactor()
	for p.peek().TokenType == tStar || p.peek().TokenType == tDiv || p.peek().TokenType == tMod {
		operator := p.next()
		binOp := &BinaryOperator{
			operator:     *operator,
			leftOperand:  term,
			rightOperand: p.ParseFactor(),
		}
		term = binOp
	}
	return term
}

func (p *Parser) ParseFactor() Expression {
	token := p.peek()
	switch token.TokenType {
	case tMem, tMemSize, tDefense, tOffense, tSize, tEnergy, tPass, tPosture:
		return p.ParseMemNode()
	case tNearby, tAhead, tRandom, tSmell:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented", token.Lexeme))
	case tMinus:
		_ = p.next()
		return &UnaryOperator{operator: LexedToken{tMinus, "-"}, operand: p.ParseFactor()}
	case tLParen:
		_ = p.next()
		expression := p.ParseExpression()
		paren := p.next()
		if paren.TokenType != tRParen {
			panic(fmt.Sprintf("critterworld: parse error: In (p *Parser).ParseFactor(), expected ')' but saw: %s", paren.Lexeme))
		}
		return expression
	case tNumber:
		number := p.ParseNumber()
		return number
	default:
		panic(fmt.Sprintf("critterworld: parse error: In (p *Parser).ParseFactor(), unexpected symbol: %s", token.Lexeme))
	}
}

func (p *Parser) ParseAction() ActionInterface {
	if p.peek().TokenType == tServe {
		_ = p.next()
		return &ServeAction{operand: p.ParseExpression()}
	}
	return &Action{actionType: *p.next()}
}

func (p *Parser) ParseUpdate() *Update {
	return &Update{destination: p.ParseMemNode(), source: p.ParseUpdateSource()}
}

func (p *Parser) ParseUpdateSource() Expression {
	if p.peek().TokenType != tAssign {
		panic(fmt.Sprintf("critterworld: parse error: in (p *Parser).ParseUpdateSource(), got %s but expected ':='", p.peek().TokenType))
	}
	_ = p.next()
	return p.ParseExpression()
}

func (p *Parser) ParseMemNode() *MemNode {
	token := p.next()
	switch token.TokenType {
	case tMemSize:
		return &MemNode{operand: &Number{value: 0}}
	case tDefense:
		return &MemNode{operand: &Number{value: 1}}
	case tOffense:
		return &MemNode{operand: &Number{value: 2}}
	case tSize:
		return &MemNode{operand: &Number{value: 3}}
	case tEnergy:
		return &MemNode{operand: &Number{value: 4}}
	case tPass:
		return &MemNode{operand: &Number{value: 5}}
	case tPosture:
		return &MemNode{operand: &Number{value: 6}}
	case tMem:
		token = p.next()
		if token.TokenType != tLBracket {
			panic(fmt.Sprintf("critterworld: Parse error: In (p *Parser).ParseMemNode(), got %s, but expected '['.", token.Lexeme))
		}
		expression := p.ParseExpression()
		token = p.next()
		if token.TokenType != tRBracket {
			panic(fmt.Sprintf("critterworld: Parse error: In (p *Parser).ParseMemNode(), got %s, but expected ']'.", token.Lexeme))
		}
		return &MemNode{operand: expression}
	default:
		panic(fmt.Sprintf("critterworld: Parse error: In (p *Parser).ParseMemNode(), got %s, but expected 'mem'.", token.Lexeme))
	}
}

func (p *Parser) ParseNumber() *Number {
	token := p.next()
	value, err := strconv.Atoi(token.Lexeme)
	if err != nil {
		panic(fmt.Sprintf("critterworld: parse error in (p *Parser).ParseNumber(), got %s, but expected a number.", token.Lexeme))
	}
	return &Number{value: value}
}
