package day18

import (
	"strconv"
	"unicode"
)

type operator func(int, int) int

type frame struct {
	value int
	op    operator
}

type stack []*frame

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(f *frame) {
	*s = append(*s, f)
}

func (s *stack) Pop() (*frame, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return element, true
	}
}

func (s *stack) Head() *frame {
	index := len(*s) - 1
	head := (*s)[index]
	return head
}

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func runCalc(source string) int {
	var frames stack
	runes := []rune(source)
	valueBuffer := []rune{}
	expectingOperator := false
	expectingSpace := false

	frames.Push(&frame{
		value: 0,
		op:    add,
	})

	for _, r := range runes {
		if expectingSpace {
			expectingSpace = false
		} else if expectingOperator {
			head := (frames).Head()
			switch r {
			case '+':
				head.op = add
			case '*':
				head.op = multiply
			}
			expectingOperator = false
			expectingSpace = true
		} else {
			if unicode.IsDigit(r) {
				valueBuffer = append(valueBuffer, r)
			} else {
				if len(valueBuffer) > 0 {
					value, _ := strconv.Atoi(string(valueBuffer))
					valueBuffer = []rune{}
					runFrame(&frames, value)
				}
				switch r {
				case ' ':
					expectingOperator = true
				case '(':
					frames.Push(&frame{
						value: 0,
						op:    add,
					})
				case ')':
					popped, _ := frames.Pop()
					value := popped.value
					runFrame(&frames, value)
				}
			}
		}
	}

	if len(valueBuffer) > 0 {
		value, _ := strconv.Atoi(string(valueBuffer))
		valueBuffer = []rune{}
		runFrame(&frames, value)
	}
	result, _ := frames.Pop()
	return result.value
}

func runFrame(frames *stack, value int) {
	fromStack, _ := frames.Pop()
	fromStack.value = fromStack.op(fromStack.value, value)
	frames.Push(fromStack)
}

func runHomework(homework *[]string) (sum int) {
	for _, line := range *homework {
		sum += runCalc(line)
	}
	return
}
