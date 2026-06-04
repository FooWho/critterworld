package parser

import (
	"fmt"
	"strconv"
)

type Parser struct {
	tokens     []LexedToken
	cursor     int
	tokenCount int
}

func NewParser(input []LexedToken) *Parser {
	var p Parser = Parser{tokens: input, cursor: 0, tokenCount: len(input)}
	return &p
}

func (p *Parser) peek() LexedToken {
	return p.tokens[p.cursor]
}

func (p *Parser) next() LexedToken {
	token := p.tokens[p.cursor]
	p.cursor++
	return token
}

func (p *Parser) Parse() *Program {
	program := Program{}
	program.rules = make([]*Rule, 0)
	for p.cursor < p.tokenCount {
		rule := p.ParseRule()
		program.rules = append(program.rules, rule)
	}
	return &program
}

func (p *Parser) ParseRule() *Rule {
	rule := Rule{}
	rule.condition = p.ParseCondition()

	return &rule
}

func (p *Parser) ParseCondition() BooleanOperator {
	conjunction := p.ParseConjunction()

	for p.peek().TokenType == T_OR {
		op := p.next()
		logOp := LogicalOperator{}
		logOp.leftOperand = conjunction
		logOp.operator = op
		logOp.rightOperand = p.ParseConjunction()
		conjunction = &logOp
	}

	return conjunction
}

func (p *Parser) ParseConjunction() BooleanOperator {
	relation := p.ParseRelation()
	for p.peek().TokenType == T_AND {
		op := p.next()
		logOp := LogicalOperator{}
		logOp.leftOperand = relation
		logOp.operator = op
		logOp.rightOperand = p.ParseRelation()
		relation = &logOp
	}
	return relation
}

func (p *Parser) ParseRelation() BooleanOperator {
	if p.peek().TokenType == T_L_BRACE {
		_ = p.next()
		condition := p.ParseCondition()
		if p.peek().TokenType != T_R_BRACE {
			panic(fmt.Sprintf("critterworld: parse error: expected '}' in (p *Parser).ParseRelation(), got %s", p.peek().Lexeme))
		}
		_ = p.next()
		return condition
	}

	leftOperand := p.ParseExpression()
	operator := p.next()
	if operator.TokenType != T_GREAT && operator.TokenType != T_GEQU && operator.TokenType != T_EQU &&
		operator.TokenType != T_LEQU && operator.TokenType != T_LESS && operator.TokenType != T_NEQU {
		panic(fmt.Sprintf("critterworld: parse error: expected RelationalOperator in (p *Parser).ParseRelation(), got %s", p.peek().Lexeme))
	}
	rightOperand := p.ParseExpression()
	relOp := RelationalOperator{operator: operator, leftOperand: leftOperand, rightOperand: rightOperand}
	return &relOp
}

func (p *Parser) ParseExpression() Expression {
	expression := p.ParseTerm()
	operator := p.peek()
	for operator.TokenType == T_PLUS || operator.TokenType == T_MINUS {
		operator := p.next()
		binOp := BinaryOperator{}
		binOp.leftOperand = expression
		binOp.operator = operator
		binOp.rightOperand = p.ParseTerm()
		expression = &binOp
		operator = p.peek()
	}
	return expression
}

func (p *Parser) ParseTerm() Expression {
	term := p.ParseFactor()
	operator := p.peek()
	for operator.TokenType == T_STAR || operator.TokenType == T_DIV || operator.TokenType == T_MOD {
		operator := p.next()
		binOp := BinaryOperator{}
		binOp.leftOperand = term
		binOp.operator = operator
		binOp.rightOperand = p.ParseFactor()
		term = &binOp
		operator = p.peek()
	}
	return term
}

func (p *Parser) ParseFactor() Expression {
	token := p.peek()
	switch token.TokenType {
	case T_MEM, T_MEMSIZE, T_DEFENSE, T_OFFENSE, T_SIZE, T_ENERGY, T_PASS, T_POSTURE:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case T_NEARBY, T_AHEAD, T_RANDOM, T_SMELL:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case T_MINUS:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case T_L_PAREN:
		panic(fmt.Sprintf("critterworld: Unimplemented: In (p *Parser).ParseFactor(), got %s, but handler is not implemented.", token.Lexeme))
	case T_NUMBER:
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

func (p *Parser) ParseNumber() Expression {
	token := p.next()
	value, err := strconv.Atoi(token.Lexeme)
	if err != nil {
		panic(fmt.Sprintf("critterworld: parse error in (p *Parser).ParseNumber(), got %s, but expected a number.", token.Lexeme))
	}
	number := Number{value: value}
	return &number
}
