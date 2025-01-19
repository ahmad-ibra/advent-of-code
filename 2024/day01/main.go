package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// read the file input.txt and print each line
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	// split data by line
	lines := strings.Split(string(data), "\n")

	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		// split line by space
		parts := strings.Split(line, " ")
		count := len(parts)

		leftInt, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		left = append(left, leftInt)

		rightInt, err := strconv.Atoi(parts[count-1])
		if err != nil {
			panic(err)
		}
		right = append(right, rightInt)
	}

	sort.Ints(left)
	sort.Ints(right)

	diff := make([]int, len(left))
	for i := 0; i < len(left); i++ {
		d := left[i] - right[i]
		if d < 0 {
			d = -d
		}
		diff[i] = d
	}

	sumDiff := 0
	for _, d := range diff {
		sumDiff += d
	}

	fmt.Println("Part 1:", sumDiff)

	similarityScore := 0
	for i := 0; i < len(left); i++ {
		cur := left[i]
		for j := 0; j < len(left); j++ {
			r := right[j]
			if cur < r {
				continue
			}
			if cur > r {
				continue
			}
			similarityScore += cur
		}
	}

	fmt.Println("Part 2:", similarityScore)
}
