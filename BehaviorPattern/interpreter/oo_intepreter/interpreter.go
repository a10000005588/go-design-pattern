package main

import (
	"strconv"
	"strings"
)

const (
	SUM = "sum"
	SUB = "sub"
)

type Interpreter interface {
	Read() int
}

type value int

func (v *value) Read() int {
	return int(*v)
}

type operationSum struct {
	Left  Interpreter
	Right Interpreter
}

func (a *operationSum) Read() int {
	return a.Left.Read() + a.Right.Read()
}

type operationSubtract struct {
	Left  Interpreter
	Right Interpreter
}

func (s *operationSubtract) Read() int {
	return s.Left.Read() - s.Right.Read()
}

func operatorFactory(o string, left, right Interpreter) Interpreter {
	switch o {
	case SUM:
		return &operationSum{
			Left:  left,
			Right: right,
		}
	case SUB:
		return &operationSubtract{
			Left:  left,
			Right: right,
		}
	}

	return nil
}

// polishNotationStack store the type which is implement Read() of Interpreter interface
type polishNotationStack []Interpreter

func (p *polishNotationStack) Push(s Interpreter) {
	*p = append(*p, s)
}

func (p *polishNotationStack) Pop() Interpreter {
	length := len(*p)

	if length > 0 {
		temp := (*p)[length-1]
		*p = (*p)[:length-1]
		return temp
	}

	return nil
}

func main() {
	stack := polishNotationStack{}
	operators := strings.Split("3 4 sum 2 sub", " ")

	for _, operatorString := range operators {
		if operatorString == SUM || operatorString == SUB {
			right := stack.Pop()
			left := stack.Pop()
			// via Facotry to generate the corrsponding operation.
			// both left and right struct have implemented the Read() of Interpreter interface.
			mathFunc := operatorFactory(operatorString, left, right)
			// make the result to be the "value" type.
			res := value(mathFunc.Read())
			// value has implemented the Read(), thus it can be stored in polishNotationStack{} which store the Interpreter type.
			stack.Push(&res)
		} else {
			// convert string number to integer number
			val, err := strconv.Atoi(operatorString)
			if err != nil {
				panic(err)
			}
			temp := value(val)
			stack.Push(&temp)
		}
	}

	println(int(stack.Pop().Read()))
}
