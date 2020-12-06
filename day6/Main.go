package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// IO -------------------------------------------
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

//-----------------------------------------------

func countAnswersP1(lines []string) {

	// init variables
	newGroup := true
	answerCount := 0
	answers := make(map[string]bool)

	for _, line := range lines {
		if newGroup {
			answers = make(map[string]bool)
			newGroup = false
		}

		// split individual characters
		answer := strings.Split(line, "")

		for _, character := range answer {
			answers[character] = true
		}

		// if end of group
		if line == "" {
			newGroup = true
			answerCount += len(answers)
		}
	}

	// add the last group to the map
	answerCount += len(answers)

	fmt.Println("P1 sum counts: ", answerCount)
}

func genCount(answers map[string]int, size int) int {

	count := 0
	for _, value := range answers {
		if value == size {
			count++
		}
	}

	return count
}

func countAnswersP2(lines []string) {

	// init variables
	newGroup := true
	answerCount := 0
	groupSize := 0
	answers := make(map[string]int)

	for _, line := range lines {
		if newGroup {
			answers = make(map[string]int)
			newGroup = false
			groupSize = 0
		}

		// split individual characters
		answer := strings.Split(line, "")

		for _, character := range answer {
			answers[character] = answers[character] + 1
		}

		// if end of group
		if line == "" {
			newGroup = true
			answerCount += genCount(answers, groupSize)

		} else {
			groupSize++
		}
	}

	// add the last group to the map
	answerCount += genCount(answers, groupSize)

	fmt.Println("P2 sum counts: ", answerCount)
}

func main() {

	path := pathToFile("input.txt")

	lines, err := readLines(path)
	check(err)

	countAnswersP1(lines)
	countAnswersP2(lines)
}
