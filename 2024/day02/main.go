package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fbyte, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(fbyte), "\n")

	safeLevels := 0
	safeLevelsDampened := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		lineLevelsStr := strings.Split(line, " ")
		lineLevelsInt := make([]int, len(lineLevelsStr))
		for i, v := range lineLevelsStr {
			num, err := strconv.Atoi(v)
			if err != nil {
				panic(err)
			}

			lineLevelsInt[i] = num
		}

		if isSafe(lineLevelsInt) {
			safeLevels++
		}

		if isSafe(lineLevelsInt) {
			safeLevelsDampened++
		} else {
			if isSafeLevelDampened(lineLevelsInt) {
				safeLevelsDampened++
			}
		}
	}

	fmt.Println("Part 1:", safeLevels)
	fmt.Println("Part 2:", safeLevelsDampened)
}

func isSafe(levels []int) bool {
	// must be all increasing or all decreasing
	if levels[1] > levels[0] {
		return isSafeIncreasing(levels)
	}

	if levels[1] < levels[0] {
		return isSafeDecreasing(levels)
	}

	return false
}

func isSafeIncreasing(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		if levels[i]-levels[i-1] > 3 {
			return false
		}

		if levels[i]-levels[i-1] < 1 {
			return false
		}
	}
	return true
}

func isSafeDecreasing(levels []int) bool {
	// must be all decreasing by 1 2 or 3
	for i := 1; i < len(levels); i++ {
		if levels[i]-levels[i-1] < -3 {
			return false
		}

		if levels[i]-levels[i-1] > -1 {
			return false
		}
	}
	return true
}

func isSafeLevelDampened(levels []int) bool {
	for i := 0; i < len(levels); i++ {
		// put everything before index i in levelsCopy
		levelsLeft := levels[:i]
		levelsRight := levels[i+1:]

		newLevels := make([]int, len(levelsLeft)+len(levelsRight))
		copy(newLevels, levelsLeft)
		copy(newLevels[len(levelsLeft):], levelsRight)

		if isSafe(newLevels) {
			return true
		}
	}

	return false
}
