package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

// helpers -----------------------------------------------------
func minMax(lines []string) {
	var intLines = []int{}

	// convert string array to int array
	for _, i := range lines {
		j, _ := strconv.Atoi(i)
		intLines = append(intLines, j)
	}

	sort.Ints(intLines)
	fmt.Println("P2 Sum is: ", intLines[0]+intLines[len(intLines)-1])

}

// tweaking solution from day01 for today
func twoSum(lines []string, goal int) bool {
	m := make(map[int]bool)

	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		otherNum := goal - num
		if m[otherNum] {
			return false
		}
		m[num] = true
	}
	return true
}

//--------------------------------------------------------------

func findContiguousSet(lines []string, goal int) {

	// start by checkout all the numbers
	min := 0
	max := len(lines) - 1
	for {
		sum := 0
		for i := max; i >= min; i-- {
			nextNum, _ := strconv.Atoi(lines[i])
			sum += nextNum
			if sum > goal {
				break
			}
			if sum == goal {
				if i != max {
					// this condition is needed because it considers an array of 1 to be contiguous
					// technically its true, but thats not what we're looking for
					minMax(lines[i : max+1])
					break
				}
			}
		}
		max--
		if min == max {
			break
		}

	}

}

func findNotSum(lines []string, preamble int) int {
	for i := preamble; i < len(lines); i++ {
		// array of preamble
		subSlice := lines[i-preamble : i]
		goal, _ := strconv.Atoi(lines[i])

		isNotSum := twoSum(subSlice, goal)
		if isNotSum {
			fmt.Println("P1 goal: ", goal)
			return goal
		}
	}
	return 0
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	goal := findNotSum(lines, 25)
	findContiguousSet(lines, goal)

}
