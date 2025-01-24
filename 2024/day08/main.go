package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	SKIP = "."
)

type Position struct {
	row int
	col int
}

type CityMap struct {
	maxRows  int
	maxCols  int
	antennas map[string][]Position // map of antenna type to list of positions
}

func (cm CityMap) CountAntinodes() int {
	antinodes := make(map[Position]struct{})
	for _, antennaList := range cm.antennas {
		for i := 0; i < len(antennaList); i++ {
			for j := i + 1; j < len(antennaList); j++ {
				antinodePositions := getAntinodePositions(antennaList[i], antennaList[j])
				for _, antinode := range antinodePositions {
					if antinode.row < 0 || antinode.row >= cm.maxRows || antinode.col < 0 || antinode.col >= cm.maxCols {
						continue
					}
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func getAntinodePositions(p1 Position, p2 Position) []Position {
	antinodePositions := []Position{}

	rDiff := p2.row - p1.row
	cDiff := p2.col - p1.col

	antinodePositions = append(antinodePositions, Position{p1.row - rDiff, p1.col - cDiff})
	antinodePositions = append(antinodePositions, Position{p2.row + rDiff, p2.col + cDiff})

	return antinodePositions
}

func (cm CityMap) CountAntinodesResonant() int {
	antinodes := make(map[Position]struct{})
	for _, antennaList := range cm.antennas {
		for i := 0; i < len(antennaList); i++ {
			for j := i + 1; j < len(antennaList); j++ {
				antinodePositions := getAntinodePositionsResonant(antennaList[i], antennaList[j], cm.maxRows, cm.maxCols)
				for _, antinode := range antinodePositions {
					antinodes[antinode] = struct{}{}
				}
			}
		}
	}
	return len(antinodes)
}

func getAntinodePositionsResonant(p1 Position, p2 Position, maxRows, maxCols int) []Position {
	antinodePositions := []Position{}

	rDiff := p2.row - p1.row
	cDiff := p2.col - p1.col

	antinodePositions = append(antinodePositions, p1)
	// check before p1
	for {
		cur := antinodePositions[len(antinodePositions)-1]
		p := Position{cur.row - rDiff, cur.col - cDiff}
		if p.row < 0 || p.row >= maxRows || p.col < 0 || p.col >= maxCols {
			break
		}
		antinodePositions = append(antinodePositions, p)
	}

	antinodePositions = append(antinodePositions, p2)
	// check after p2
	for {
		cur := antinodePositions[len(antinodePositions)-1]
		p := Position{cur.row + rDiff, cur.col + cDiff}
		if p.row < 0 || p.row >= maxRows || p.col < 0 || p.col >= maxCols {
			break
		}
		antinodePositions = append(antinodePositions, p)
	}

	return antinodePositions
}

func parseInput(str string) CityMap {
	bytes, _ := os.ReadFile(str)
	cm := CityMap{
		antennas: make(map[string][]Position),
		maxRows:  len(strings.Split(string(bytes), "\n")) - 1,
		maxCols:  len(strings.Split(string(bytes), "\n")[0]),
	}

	for r, line := range strings.Split(string(bytes), "\n") {
		if line == "" {
			continue
		}
		for c, cell := range strings.Split(line, "") {
			if cell == SKIP {
				continue
			}

			if _, ok := cm.antennas[cell]; !ok {
				cm.antennas[cell] = []Position{}
			}
			cm.antennas[cell] = append(cm.antennas[cell], Position{r, c})
		}
	}

	return cm
}

func main() {

	cityMap := parseInput("input.txt")
	//cityMap := parseInput("sample.txt")

	fmt.Println("Part1:", cityMap.CountAntinodes())
	fmt.Println("Part2:", cityMap.CountAntinodesResonant())

}
