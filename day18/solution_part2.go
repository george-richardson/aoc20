package day18

import (
	"strconv"
	"unicode"
)

const addition rune = '+'
const multiplication rune = '*'

type expression interface {
	calculate() int
}

type literal int

func (l literal) calculate() int {
	return int(l)
}

type compoundExpression struct {
	expressions *[]expression
	operations  *[]rune
}

func (c compoundExpression) calculate() int {
	ops := *c.operations
	opsl := len(ops)
	exps := *c.expressions
	for i := opsl - 1; i >= 0; i-- {
		if ops[i] == addition {
			left := exps[i]
			right := exps[i+1]
			newLiteral := literal(add(left.calculate(), right.calculate()))
			ops = append(ops[0:i], ops[i+1:]...)
			exps = append(exps[0:i], append([]expression{expression(newLiteral)}, exps[i+2:]...)...)
		}
	}

	opsl = len(ops)
	for i := opsl - 1; i >= 0; i-- {
		if ops[i] == multiplication {
			left := exps[i]
			right := exps[i+1]
			newLiteral := literal(multiply(left.calculate(), right.calculate()))
			ops = append(ops[0:i], ops[i+1:]...)
			exps = append(exps[0:i], append([]expression{expression(newLiteral)}, exps[i+2:]...)...)
		}
	}

	return exps[0].calculate()
}

type expressionStack []*compoundExpression

func (s *expressionStack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *expressionStack) Push(f *compoundExpression) {
	*s = append(*s, f)
}

func (s *expressionStack) Pop() (*compoundExpression, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *expressionStack) Head() *compoundExpression {
	index := len(*s) - 1
	head := (*s)[index]
	return head
}

func parseCalc(source string) *compoundExpression {
	var frames expressionStack
	runes := []rune(source)
	valueBuffer := []rune{}
	expectingOperator := false
	expectingSpace := false

	frames.Push(&compoundExpression{
		expressions: &[]expression{},
		operations:  &[]rune{},
	})

	for _, r := range runes {
		if expectingSpace {
			expectingSpace = false
		} else if expectingOperator {
			head := (frames).Head()
			switch r {
			case '+':
				newOps := append(*head.operations, addition)
				head.operations = &newOps
			case '*':
				newOps := append(*head.operations, multiplication)
				head.operations = &newOps
			}
			expectingOperator = false
			expectingSpace = true
		} else {
			if unicode.IsDigit(r) {
				valueBuffer = append(valueBuffer, r)
			} else {
				if len(valueBuffer) > 0 {
					head := (frames).Head()
					value, _ := strconv.Atoi(string(valueBuffer))
					valueBuffer = []rune{}
					newExps := append(*head.expressions, literal(value))
					head.expressions = &newExps
				}
				switch r {
				case ' ':
					expectingOperator = true
				case '(':
					frames.Push(&compoundExpression{
						expressions: &[]expression{},
						operations:  &[]rune{},
					})
				case ')':
					popped, _ := frames.Pop()
					head := (frames).Head()
					newExps := append(*head.expressions, popped)
					head.expressions = &newExps
				}
			}
		}
	}

	if len(valueBuffer) > 0 {
		value, _ := strconv.Atoi(string(valueBuffer))
		valueBuffer = []rune{}
		newExps := append(*frames.Head().expressions, literal(value))
		frames.Head().expressions = &newExps
	}
	result, _ := frames.Pop()
	return result
}

func runCalcPart2(source string) int {
	exp := *parseCalc(source)
	return exp.calculate()
}

func runHomeworkPart2(homework *[]string) (sum int) {
	for _, line := range *homework {
		sum += runCalcPart2(line)
	}
	return
}
