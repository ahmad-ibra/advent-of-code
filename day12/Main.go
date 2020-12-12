package main

import (
	"fmt"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func part1(lines []string) {

	var p position
	p.dir = east

	for _, line := range lines {
		move(&p, line)
	}

	fmt.Println("FinalPos| NS: ", p.NS, ", EW: ", p.EW, ", dir: ", p.dir)
	fmt.Println("P1: The Manhatten Distance is: ", p.manhattenDistance())
}

func part2(lines []string) {

	// init a waypoint 10 east and 1 north of ship
	var wp waypoint
	wp.NS = 1
	wp.EW = 10

	// init a ship at the 0 position
	var p position
	p.WP = wp

	for _, line := range lines {
		moveP2(&p, line)
	}

	fmt.Println("FinalPos| NS: ", p.NS, ", EW: ", p.EW, ", dir: ", p.dir)
	fmt.Println("P2: The Manhatten Distance is: ", p.manhattenDistance())
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	part1(lines)
	part2(lines)

}
