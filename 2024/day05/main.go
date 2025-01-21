package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	Before string
	After  string
}

type RuleSet struct {
	Rules []Rule
}

func (r *RuleSet) AddRule(rule Rule) {
	r.Rules = append(r.Rules, rule)
}

func (r *RuleSet) GetBefore(page string) []string {
	before := make([]string, 0)

	for _, rule := range r.Rules {
		if rule.After == page {
			before = append(before, rule.Before)
		}
	}

	return before
}

func (r *RuleSet) GetAfter(page string) []string {
	after := make([]string, 0)

	for _, rule := range r.Rules {
		if rule.Before == page {
			after = append(after, rule.After)
		}
	}

	return after
}

func (r *RuleSet) IsValid(input PageOrder) bool {
	visited := make(map[string]struct{})

	for _, curPage := range input {
		after := r.GetAfter(curPage)
		for _, a := range after {
			if _, ok := visited[a]; ok {
				return false
			}
		}

		visited[curPage] = struct{}{}
	}

	return true
}

func (r *RuleSet) FixOrder(input PageOrder) PageOrder {
	fixed := make([]string, 0)

	// find first page, page whhere no others in input are before it, add it to fixed list
	for len(input) > 0 {
		for i, curPage := range input {
			before := r.GetBefore(curPage)
			if len(before) == 0 {
				fixed = append(fixed, curPage)

				// remove page from input
				input = append(input[:i], input[i+1:]...)
				break
			}

			// if no pages in input are in the befor list
			if !input.ContainsAny(before) {
				fixed = append(fixed, curPage)

				// remove page from input
				input = append(input[:i], input[i+1:]...)
				break
			}
		}
	}

	return fixed
}

type PageOrder []string

func (p PageOrder) MiddlePage() string {
	return p[len(p)/2]
}

func (p PageOrder) Contains(input string) bool {
	for _, i := range p {
		if i == input {
			return true
		}
	}

	return false
}

func (p PageOrder) ContainsAny(input []string) bool {
	for _, i := range input {
		if p.Contains(i) {
			return true
		}
	}

	return false

}

func parseInput(input string) (RuleSet, []PageOrder) {
	rs := RuleSet{Rules: []Rule{}}
	po := make([]PageOrder, 0)

	bytes, err := os.ReadFile(input)
	if err != nil {
		panic(err)
	}

	for _, line := range strings.Split(string(bytes), "\n") {
		if strings.Contains(line, "|") {
			parts := strings.Split(line, "|")
			rs.AddRule(Rule{Before: parts[0], After: parts[1]})
		}

		if strings.Contains(line, ",") {
			parts := strings.Split(line, ",")
			po = append(po, PageOrder(parts))
		}
	}

	return rs, po
}

func main() {
	ruleSet, pageOrders := parseInput("input.txt")

	sumP1 := 0
	sumP2 := 0
	for _, po := range pageOrders {
		if ruleSet.IsValid(po) {
			num, err := strconv.Atoi(po.MiddlePage())
			if err != nil {
				panic(err)
			}

			sumP1 += num
		} else {
			fixedPo := ruleSet.FixOrder(po)
			num, err := strconv.Atoi(fixedPo.MiddlePage())
			if err != nil {
				panic(err)
			}
			sumP2 += num
		}
	}

	fmt.Println("Part 1:", sumP1)
	fmt.Println("Part 2:", sumP2)
}
