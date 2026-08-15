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

func dailyTemperatures(temperatures []int) []int {

	stack := Stack[int]{}
	answer := make([]int, len(temperatures))

	for idx, value := range temperatures {
		if stack.IsEmpty() {
			stack.Push(idx)
			continue
		}
		stackTop, err := stack.Top()
		if err != nil {
			return []int{}
		}
		if value <= temperatures[stackTop] {
			stack.Push(idx)
			continue
		}
		size := stack.Size()
		for i := size - 1; i >= 0; i-- {
			temperatureIdx, err := stack.Top()
			if err != nil {
				return []int{}
			}
			if value > temperatures[temperatureIdx] {
				popedIdx, err := stack.Pop()
				if err != nil {
					return []int{}
				}
				answer[popedIdx] = idx - popedIdx
			} else {
				break
			}
		}
		stack.Push(idx)
	}
	return answer
}
