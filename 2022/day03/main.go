package main

import (
	"fmt"
	"os"
	"strings"
)

type rucksack struct {
	left       map[string]int
	right      map[string]int
	duplicated string
}

func newRuckSack(items string) *rucksack {
	itemsPerSide := len(items) / 2

	l := make(map[string]int)
	r := make(map[string]int)
	duplicated := ""

	compartment := l

	i := 0
	for _, c := range items {
		stringChar := string(c)
		if i == itemsPerSide {
			//switch to right compartment
			compartment = r
		}
		if val, ok := compartment[stringChar]; ok {
			compartment[stringChar] = val + 1
		} else {
			compartment[stringChar] = 1
		}

		if i >= itemsPerSide {
			if _, ok := l[stringChar]; ok {
				duplicated = stringChar
			}
		}
		i++
	}

	return &rucksack{left: l, right: r, duplicated: duplicated}
}

func toPriority(s string) int {
	char := int(s[0])
	if 'a' <= char && char <= 'z' {
		// Lowercase letter (a-z): priority 1-26
		return int(char - 'a' + 1)
	} else if 'A' <= char && char <= 'Z' {
		// Uppercase letter (A-Z): priority 27-52
		return int(char - 'A' + 27)
	}
	return -1
}

func partA(lines []string) {
	score := 0
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		rs := newRuckSack(line)

		score += toPriority(rs.duplicated)
	}
	fmt.Printf("\nScore is %v\n", score)
}

type itemsSet struct {
	items map[string]struct{}
}

func newItemSet(items string) *itemsSet {
	s := make(map[string]struct{})
	for _, c := range items {
		s[string(c)] = struct{}{}
	}

	return &itemsSet{items: s}
}

func (is *itemsSet) commonItems(is2 *itemsSet) *itemsSet {
	s := make(map[string]struct{})

	for key := range is.items {
		if _, ok := is2.items[key]; ok {
			s[key] = struct{}{}
		}
	}

	return &itemsSet{items: s}
}

func partB(lines []string) {
	score := 0
	itemSets := make([]*itemsSet, 3)

	for i, line := range lines {
		group := i % 3
		if len(line) == 0 {
			continue
		}

		is := newItemSet(line)
		itemSets[group] = is

		if group == 2 {
			is0 := itemSets[0]
			is1 := itemSets[1]
			is2 := itemSets[2]

			common1 := is0.commonItems(is1)
			commonFinal := common1.commonItems(is2)

			for key := range commonFinal.items {
				score += toPriority(key)
			}
		}
	}

	fmt.Printf("\nScore is %v\n", score)

}

func main() {

	lines, err := inputSplitByLine("day03/input.txt")
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
