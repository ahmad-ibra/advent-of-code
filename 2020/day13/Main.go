package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func part1(notes []string) {
	line1 := notes[0]
	startTime, _ := strconv.Atoi(line1)

	busses := strings.Split(notes[1], ",")

	chosenBus := 0
	earliestDeparture := 999999999

	for _, bus := range busses {
		if bus != "x" {
			departureTime := startTime
			busNum, _ := strconv.Atoi(bus)
			for {
				depart := departureTime % busNum
				if departureTime > earliestDeparture {
					break
				}
				if depart == 0 {
					earliestDeparture = departureTime
					chosenBus = busNum
					break
				}
				departureTime++
			}
		}
	}

	waitTime := earliestDeparture - startTime

	fmt.Println("chosenBus: ", chosenBus, ", earliestDeparture: ", earliestDeparture, ", waitTime: ", waitTime)
	fmt.Println("Part1: ", chosenBus*waitTime)
}

/*
 * Attempt 1 at part 2, a brute force approach
 * Taking too long on the real input data
 */
func part2a(notes []string) {

	busses := strings.Split(notes[1], ",")

	firstBus, _ := strconv.Atoi(busses[0])
	time := firstBus

	finished := false

	// going forward through time by firstBus
	for {
		tempTime := time

		// check remaining busses
		for i := 1; i < len(busses); i++ {
			tempTime++
			nextBusString := busses[i]
			if nextBusString != "x" {
				nextBus, _ := strconv.Atoi(nextBusString)
				if tempTime%nextBus != 0 {
					break
				}
				// we've gone through all busses
				if i == len(busses)-1 {
					finished = true
					fmt.Println("P2: ", time)
					break
				}
			}
		}
		if finished {
			break
		}

		time += firstBus
	}
}

/*
 * Attempt 2 at part 2, a "optimized" brute force approach
 * Still taking too long on the real input data
 * Clearly I need to approach the problem in a different way
 */
func part2b(notes []string) {

	busses := strings.Split(notes[1], ",")

	// find largest number in the array and its offset from first item
	largest := 0
	offset := 0
	for i := 0; i < len(busses); i++ {
		nextBusString := busses[i]
		if nextBusString != "x" {
			nextBus, _ := strconv.Atoi(nextBusString)
			if nextBus > largest {
				largest = nextBus
				offset = i
			}
		}
	}

	// create map of busses with their offset from largest
	busMap := make(map[int]int)
	for i := 0; i < len(busses); i++ {
		nextBusString := busses[i]
		if nextBusString != "x" {
			nextBus, _ := strconv.Atoi(nextBusString)
			busMap[nextBus] = i - offset
		}
	}

	// make a sorted list of buses
	busList := []int{}
	for i := 0; i < len(busses); i++ {
		nextBusString := busses[i]
		if nextBusString != "x" {
			nextBus, _ := strconv.Atoi(nextBusString)
			if nextBus < largest {
				busList = append(busList, nextBus)
			}
		}
	}
	sort.Ints(busList)

	time := largest
	// going forward through time by largest
	for {
		countAll := 0
		finished := false

		for i := len(busList) - 1; i >= 0; i-- {
			bus := busList[i]
			thisOffset := busMap[bus]
			if (time+thisOffset)%bus != 0 {
				break
			}
			countAll++

			if countAll == len(busList) {
				finished = true
				// set the final time
				time = time - offset
				break
			}
		}
		if finished {
			break
		}
		time += largest
	}
	fmt.Println(time)
}

/*
 * Attempt 3 at part 2
 * This is an itterative approach where we incrase the step size
 * as we find busses that match the right condition.
 * This approach is made possible by the fact that all the bus
 * numbers are prime numbers
 */
func part2(notes []string) {

	busses := strings.Split(notes[1], ",")

	busList := []int{}
	for i := 0; i < len(busses); i++ {
		nextBusString := busses[i]
		if nextBusString != "x" {
			nextBus, _ := strconv.Atoi(nextBusString)
			busList = append(busList, nextBus)

		}
	}

	firstBus := busList[0]
	time := firstBus
	stepSize := firstBus

	busMap := make(map[int]int)
	offset := 0
	for i := 1; i < len(busses); i++ {
		offset++
		nextBusString := busses[i]
		if nextBusString != "x" {
			nextBus, _ := strconv.Atoi(nextBusString)
			busMap[nextBus] = offset
		}
	}

	// for each bus
	for i := 1; i < len(busList); i++ {
		bus := busList[i]
		offset := busMap[bus]

		// loop forward through time
		for {
			// we hit the right condition, set the next step size and break
			if (time+offset)%bus == 0 {
				stepSize = stepSize * bus
				break
			}
			time += stepSize
		}
	}
	fmt.Println("Part2: ", time)
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	part1(lines)
	part2(lines)

}
