package main

import (
	in "aoc2022/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type xCycle struct {
	num    int
	during int
	after  int
}

func solution(lines []string) {
	sigs := make([]xCycle, 0)

	var prev xCycle = xCycle{num: 0, during: 1, after: 1}
	var cur xCycle

	cycleNum := 1
	for _, line := range lines {
		if line != "" {
			if line == "noop" {
				cur = xCycle{num: cycleNum, during: prev.after, after: prev.after}
				sigs = append(sigs, cur)
				cycleNum++
				prev = cur
			} else {
				cur = xCycle{num: cycleNum, during: prev.after, after: prev.after}
				sigs = append(sigs, cur)
				cycleNum++
				prev = cur

				parts := strings.Split(line, " ")
				add, err := strconv.Atoi(parts[1])
				if err != nil {
					panic(err)
				}
				cur = xCycle{num: cycleNum, during: prev.after, after: prev.after + add}
				sigs = append(sigs, cur)
				cycleNum++
				prev = cur
			}
		}
	}

	str := 0
	for i := 0; i < len(sigs); i++ {
		cur := sigs[i]
		if (cur.num-20)%40 == 0 {
			str += cur.num * cur.during
		}
	}

	fmt.Printf("Part A: Sum of signals is %v\n", str)

	fmt.Println("Part B:")
	for j := 0; j < len(sigs); j++ {
		if j%40 == 0 {
			fmt.Println()
		}

		c := sigs[j]

		if math.Abs(float64((j%40)-c.during)) <= 1 {
			fmt.Printf("#")
		} else {
			fmt.Print(".")
		}
	}
}

func main() {
	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
}
