package main

import (
	"fmt"
	"strings"
)

func countAnswersP1(groupedLines [][]string) {

	// init variables
	answerCount := 0
	answers := make(map[string]bool)

	// for each group
	for _, group := range groupedLines {
		answers = make(map[string]bool)

		// for each line in group
		for _, line := range group {
			// split individual characters
			answer := strings.Split(line, "")

			// for each answer
			for _, character := range answer {
				answers[character] = true
			}
		}
		answerCount += len(answers)
	}

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

func countAnswersP2(groupedLines [][]string) {

	// init variables
	answerCount := 0
	groupSize := 0
	answers := make(map[string]int)

	// for each group
	for _, group := range groupedLines {
		answers = make(map[string]int)

		// for each line in group
		for i, line := range group {
			groupSize = i + 1

			// split individual characters
			answer := strings.Split(line, "")

			// for each answer
			for _, character := range answer {
				answers[character] = answers[character] + 1
			}
		}
		answerCount += genCount(answers, groupSize)
	}

	fmt.Println("P2b sum counts: ", answerCount)
}

func main() {

	path := pathToFile("input.txt")

	lines, err := readLines(path)
	check(err)

	groupedLines := groupLines(lines)

	countAnswersP1(groupedLines)
	countAnswersP2(groupedLines)
}
