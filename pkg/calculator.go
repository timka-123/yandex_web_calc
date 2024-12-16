package pkg

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"text/scanner"
)

const (
	TokenTypeOperand  = 1
	TokenTypeOperator = 2
)

var priorities map[string]float64
var associativities map[string]bool

func init() {
	priorities = make(map[string]float64, 0)
	associativities = make(map[string]bool, 0)

	priorities["+"] = 0
	priorities["-"] = 0
	priorities["*"] = 1
	priorities["/"] = 1
}

type Token struct {
	Type  int
	Value interface{}
}

func tryGetOperand(str string) *Token {
	value, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		return nil
	}
	return NewOperandToken(value)
}

func NewOperandToken(val float64) *Token {
	return NewToken(val, TokenTypeOperand)
}

func NewOperatorToken(val string) *Token {
	return NewToken(val, TokenTypeOperator)
}

func NewToken(val interface{}, typ int) *Token {
	return &Token{Value: val, Type: typ}
}

func (token *Token) IsOperand(val float64) bool {
	return token.Type == TokenTypeOperand && token.Value.(float64) == val
}

func (token *Token) IsOperator(val string) bool {
	return token.Type == TokenTypeOperator && token.Value.(string) == val
}

func Scan(input string) ([]string, error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(input))

	var tok rune
	var result = make([]string, 0)
	for tok != scanner.EOF {
		tok = s.Scan()
		value := strings.TrimSpace(s.TokenText())
		if len(value) > 0 {
			result = append(result, s.TokenText())
		}
	}
	return result, nil
}

func Parse(tokens []string) ([]*Token, error) {
	var ret []*Token

	var operators []string
	for _, token := range tokens {
		operandToken := tryGetOperand(token)
		if operandToken != nil {
			ret = append(ret, operandToken)
		} else {
			if token == "(" {
				operators = append(operators, token)
			} else if token == ")" {
				foundLeftParenthesis := false
				for len(operators) > 0 {
					oper := operators[len(operators)-1]
					operators = operators[:len(operators)-1]
					if oper == "(" {
						foundLeftParenthesis = true
						break
					} else {
						ret = append(ret, NewOperatorToken(oper))
					}
				}
				if !foundLeftParenthesis {
					return nil, errors.New("error in expression")
				}
			} else {
				priority, ok := priorities[token]
				if !ok {
					return nil, fmt.Errorf("unknown operator: %v", token)
				}
				rightAssociative := associativities[token]
				for len(operators) > 0 {
					top := operators[len(operators)-1]
					if top == "(" {
						break
					}
					prevPriority := priorities[top]
					if (rightAssociative && priority < prevPriority) || (!rightAssociative && priority <= prevPriority) {
						operators = operators[:len(operators)-1]
						ret = append(ret, NewOperatorToken(top))
					} else {
						break
					}
				}
				operators = append(operators, token)
			}
		}
	}

	for len(operators) > 0 {
		operator := operators[len(operators)-1]
		operators = operators[:len(operators)-1]
		if operator == "(" {
			return nil, errors.New("error in expression")
		}
		ret = append(ret, NewOperatorToken(operator))
	}
	return ret, nil
}

func Evaluate(tokens []*Token) (float64, error) {
	if tokens == nil {
		return 0, errors.New("tokens cannot be nil")
	}
	var stack []float64
	for _, token := range tokens {
		if token.Type == TokenTypeOperand {
			val := token.Value.(float64)
			stack = append(stack, val)
		} else {
			if len(stack) < 2 {
				return 0, errors.New("missing operand")
			}
			arg1, arg2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			val, err := evaluateOperator(token.Value.(string), arg1, arg2)
			if err != nil {
				return 0, err
			}
			stack = append(stack, val)
		}
	}
	if len(stack) != 1 {
		return 0, errors.New("stack corrupted")
	}
	return stack[len(stack)-1], nil
}

func evaluateOperator(oper string, a, b float64) (float64, error) {
	switch oper {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("division by zero")
		}
		return a / b, nil
	default:
		return 0, errors.New("unknown operator: " + oper)
	}
}

func Calc(expression string) (float64, error) {
	tokens, err := Scan(expression)
	if err != nil {
		return 0, err
	}
	parsed, err := Parse(tokens)
	if err != nil {
		return 0, err
	}
	return Evaluate(parsed)
}
