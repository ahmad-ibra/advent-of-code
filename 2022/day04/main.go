package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type section struct {
	start int
	end   int
}

func fullyContained(s1, s2 *section) bool {
	if s1.start <= s2.start && s1.end >= s2.end {
		return true
	}
	if s2.start <= s1.start && s2.end >= s1.end {
		return true
	}
	return false
}

func partiallyContained(s1, s2 *section) bool {
	if s1.end < s2.start || s2.end < s1.start {
		return false
	}
	return true
}

func newSection(s string) *section {
	sections := strings.Split(s, "-")
	startSection := sections[0]
	endSection := sections[1]

	sNum, _ := strconv.Atoi(startSection)
	eNum, _ := strconv.Atoi(endSection)

	return &section{start: sNum, end: eNum}
}

func partA(lines []string) {

	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		pairs := strings.Split(line, ",")
		s1 := newSection(pairs[0])
		s2 := newSection(pairs[1])

		if fullyContained(s1, s2) {
			count++
		}
	}

	fmt.Printf("\nThere are %v sections that are fully contained\n", count)

}

func partB(lines []string) {
	count := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		pairs := strings.Split(line, ",")
		s1 := newSection(pairs[0])
		s2 := newSection(pairs[1])

		if partiallyContained(s1, s2) {
			count++
		}
	}

	fmt.Printf("\nThere are %v sections that are partially contained\n", count)
}

func main() {

	lines, err := inputSplitByLine("day04/input.txt")
	if err != nil {
		panic(err)
	}

	partA(lines)
	partB(lines)

}

func inputSplitByLine(fileLoc string) ([]string, error) {

	dat, err := os.ReadFile(fileLoc)
	if err != nil {
		return nil, err
	}

	groups := strings.Split(string(dat), "\n")

	return groups, nil
}
