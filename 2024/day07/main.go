package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	answer int
	nums   []int
}

// checks if it can be solved with only addition and multiplication
func (e Equation) isSolvablePlusTimes() bool {
	// impossible if there are no numbers
	if len(e.nums) == 0 {
		return false
	}

	// check if its solved
	if len(e.nums) == 1 {
		return e.answer == e.nums[0]
	}

	// answer is less than the smallest number
	if e.answer < e.nums[0] {
		return false
	}

	// add
	nums := []int{}
	nums = append(nums, e.nums[0]+e.nums[1])
	nums = append(nums, e.nums[2:]...)
	eAdd := Equation{answer: e.answer, nums: nums}

	// multiply
	nums = []int{}
	nums = append(nums, e.nums[0]*e.nums[1])
	nums = append(nums, e.nums[2:]...)
	eMultiply := Equation{answer: e.answer, nums: nums}

	return eAdd.isSolvablePlusTimes() || eMultiply.isSolvablePlusTimes()
}

// checks if it can be solved with only addition, multiplication, and concatenation
func (e Equation) isSolvablePlusTimesConcat() bool {
	// impossible if there are no numbers
	if len(e.nums) == 0 {
		return false
	}

	// check if its solved
	if len(e.nums) == 1 {
		return e.answer == e.nums[0]
	}

	// answer is less than the smallest number
	if e.answer < e.nums[0] {
		return false
	}

	// add
	nums := []int{}
	nums = append(nums, e.nums[0]+e.nums[1])
	nums = append(nums, e.nums[2:]...)
	eAdd := Equation{answer: e.answer, nums: nums}

	// multiply
	nums = []int{}
	nums = append(nums, e.nums[0]*e.nums[1])
	nums = append(nums, e.nums[2:]...)
	eMultiply := Equation{answer: e.answer, nums: nums}

	// concatenate
	nums = []int{}
	num1Str := strconv.Itoa(e.nums[0]) + strconv.Itoa(e.nums[1])
	num1, err := strconv.Atoi(num1Str)
	if err != nil {
		panic(err)
	}
	nums = append(nums, num1)
	nums = append(nums, e.nums[2:]...)
	eConcat := Equation{answer: e.answer, nums: nums}

	return eAdd.isSolvablePlusTimesConcat() || eMultiply.isSolvablePlusTimesConcat() || eConcat.isSolvablePlusTimesConcat()
}

func parseInput(input string) []Equation {
	equations := []Equation{}

	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		if line == "" {
			continue
		}

		equation := Equation{
			nums: []int{},
		}

		parts := strings.Split(line, ":")
		num, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		equation.answer = num

		for _, numStr := range strings.Split(parts[1], " ") {
			if numStr == "" {
				continue
			}
			num, err := strconv.Atoi(numStr)
			if err != nil {
				panic(err)
			}

			equation.nums = append(equation.nums, num)
		}

		equations = append(equations, equation)
	}

	return equations
}

func main() {
	equations := parseInput("input.txt")

	ansP1 := 0
	ansP2 := 0
	for _, equation := range equations {
		if equation.isSolvablePlusTimes() {
			ansP1 += equation.answer
		}
		if equation.isSolvablePlusTimesConcat() {
			ansP2 += equation.answer
		}
	}

	fmt.Println("Part 1:", ansP1)
	fmt.Println("Part 2:", ansP2)
}
