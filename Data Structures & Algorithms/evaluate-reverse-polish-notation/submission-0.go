import (
	"errors"
	"slices"
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
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("Stack is empty")
	}

	idx := s.Size() - 1
	elem := s.elements[idx]
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

func evalRPN(tokens []string) int {

	operations := []string{"+", "-", "*", "/"}
	stack := Stack[int]{}

	for _, val := range tokens {
		if slices.Contains(operations, val) {
			var res int

			numRight, err := stack.Pop()
			if err != nil {
				//fmt.Println("Error in Pop right ")
				return -1
			}
			numLeft, err := stack.Pop()
			if err != nil {
				//fmt.Println("Error in Pop left")
				return -1
			}

			switch val {
			case "+":
				res = numLeft + numRight
			case "-":
				res = numLeft - numRight
			case "*":
				res = numLeft * numRight
			case "/":
				res = numLeft / numRight
			}
			stack.Push(res)
			continue
		}
		number, err := strconv.Atoi(val)
		if err != nil {
			//fmt.Println("Error in coversion string to int")
			return -1
		}
		stack.Push(number)
	}
	result, err := stack.Top()
	if err != nil {
		//fmt.Println("Error in Top")
		return -1
	}
	return result
}
