package stack

import (
	"fmt"
)

// TODO: make stack generic

type Stack struct {
	st []string
}

func NewStack() *Stack {
	return &Stack{st: make([]string, 0)}
}

func (s *Stack) Print() {
	fmt.Printf("Stack size is %v\n", len(s.st))
	fmt.Println(s.st)
}

func (s *Stack) Push(str string) {
	s.st = append(s.st, str)
}

func (s *Stack) Pop() string {
	length := len(s.st)
	if length > 0 {
		str := s.st[length-1]
		s.st = s.st[:length-1]
		return str
	}
	return ""
}

func (s *Stack) Peek() string {
	length := len(s.st)
	if length > 0 {
		return s.st[length-1]
	}
	return ""
}

func (s *Stack) Reverse() {
	st := s.st
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

// TODO: make this an actual test that lives in the stack package
func testStack() {
	s := NewStack()
	s.Print()
	s.Push("H")
	s.Print()
	fmt.Printf("value of peek is %v\n", s.Peek())
	s.Print()
	b := s.Pop()
	fmt.Printf("popped item is %v\n", b)
	s.Print()
	s.Push("a")
	s.Push("b")
	s.Push("c")
	s.Push("d")
	s.Push("e")
	s.Print()
	s.Reverse()
	s.Print()
}
