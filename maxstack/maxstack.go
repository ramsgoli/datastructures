package maxstack

import (
	"errors"
)

type MaxStack struct {
	stack []int
	maxes []int
}

// Utility function to find the max of two ints
func max(a int, b int) int {
	if a < b {
		return b
	}
	return a
}

// Push the argument onto the stack
func (s *MaxStack) Push(val int) {
	s.stack = append(s.stack, val)

	if s.maxes == nil {
		s.maxes = []int{val}
	} else {
		s.maxes = append(s.maxes, max(s.maxes[len(s.maxes)-1], val))
	}
}

// Pop the top most value off of the stack
func (s *MaxStack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("stack has 0 length")
	}
	poppedVal := s.stack[len(s.stack)-1]

	s.stack = s.stack[:len(s.stack)-1]
	s.maxes = s.maxes[:len(s.maxes)-1]

	return poppedVal, nil
}

// Max returns the maximum value present in the stack
func (s *MaxStack) Max() int {
	return s.maxes[len(s.maxes)-1]
}
