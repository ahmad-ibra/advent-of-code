package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func pathToFile(fName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	check(err)

	return filepath.Join(dir, fName)
}

func countTrees(treeMap []string, rightDelta int, downDelta int) int {

	widthOfMap := len(treeMap[0])
	xPos := 0
	treeCount := 0
	tree := "#"

	downIncrement := 0
	for _, line := range treeMap {
		// if we've moved down enough
		if downIncrement == downDelta {

			xPos += rightDelta
			moddedXPos := xPos % widthOfMap

			currentSpot := string(line[moddedXPos])

			if currentSpot == tree {
				treeCount++
			}

			downIncrement = 0
		}
		downIncrement++
	}
	return treeCount
}

func main() {

	path := pathToFile("input.txt")

	lines, err := readLines(path)
	check(err)

	fmt.Println("Part 1: ", countTrees(lines, 3, 1))

	fmt.Println("Part 2: ", countTrees(lines, 1, 1)*countTrees(lines, 3, 1)*countTrees(lines, 5, 1)*countTrees(lines, 7, 1)*countTrees(lines, 1, 2))

}
