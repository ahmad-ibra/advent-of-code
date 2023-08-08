package main

import (
	in "aoc2022/input"
)

func solution(lines []string) {

}

func main() {
	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
}
