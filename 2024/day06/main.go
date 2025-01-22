package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	EMPTY    = "."
	GUARD    = "^"
	OBSTACLE = "#"
)

type Position struct {
	row int
	col int
}

func (p Position) toString() string {
	return "(" + strconv.Itoa(p.row) + "," + strconv.Itoa(p.col) + ")"
}

type GuardPath struct {
	guard            Position
	guardDirection   string
	maxRow           int
	maxCol           int
	obstacles        map[string]struct{}
	visited          map[string]struct{}
	visitedDirection map[string]struct{}
}

func (gp *GuardPath) DeepCopy() GuardPath {
	copy := GuardPath{
		guard:            gp.guard,
		guardDirection:   gp.guardDirection,
		maxRow:           gp.maxRow,
		maxCol:           gp.maxCol,
		obstacles:        make(map[string]struct{}),
		visited:          make(map[string]struct{}),
		visitedDirection: make(map[string]struct{}),
	}

	for key, value := range gp.obstacles {
		copy.obstacles[key] = value
	}
	for key, value := range gp.visited {
		copy.visited[key] = value
	}
	for key, value := range gp.visitedDirection {
		copy.visitedDirection[key] = value
	}

	return copy
}

func (gp *GuardPath) inFront(o Position) Position {
	switch gp.guardDirection {
	case "N":
		return Position{o.row - 1, o.col}
	case "S":
		return Position{o.row + 1, o.col}
	case "E":
		return Position{o.row, o.col + 1}
	case "W":
		return Position{o.row, o.col - 1}
	}
	return Position{}
}

func (gp *GuardPath) turnRight() {
	switch gp.guardDirection {
	case "N":
		gp.guardDirection = "E"
	case "E":
		gp.guardDirection = "S"
	case "S":
		gp.guardDirection = "W"
	case "W":
		gp.guardDirection = "N"
	}
}

// returns true if the guard can move to the next position
func (gp *GuardPath) nextPosition() bool {
	// set current position as visited
	gp.visited[gp.guard.toString()] = struct{}{}
	gp.visitedDirection[gp.guard.toString()+gp.guardDirection] = struct{}{}

	if _, ok := gp.obstacles[gp.inFront(gp.guard).toString()]; ok {
		// if obstacle in front, turn right
		gp.turnRight()
	} else {
		// otherwise move forward
		gp.guard = gp.inFront(gp.guard)
	}

	// if guard has left the map, return false
	if gp.guard.row < 0 || gp.guard.row >= gp.maxRow || gp.guard.col < 0 || gp.guard.col >= gp.maxCol {
		return false
	}

	return true
}

func (gp *GuardPath) hasLoop() bool {
	for gp.nextPosition() {
		if _, ok := gp.visitedDirection[gp.guard.toString()+gp.guardDirection]; ok {
			return true
		}
	}
	return false
}

func parseInput(filename string) GuardPath {
	byte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	gp := GuardPath{
		maxRow:           len(strings.Split(string(byte), "\n")),
		maxCol:           len(strings.Split(string(byte), "\n")[0]),
		visited:          make(map[string]struct{}),
		obstacles:        make(map[string]struct{}),
		visitedDirection: make(map[string]struct{}),
		guardDirection:   "N",
	}

	for r, line := range strings.Split(string(byte), "\n") {
		if len(line) == 0 {
			gp.maxRow--
			continue
		}
		for c, char := range line {
			if string(char) == GUARD {
				gp.guard = Position{r, c}
			}
			if string(char) == OBSTACLE {
				gp.obstacles[Position{r, c}.toString()] = struct{}{}
			}
		}
	}

	return gp
}

func main() {

	gp := parseInput("input.txt")
	//gp := parseInput("sample.txt")

	gp1 := gp.DeepCopy()
	for gp1.nextPosition() {
	}

	fmt.Println("Part1:", len(gp1.visited))

	countLoopObstacles := 0
	for r := 0; r < gp.maxRow; r++ {
		for c := 0; c < gp.maxCol; c++ {
			// try replacing every non guard and non obstacle position with an obstacle, see if it leads to a loop
			gp2 := gp.DeepCopy()
			if gp2.guard.row == r && gp2.guard.col == c {
				continue
			}
			if _, ok := gp2.obstacles[Position{r, c}.toString()]; ok {
				continue
			}

			// adding new obstacle at chosen location
			gp2.obstacles[Position{r, c}.toString()] = struct{}{}

			if gp2.hasLoop() {
				countLoopObstacles++
				delete(gp2.obstacles, Position{r, c}.toString())
				continue
			}
			delete(gp2.obstacles, Position{r, c}.toString())
		}
	}

	fmt.Println("Part2:", countLoopObstacles)
}
