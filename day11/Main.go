package main

import (
	"fmt"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func buildMap(lines []string) [][]rune {
	var mapOfLines [][]rune

	for _, line := range lines {
		characters := []rune(line)

		mapOfLines = append(mapOfLines, characters)

	}
	return mapOfLines
}

func countOccupiedP1(y int, x int, seatMap [][]rune) int {
	occupied := 0

	maxX := len(seatMap[0]) - 1
	maxY := len(seatMap) - 1

	// left top
	if x != 0 && y != 0 {
		leftTop := string(seatMap[y-1][x-1])
		if leftTop == "#" {
			occupied++
		}
	}
	// middle top
	if y != 0 {
		middleTop := string(seatMap[y-1][x])
		if middleTop == "#" {
			occupied++
		}
	}
	// right top
	if x != maxX && y != 0 {
		rightTop := string(seatMap[y-1][x+1])
		if rightTop == "#" {
			occupied++
		}
	}
	// left middle
	if x != 0 {
		leftMiddle := string(seatMap[y][x-1])
		if leftMiddle == "#" {
			occupied++
		}
	}

	// right middle
	if x != maxX {
		rightMiddle := string(seatMap[y][x+1])
		if rightMiddle == "#" {
			occupied++
		}
	}

	// left bottom
	if x != 0 && y != maxY {
		leftBottom := string(seatMap[y+1][x-1])
		if leftBottom == "#" {
			occupied++
		}
	}

	// middle bottom
	if y != maxY {
		middleBottom := string(seatMap[y+1][x])
		if middleBottom == "#" {
			occupied++
		}
	}

	// right bottom
	if x != maxX && y != maxY {
		rightBottom := string(seatMap[y+1][x+1])
		if rightBottom == "#" {
			occupied++
		}
	}
	return occupied
}

func timeLoopP1(seatMap [][]rune) {
	// loop on each round
	i := 1

	for {
		occupied := 0
		hasChanged := false

		nextMap := [][]rune{}
		for y := 0; y < len(seatMap); y++ {
			curRunes := []rune{}
			for x := 0; x < len(seatMap[y]); x++ {
				curRunes = append(curRunes, seatMap[y][x])
			}
			nextMap = append(nextMap, curRunes)
		}
		occupiedSeat := []rune("#")
		seat := []rune("L")

		// loop on each row
		for y := 0; y < len(seatMap); y++ {

			// loop on each column
			for x := 0; x < len(seatMap[y]); x++ {

				curChar := string(seatMap[y][x])

				// if seat is empty
				if curChar == "L" {
					// if all adjacent seats are empty, occupy the seat
					if countOccupiedP1(y, x, seatMap) == 0 {
						nextMap[y][x] = occupiedSeat[0]
						hasChanged = true
						occupied++
					}
				}
				// if seat is occupied
				if curChar == "#" {
					// if 4 or more adjacent seats are occupied, empty the seat
					if countOccupiedP1(y, x, seatMap) >= 4 {
						nextMap[y][x] = seat[0]
						hasChanged = true
					} else {
						occupied++
					}
				}
			}
		}

		// set seatMap to nextMap
		for y := 0; y < len(seatMap); y++ {
			for x := 0; x < len(seatMap[y]); x++ {
				seatMap[y][x] = nextMap[y][x]
			}
		}

		seatMap = nextMap
		i++

		if !hasChanged {
			fmt.Println("P1 occupied seats: ", occupied)
			break
		}
	}
}

func countOccupiedP2(y int, x int, seatMap [][]rune) int {
	occupied := 0

	maxX := len(seatMap[0]) - 1
	maxY := len(seatMap) - 1

	// left top
	iterX := x
	iterY := y
	for {
		if iterX != 0 && iterY != 0 {
			iterX--
			iterY--
			leftTop := string(seatMap[iterY][iterX])
			if leftTop == "#" {
				occupied++
				break
			}
			if leftTop == "L" {
				break
			}
		} else {
			break
		}
	}
	// middle top
	iterX = x
	iterY = y
	for {
		if iterY != 0 {
			iterY--
			middleTop := string(seatMap[iterY][iterX])
			if middleTop == "#" {
				occupied++
				break
			}
			if middleTop == "L" {
				break
			}
		} else {
			break
		}
	}
	// right top
	iterX = x
	iterY = y
	for {
		if iterX != maxX && iterY != 0 {
			iterY--
			iterX++
			rightTop := string(seatMap[iterY][iterX])
			if rightTop == "#" {
				occupied++
				break
			}
			if rightTop == "L" {
				break
			}
		} else {
			break
		}
	}
	// left middle
	iterX = x
	iterY = y
	for {
		if iterX != 0 {
			iterX--
			leftMiddle := string(seatMap[iterY][iterX])
			if leftMiddle == "#" {
				occupied++
				break
			}
			if leftMiddle == "L" {
				break
			}
		} else {
			break
		}
	}
	// right middle
	iterX = x
	iterY = y
	for {
		if iterX != maxX {
			iterX++
			rightMiddle := string(seatMap[iterY][iterX])
			if rightMiddle == "#" {
				occupied++
				break
			}
			if rightMiddle == "L" {
				break
			}
		} else {
			break
		}
	}
	// left bottom
	iterX = x
	iterY = y
	for {
		if iterX != 0 && iterY != maxY {
			iterX--
			iterY++
			leftBottom := string(seatMap[iterY][iterX])
			if leftBottom == "#" {
				occupied++
				break
			}
			if leftBottom == "L" {
				break
			}
		} else {
			break
		}
	}
	// middle bottom
	iterX = x
	iterY = y
	for {
		if iterY != maxY {
			iterY++
			middleBottom := string(seatMap[iterY][iterX])
			if middleBottom == "#" {
				occupied++
				break
			}
			if middleBottom == "L" {
				break
			}
		} else {
			break
		}
	}
	// right bottom
	iterX = x
	iterY = y
	for {
		if iterX != maxX && iterY != maxY {
			iterY++
			iterX++
			rightBottom := string(seatMap[iterY][iterX])
			if rightBottom == "#" {
				occupied++
				break
			}
			if rightBottom == "L" {
				break
			}
		} else {
			break
		}
	}
	return occupied
}

func timeLoopP2(seatMap [][]rune) {
	// loop on each round
	for {
		occupied := 0
		hasChanged := false

		nextMap := [][]rune{}
		for y := 0; y < len(seatMap); y++ {
			curRunes := []rune{}
			for x := 0; x < len(seatMap[y]); x++ {
				curRunes = append(curRunes, seatMap[y][x])
			}
			nextMap = append(nextMap, curRunes)
		}
		occupiedSeat := []rune("#")
		seat := []rune("L")

		// loop on each row
		for y := 0; y < len(seatMap); y++ {

			// loop on each column
			for x := 0; x < len(seatMap[y]); x++ {

				curChar := string(seatMap[y][x])

				// if seat is empty
				if curChar == "L" {
					// if all adjacent seats are empty, occupy the seat
					if countOccupiedP2(y, x, seatMap) == 0 {
						nextMap[y][x] = occupiedSeat[0]
						hasChanged = true
						occupied++
					}
				}
				// if seat is occupied
				if curChar == "#" {
					// if 4 or more adjacent seats are occupied, empty the seat
					if countOccupiedP2(y, x, seatMap) >= 5 {
						nextMap[y][x] = seat[0]
						hasChanged = true
					} else {
						occupied++
					}
				}
			}
		}

		// set seatMap to nextMap
		for y := 0; y < len(seatMap); y++ {
			for x := 0; x < len(seatMap[y]); x++ {
				seatMap[y][x] = nextMap[y][x]
			}
		}

		seatMap = nextMap
		if !hasChanged {
			fmt.Println("P2 occupied seats: ", occupied)
			break
		}
	}
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	seatMap := buildMap(lines)

	timeLoopP1(seatMap)
	timeLoopP2(seatMap)
}
