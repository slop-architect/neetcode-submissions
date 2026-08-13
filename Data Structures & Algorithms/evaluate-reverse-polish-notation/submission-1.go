import (
	"errors"
)

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}

func (s *Stack[T]) Push(val T) {
	s.elements = append(s.elements, val)
}

func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("Stack is empty")
	}

	idx := s.Size() - 1
	elem := s.elements[idx]
	s.elements[idx] = zero
	s.elements = s.elements[:idx]
	return elem, nil
}

func (s *Stack[T]) Top() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	return s.elements[s.Size()-1], nil

}

func evalLeftRight(operation string, stack *Stack[int]) int {
	numRight, err := stack.Pop()
	if err != nil {
		return -1
	}
	numLeft, err := stack.Pop()
	if err != nil {
		return -1
	}

	switch operation {
	case "+":
		return numLeft + numRight
	case "-":
		return numLeft - numRight
	case "*":
		return numLeft * numRight
	case "/":
		return numLeft / numRight
	}
	return 0
}

func evalRPN(tokens []string) int {

	stack := Stack[int]{}

	for _, val := range tokens {
		var res int
		switch val {
		case "+", "-", "*", "/":
			res = evalLeftRight(val, &stack)
			stack.Push(res)
		default:
			number, err := strconv.Atoi(val)
			if err != nil {
				return -1
			}
			stack.Push(number)
		}
	}
	result, err := stack.Top()
	if err != nil {
		return -1
	}
	return result
}

