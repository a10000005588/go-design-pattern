package interpreter

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
	MUL = "mul"
	DIV = "div"
)

func getOperationFunc(o string) func(a, b int) int {
	switch o {
	case SUM:
		return func(a, b int) int {
			return a + b
		}
	case SUB:
		return func(a, b int) int {
			return a - b
		}
	case MUL:
		return func(a, b int) int {
			return a * b
		}
	case DIV:
		return func(a, b int) int {
			return a / b
		}
	}

	return nil
}

func isOperator(o string) bool {
	if o == SUM || o == SUB || o == MUL || o == DIV {
		return true
	}

	return false
}

type polishNotationStack []int

func (p *polishNotationStack) Push(s int) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() int {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return 0
}

func Calculate(o string) (int, error) {
	// create a stack data structure storing the operator and value
	stack := polishNotationStack{}
	// The string will be "3 4 sum 2 sub"
	// Thus, split it with a space " "
	operators := strings.Split(o, " ")

	for _, operatorString := range operators {
		// check the operators is "sum" "sub" "mul" "div" or not
		if isOperator(operatorString) {
			// is operator, then pop two integers from stack
			right := stack.Pop()
			left := stack.Pop()
			// check what kind of math operation.
			mathFunc := getOperationFunc(operatorString)
			res := mathFunc(left, right)
			// put the result in the stack
			stack.Push(res)
		} else {
			// strcov.Atoi is convert string number to int
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				return 0, err
			}
			// push the integer into stack
			stack.Push(val)
		}
	}

	return int(stack.Pop()), nil
}
