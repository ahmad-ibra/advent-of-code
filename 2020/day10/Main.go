package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func strToInt(lines []string) []int {
	var intLines = []int{}
	// convert string array to int array
	for _, i := range lines {
		j, _ := strconv.Atoi(i)
		intLines = append(intLines, j)
	}

	return intLines
}

// could be improved if i pass in a map[int]bool of all the jolts and the highest jolt value
// if that was the case then i wouldn't have needed to sort the input
func useAllAdapters(jolts []int) {
	sum1Jolt := 0
	sum2Jolt := 0
	sum3Jolt := 0

	curCharge := 0

	for i := 0; i < len(jolts); i++ {
		nextCharge := jolts[i]
		if nextCharge-curCharge == 1 {
			sum1Jolt++
		} else if nextCharge-curCharge == 2 {
			sum2Jolt++
		} else if nextCharge-curCharge == 3 {
			sum3Jolt++
		} else {
			fmt.Println("NextCharge: ", nextCharge, ", CurCharge: ", curCharge)
			panic("too large of a dif")
		}
		curCharge = nextCharge

	}
	// built in jolt adapter is 3 higher than the highest rated
	sum3Jolt++

	fmt.Println("sum1Jolt: ", sum1Jolt, ", sum2Jolt: ", sum2Jolt, ", sum3Jolt: ", sum3Jolt)

	fmt.Println("P1: ", sum1Jolt*sum3Jolt)
}

// use dynamic programming algorithm
func countWaysToUseAdapters(joltMap map[int]bool, curCharge int, memo map[int]int) int {
	if memo[curCharge] > 0 {
		return memo[curCharge]
	}

	// base case
	if curCharge == 1 {
		memo[curCharge] = 1
		return 1
	} else if curCharge == 2 {
		// find any numbers that 2 can connect to (ie is there a 1)
		sum := 1
		if joltMap[1] {
			sum += countWaysToUseAdapters(joltMap, 1, memo)
		}
		memo[curCharge] = sum
		return sum
	} else if curCharge == 3 {
		// find any numbers that 3 can connect to (ie is there a 1 or a 2)
		sum := 1
		if joltMap[2] {
			sum += countWaysToUseAdapters(joltMap, 2, memo)
		}
		if joltMap[1] {
			sum += countWaysToUseAdapters(joltMap, 1, memo)
		}
		memo[curCharge] = sum
		return sum
	}

	// recursive case
	sum := 0
	if joltMap[curCharge-1] {
		sum += countWaysToUseAdapters(joltMap, curCharge-1, memo)
	}
	if joltMap[curCharge-2] {
		sum += countWaysToUseAdapters(joltMap, curCharge-2, memo)
	}
	if joltMap[curCharge-3] {
		sum += countWaysToUseAdapters(joltMap, curCharge-3, memo)
	}
	memo[curCharge] = sum
	return sum

}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	intJolts := strToInt(lines)

	sort.Ints(intJolts)

	useAllAdapters(intJolts)

	memo := make(map[int]int)
	joltMap := make(map[int]bool)
	for _, i := range intJolts {
		joltMap[i] = true
	}

	sum := countWaysToUseAdapters(joltMap, intJolts[len(intJolts)-1], memo)
	fmt.Println("P2: ", sum)

}
