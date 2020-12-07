package main

import (
	"fmt"
	"strconv"
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

type bag struct {
	colour string
	num    int
}

func parseRules(rules []string) map[string][]bag {
	ruleMap := make(map[string][]bag)

	for _, rule := range rules {
		words := strings.Split(rule, " ")
		container := words[0] + words[1]
		internalBag := []bag{}

		if words[4] == "no" {
			bag := bag{
				colour: "",
				num:    0,
			}
			internalBag = append(internalBag, bag)
			ruleMap[container] = internalBag
		} else {
			// if there are bags
			for i := 4; i < len(words); i += 4 {

				total, err := strconv.Atoi(words[i])
				panicer.Check(err)

				bag := bag{
					colour: words[i+1] + words[i+2],
					num:    total,
				}
				internalBag = append(internalBag, bag)
			}
			ruleMap[container] = internalBag
		}
	}

	return ruleMap
}

func countContains(colour string, ruleMap map[string][]bag) {

	containsColour := make(map[string]string)
	containsColour[colour] = "init"
	count := 0

	finished := false
	for !finished {

		for container, bags := range ruleMap {
			for _, bag := range bags {
				if containsColour[bag.colour] != "" {
					containsColour[container] = "yes"
				}
			}
		}

		//exit condition
		currentCount := 0
		for _, val := range containsColour {
			if val == "yes" {
				currentCount++
			}
		}

		if currentCount == count {
			finished = true
		}
		count = currentCount
	}

	fmt.Println("P1 count of bags that can contain ", colour, " is ", count)
}

// the idea is to follow down the tree of all bags we must internally contain
// this is recursive, but we can use memoization to improve performance
func internallyContains(colour string, ruleMap map[string][]bag, memo map[string]int) int {

	// base case
	internalBags := ruleMap[colour]
	if len(internalBags) == 1 {
		if internalBags[0].num == 0 {
			return 0
		}
	}

	// recursive case
	count := 0
	for _, internalBag := range internalBags {
		for i := 0; i < internalBag.num; i++ {

			if memo[internalBag.colour] > 0 {
				count += 1 + memo[internalBag.colour]
			} else {
				bagContains := internallyContains(internalBag.colour, ruleMap, memo)
				memo[internalBag.colour] = bagContains
				count += 1 + bagContains
			}

		}
	}

	return count
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)

	panicer.Check(err)

	ruleMap := parseRules(lines)

	countContains("shinygold", ruleMap)

	memo := make(map[string]int)
	count := internallyContains("shinygold", ruleMap, memo)

	fmt.Println("P2 count of bags contained is ", count)

}
