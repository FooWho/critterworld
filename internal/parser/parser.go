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

func (p *Parser) Parse() *Program {
	program := &Program{}
	program.rules = make([]*Rule, 0)
	for p.cursor < p.tokenCount {
		rule := p.ParseRule()
		program.rules = append(program.rules, rule)
	}
	return program
}

func (p *Parser) ParseRule() *Rule {
	return &Rule{condition: p.ParseCondition()}
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
		if p.peek().TokenType != tRBrace {

		}
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
	switch operator.TokenType {
	case tGreat, tGequ, tEqu, tLequ, tLess, tNequ:
		rightOperand := p.ParseExpression()
		return &RelationalOperator{operator: *operator, leftOperand: leftOperand, rightOperand: rightOperand}
	default:
		panic(fmt.Sprintf("critterworld: parse error: expected RelationalOperator in (p *Parser).ParseRelation(), got %s", p.peek().Lexeme))
	}
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
	case tMem, tMemSize, tDefense, TOffense, tSize, tEnergy, tPass, tPosture:
		memnode := p.ParseMemNode()
		return memnode
	case tNearby, tAhead, tRandom, tSmell:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case tMinus:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case tLParen:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case tNumber:
		number := p.ParseNumber()
		return number
	default:
		panic(fmt.Sprintf("critterworld: parse error: In (p *Parser).ParseFactor(), unexpected symbol: %s.", token.Lexeme))
	}
	// token = self.peek()
	// match token.tokenType:
	//
	//	case toke if toke in SET_SUGAR | {TOKENS.T_MEM}:
	//	    memNode = self.parseMemNode()
	//	    return memNode
	//	case sensor if sensor in SET_SENSORS:
	//	    sensorNode = self.parseSensor()
	//	    return sensorNode
	//	case TOKENS.T_MINUS:
	//	    op = self.getToken()
	//	    unOp = UnaryOperator()
	//	    unOp.operator = TokenLexeme(op.tokenType, op.lexeme)
	//	    unOp.operand = self.parseFactor()
	//	    return unOp
	//	case TOKENS.T_L_PAREN:
	//	    token = self.getToken()
	//	    innerFactor = self.parseExpression()
	//	    token = self.getToken()
	//	   if token.tokenType is not TOKENS.T_R_PAREN:
	//	        raise CritterParseError(token, ")")
	//	    return innerFactor
	//	case TOKENS.T_NUMBER:
	//	    number = self.parseNumber()
	//	    return number
	//	case _:
	//	    raise CritterParseError(token, "<Factor>")
}

func (p *Parser) ParseMemNode() *MemNode {
	token := p.next()
	switch token.TokenType {
	case tMemSize, tDefense, TOffense, tEnergy, tSize, tPass:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseMemNode(), got %s, but handler is not implemented.", token.Lexeme))
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
