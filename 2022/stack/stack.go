package stack

import (
	"fmt"
)

type Stack struct {
	crates []string
}

func NewStack() *Stack {
	return &Stack{crates: make([]string, 0)}
}

func (s *Stack) Print() {
	fmt.Printf("Stack size is %v\n", len(s.crates))
	fmt.Println(s.crates)
}

func (s *Stack) Push(str string) {
	s.crates = append(s.crates, str)
}

func (s *Stack) Pop() string {
	length := len(s.crates)
	if length > 0 {
		str := s.crates[length-1]

		s.crates = s.crates[:length-1]
		return str
	}
	return ""
}

func (s *Stack) Peek() string {
	length := len(s.crates)
	if length > 0 {
		return s.crates[length-1]
	}
	return ""
}

func (s *Stack) Reverse() {
	st := s.crates
	start := 0
	end := len(st) - 1

	for start <= end {
		tmp := st[start]
		st[start] = st[end]
		st[end] = tmp
		start++
		end--
	}
}
