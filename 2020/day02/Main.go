package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func countLetter(word string, letter string) int {

	count := 0
	for _, character := range word {

		stringChar := string(character)

		if stringChar == letter {
			count++
		}
	}
	return count
}

func checkValidityPart1(line string) int {

	fields := strings.Fields(line)
	rule := fields[0]
	letter := fields[1]
	pword := fields[2]

	rule2 := strings.Split(rule, "-")
	min, _ := strconv.Atoi(rule2[0])
	max, _ := strconv.Atoi(rule2[1])

	letter2 := strings.Split(letter, ":")

	countOfLetter := countLetter(pword, letter2[0])

	if min <= countOfLetter {
		if countOfLetter <= max {
			return 1
		}
	}
	return 0
}

func checkValidityPart2(line string) int {

	fields := strings.Fields(line)
	rule := fields[0]
	letter := fields[1]
	pword := fields[2]

	rule2 := strings.Split(rule, "-")
	start, _ := strconv.Atoi(rule2[0])
	end, _ := strconv.Atoi(rule2[1])

	letter2 := strings.Split(letter, ":")

	charCount := 0
	for i, character := range pword {

		stringChar := string(character)

		if start-1 == i {
			if stringChar == letter2[0] {
				charCount++
			}
		}
		if end-1 == i {
			if stringChar == letter2[0] {
				charCount++
			}
		}
	}

	if charCount == 1 {
		return 1
	}
	return 0
}

func main() {

	path := pathToFile("input.txt")

	lines, err := readLines(path)
	check(err)

	count := 0
	for _, line := range lines {
		count += checkValidityPart1(line)
	}

	fmt.Print("part 1 ")
	fmt.Println(count)

	count = 0
	for _, line := range lines {
		count += checkValidityPart2(line)
	}

	fmt.Print("part 2 ")
	fmt.Println(count)

}
