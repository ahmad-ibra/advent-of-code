package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func pathToFile(fName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)

	return filepath.Join(dir, fName)
}

// im sleepy so im brute forcing this
func bruteForceThreeSum(nums []int) {

	for a := 0; a < len(nums)-2; a++ {
		for b := a + 1; b < len(nums)-1; b++ {
			for c := b + 1; c < len(nums); c++ {
				if nums[a]+nums[b]+nums[c] == 2020 {
					fmt.Println("Success!")
					fmt.Println(nums[a], nums[b], nums[c])
					fmt.Println(nums[a] * nums[b] * nums[c])
					return
				}
			}
		}
	}
	return
}

func twoSum(lines []string) {
	m := make(map[int]bool)

	for i, line := range lines {
		num, _ := strconv.Atoi(line)
		otherNum := 2020 - num
		if m[otherNum] { // if the other number we need is in the map
			fmt.Println("Success!")
			fmt.Println(i, num)
			fmt.Println(i, otherNum)
			fmt.Println(i, num*otherNum)
			fmt.Println("-----------")
		}
		m[num] = true
	}
}

func main() {

	path := pathToFile("src/github.com/Ahmad-Ibra/advent-of-code-2020/day1/input.txt")

	lines, err := readLines(path)
	check(err)

	twoSum(lines)

	// convert slice of strings to slice of ints
	var nums []int
	for _, line := range lines {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}

	bruteForceThreeSum(nums)

	// fmt.Println(result)
}
