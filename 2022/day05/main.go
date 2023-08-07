package main

import (
	"fmt"
	"strconv"
	"strings"

	in "aoc2022/input"
	st "aoc2022/stack"
)

type instruction struct {
	count int
	from  int
	to    int
}

func newInstruction(s string) *instruction {
	parts := strings.Split(s, " ")
	c, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}
	f, err := strconv.Atoi(parts[3])
	if err != nil {
		panic(err)
	}
	t, err := strconv.Atoi(parts[5])
	if err != nil {
		panic(err)
	}
	return &instruction{count: c, from: f - 1, to: t - 1}
}

func partB(lines []string) {
	stackCount := ((len(lines[0]) - 3) / 4) + 1
	stacks := make([]st.Stack, stackCount)

	parsingStartPos := true
	for _, line := range lines {
		if line == "" {
			if parsingStartPos {
				for _, s := range stacks {
					s.Reverse()
				}
			}
			parsingStartPos = false
			continue
		}
		if parsingStartPos {
			for i := 0; i < stackCount; i++ {
				letter := string(line[mapToIndex(i)])
				if letter == " " {
					continue
				}

				_, err := strconv.Atoi(letter)
				if err == nil {
					continue
				}

				stacks[i].Push(letter)
			}
			continue
		}
		inst := newInstruction(line)
		from := &stacks[inst.from]
		to := &stacks[inst.to]
		tmp := st.NewStack()
		for i := 0; i < inst.count; i++ {
			tmp.Push(from.Pop())
		}
		for i := 0; i < inst.count; i++ {
			to.Push(tmp.Pop())
		}
	}

	fmt.Println("Part B:")
	for _, s := range stacks {
		fmt.Print(s.Peek())
	}
	fmt.Println()
}

func mapToIndex(i int) int {
	// 0 -> 1
	// 1 -> 5
	// 2 -> 9
	// etc
	return (4 * i) + 1
}

func partA(lines []string) {
	stackCount := ((len(lines[0]) - 3) / 4) + 1
	stacks := make([]st.Stack, stackCount)

	parsingStartPos := true
	for _, line := range lines {
		if line == "" {
			if parsingStartPos {
				for _, s := range stacks {
					s.Reverse()
				}
			}
			parsingStartPos = false
			continue
		}
		if parsingStartPos {
			for i := 0; i < stackCount; i++ {
				letter := string(line[mapToIndex(i)])
				if letter == " " {
					continue
				}

				_, err := strconv.Atoi(letter)
				if err == nil {
					continue
				}

				stacks[i].Push(letter)
			}
			continue
		}
		inst := newInstruction(line)
		from := &stacks[inst.from]
		to := &stacks[inst.to]
		for i := 0; i < inst.count; i++ {
			to.Push(from.Pop())
		}
	}

	fmt.Println("Part A:")
	for _, s := range stacks {
		fmt.Print(s.Peek())
	}
	fmt.Println()
}

func main() {

	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	partA(lines)
	partB(lines)
}
