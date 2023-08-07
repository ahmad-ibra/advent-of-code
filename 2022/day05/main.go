package main

import (
	"fmt"
	"os"
	"strings"

	st "aoc2022/stack"
)

func partB(lines []string) {
	for _, line := range lines {
		if line == "" {
			continue
		}
	}
}

func partA(lines []string) {
	stackCount := ((len(lines[0]) - 3) / 4) + 1
	stacks := make([]st.Stack, stackCount)

	parsingStartPos := true
	for _, line := range lines {
		if line == "" {
			parsingStartPos = false
			continue
		}
		if parsingStartPos {
			// parse the positions
			continue
		}
		// handle the instructions
	}

	for _, s := range stacks {
		fmt.Println(s.Peek())
	}
}

func main() {

	testStack()
	lines, err := inputSplitByLine("input-sample.txt")
	if err != nil {
		panic(err)
	}

	partA(lines)
	partB(lines)

}

func testStack() {
	s := st.NewStack()
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

func inputSplitByLine(fileLoc string) ([]string, error) {

	dat, err := os.ReadFile(fileLoc)
	if err != nil {
		return nil, err
	}

	groups := strings.Split(string(dat), "\n")

	return groups, nil
}
