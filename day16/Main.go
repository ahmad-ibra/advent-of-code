package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

type rule struct {
	name string

	range1min int
	range1max int

	range2min int
	range2max int
}

func (r *rule) isValid(num int) bool {
	if num >= r.range1min && num <= r.range1max {
		return true
	}
	if num >= r.range2min && num <= r.range2max {
		return true
	}
	return false
}

func generateRules(rules []string) []rule {
	var ruleArray []rule

	for _, thisRule := range rules {
		var newRule rule

		splitRule := strings.Split(thisRule, ": ")

		newRule.name = splitRule[0]
		splitRuleAgain := strings.Split(splitRule[1], " or ")

		range1 := strings.Split(splitRuleAgain[0], "-")
		range2 := strings.Split(splitRuleAgain[1], "-")

		num, _ := strconv.Atoi(range1[0])
		newRule.range1min = num
		num, _ = strconv.Atoi(range1[1])
		newRule.range1max = num
		num, _ = strconv.Atoi(range2[0])
		newRule.range2min = num
		num, _ = strconv.Atoi(range2[1])
		newRule.range2max = num

		ruleArray = append(ruleArray, newRule)
	}

	return ruleArray
}

func checkValidity(tickets []string, ruleArray []rule) {

	invalidValues := 0

	for ticketNum := 1; ticketNum < len(tickets); ticketNum++ {

		values := strings.Split(tickets[ticketNum], ",")
		for _, value := range values {
			val, _ := strconv.Atoi(value)
			// for each rule
			valid := false
			for ruleNum := 0; ruleNum < len(ruleArray); ruleNum++ {

				thisRule := ruleArray[ruleNum]

				// if not valid
				if thisRule.isValid(val) {
					valid = true
					break
				}
			}
			if !valid {
				invalidValues += val
			}
		}
	}
	fmt.Println("P1: ", invalidValues)
}

func getValid(tickets []string, ruleArray []rule) []string {

	var validTickets []string

	for ticketNum := 1; ticketNum < len(tickets); ticketNum++ {
		thisTicket := tickets[ticketNum]
		values := strings.Split(thisTicket, ",")
		validTicket := true

		for _, value := range values {
			val, _ := strconv.Atoi(value)

			// for each rule
			validVal := false
			for ruleNum := 0; ruleNum < len(ruleArray); ruleNum++ {
				thisRule := ruleArray[ruleNum]

				// if not valid
				if thisRule.isValid(val) {
					validVal = true
					break
				}
			}
			if !validVal {
				validTicket = false
			}
		}
		if validTicket {
			validTickets = append(validTickets, thisTicket)
		}
	}
	return validTickets
}

/*
 * Returns a map of columns to potential rules they're associated with
 * Key is column of the ticket (0,1,2,...)
 * Value is an internal map of string to int
 *			Internal key is the rule name
 *			Internal value is bool of if the column could be that rule
 */
func initColToRuleMap(ticket string, ruleArray []rule) map[int]map[string]bool {

	colToRuleMap := make(map[int]map[string]bool)
	values := strings.Split(ticket, ",")
	for col := 0; col < len(values); col++ {

		// generate internal map of rule -> bool
		internalMap := make(map[string]bool)

		for _, rule := range ruleArray {
			internalMap[rule.name] = true
		}
		// initialize external map
		colToRuleMap[col] = internalMap
	}
	return colToRuleMap
}

/*
 * Loops over the tickets removing invalid col possibilities from colToRuleMap
 * Removal from the colToRuleMap map is done by setting its internal map value to false
 */
func removeInvalidOptions(tickets []string, ruleArray []rule, colToRuleMap map[int]map[string]bool) map[int]map[string]bool {
	for ticketNum := 0; ticketNum < len(tickets); ticketNum++ {
		thisTicket := tickets[ticketNum]
		values := strings.Split(thisTicket, ",")

		for col, value := range values {
			val, _ := strconv.Atoi(value)

			// for each rule
			for ruleNum := 0; ruleNum < len(ruleArray); ruleNum++ {

				thisRule := ruleArray[ruleNum]

				// if not valid
				if !thisRule.isValid(val) {
					colToRuleMap[col][thisRule.name] = false
				}
			}
		}
	}
	return colToRuleMap
}

/*
 *  Go over each columns internal values.
 *  If any of them only have 1 rule, they must be the column for that rule. Disable the rule in other columns.
 */
func determineFinalOptions(tickets []string, ruleArray []rule, colToRuleMap map[int]map[string]bool) map[int]map[string]bool {

	column := 0
	key := ""
	uniqueRule := false
	noMoreUnique := false

	updatedRules := make(map[string]bool)
	for {
		// clean it out from other entries
		if uniqueRule {
			updatedRules[key] = true
			for i := 0; i < len(colToRuleMap); i++ {
				if i != column {
					colToRuleMap[i][key] = false
				}
			}
		} else {
			if noMoreUnique {
				break
			}
		}
		// reset values and search again
		uniqueRule = false
		column = 0
		key = ""

		for col, rule := range colToRuleMap {

			ruleCount := 0

			for k, v := range rule {
				if v {
					ruleCount++
					column = col
					key = k
				}
			}

			// unique rule
			if ruleCount == 1 && !updatedRules[key] {
				uniqueRule = true
				break
			} else {
				noMoreUnique = true
			}
		}
	}

	return colToRuleMap
}

/*
 * Returns the product of columns matching substring in ticket
 */
func findProductOfColumns(colToRuleMap map[int]map[string]bool, ticket string, substring string) int {
	myTicketArray := strings.Split(ticket, ",")
	output := 1
	// find columns that are "departure"
	for col, rule := range colToRuleMap {
		for k, v := range rule {
			if strings.Contains(k, substring) && v {
				val, _ := strconv.Atoi(myTicketArray[col])
				output *= val
			}
		}
	}
	return output
}

func findDepartureCounts(tickets []string, ruleArray []rule, myTicket []string) {

	colToRuleMap := initColToRuleMap(tickets[0], ruleArray)
	colToRuleMap = removeInvalidOptions(tickets, ruleArray, colToRuleMap)
	colToRuleMap = determineFinalOptions(tickets, ruleArray, colToRuleMap)

	fmt.Println("P2: ", findProductOfColumns(colToRuleMap, myTicket[1], "departure"))
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	groups := input.GroupLines(lines)
	ruleArray := generateRules(groups[0])

	// part 1
	checkValidity(groups[2], ruleArray)

	// part 2
	validTickets := getValid(groups[2], ruleArray)
	findDepartureCounts(validTickets, ruleArray, groups[1])
}
