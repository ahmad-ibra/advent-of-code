package main

import (
	"fmt"
	"strings"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func genCount(answers map[string]int, size int) int {

	count := 0
	for _, value := range answers {
		if value == size {
			count++
		}
	}

	return count
}

func countAnswers(groupedLines [][]string) {

	// init variables
	answerCountP1 := 0
	answerCountP2 := 0
	answersP1 := make(map[string]bool)
	answersP2 := make(map[string]int)
	groupSize := 0

	// for each group
	for _, group := range groupedLines {
		answersP1 = make(map[string]bool)
		answersP2 = make(map[string]int)

		// for each line in group
		for i, line := range group {
			groupSize = i + 1

			// split individual characters
			answer := strings.Split(line, "")

			// for each answer
			for _, character := range answer {
				answersP1[character] = true
				answersP2[character] = answersP2[character] + 1
			}
		}
		answerCountP1 += len(answersP1)
		answerCountP2 += genCount(answersP2, groupSize)
	}

	fmt.Println("P1 sum counts: ", answerCountP1)
	fmt.Println("P2 sum counts: ", answerCountP2)
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	groupedLines := input.GroupLines(lines)

	countAnswers(groupedLines)
}
