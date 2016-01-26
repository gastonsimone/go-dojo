// Provides a Reverse Polish Notation (RPN) evaluator
package rpn

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var operators = "+-*/^!"

func Evalrpn(tokens []string) (float64, error) {
	if len(tokens) <= 0 {
		return 0, errors.New("no tokens to evaluate.")
	}

	var err error
	var x, y float64

	stack := NewStack()
	for i, t := range tokens {

		if strings.Contains(operators, t) {
			x, err = stack.pop()
			if err != nil {
				return 0, fmt.Errorf("processing token %v: %v", i, err)
			}

			y, err = stack.pop()
			if err != nil {
				return 0, fmt.Errorf("processing token %v: %v", i, err)
			}
		}

		switch t {
		case "+":
			stack.push(x + y)
		case "-":
			stack.push(y - x)
		case "*":
			stack.push(x * y)
		case "/":
			stack.push(y / x)
		case "^":
			stack.push(math.Pow(y, x))
		default:
			x, err = strconv.ParseFloat(t, 64)
			if err != nil {
				return 0, fmt.Errorf("processing token %v: %v", i, err)
			}
			stack.push(x)
		}
	}

	return stack.pop()
}

type Stack struct {
	head *stacknode
}

type stacknode struct {
	value float64
	next  *stacknode
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) isEmpty() bool {
	return (s.head == nil)
}

func (s *Stack) push(v float64) {
	s.head = &stacknode{v, s.head}
}

func (s *Stack) pop() (float64, error) {
	if s.isEmpty() {
		return 0, errors.New("cannot pop, stack is empty.")
	}

	popped := s.head
	s.head = s.head.next
	return popped.value, nil
}
