package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	X = "X"
	M = "M"
	A = "A"
	S = "S"
)

func setupPuzzle(file string) [][]string {
	bytes, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	puzzle := make([][]string, 0)

	lines := strings.Split(string(bytes), "\n")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		puzzle = append(puzzle, strings.Split(line, ""))
	}

	return puzzle
}

func countXmas(puzzle [][]string) int {
	count := 0

	rows := len(puzzle)
	cols := len(puzzle[0])

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if puzzle[r][c] == X {
				count += checkLeft(puzzle, r, c, rows, cols)
				count += checkRight(puzzle, r, c, rows, cols)
				count += checkUp(puzzle, r, c, rows, cols)
				count += checkDown(puzzle, r, c, rows, cols)
				count += checkUpLeft(puzzle, r, c, rows, cols)
				count += checkUpRight(puzzle, r, c, rows, cols)
				count += checkDownRight(puzzle, r, c, rows, cols)
				count += checkDownLeft(puzzle, r, c, rows, cols)
			}
		}
	}

	return count
}

func checkLeft(puzzle [][]string, r, c, rows, cols int) int {
	if c < 3 {
		return 0
	}
	if puzzle[r][c-1] == M && puzzle[r][c-2] == A && puzzle[r][c-3] == S {
		return 1
	}
	return 0
}

func checkRight(puzzle [][]string, r, c, rows, cols int) int {
	if c > cols-4 {
		return 0
	}
	if puzzle[r][c+1] == M && puzzle[r][c+2] == A && puzzle[r][c+3] == S {
		return 1
	}
	return 0
}

func checkUp(puzzle [][]string, r, c, rows, cols int) int {
	if r < 3 {
		return 0
	}
	if puzzle[r-1][c] == M && puzzle[r-2][c] == A && puzzle[r-3][c] == S {
		return 1
	}
	return 0
}

func checkDown(puzzle [][]string, r, c, rows, cols int) int {
	if r > rows-4 {
		return 0
	}
	if puzzle[r+1][c] == M && puzzle[r+2][c] == A && puzzle[r+3][c] == S {
		return 1
	}
	return 0
}

func checkUpLeft(puzzle [][]string, r, c, rows, cols int) int {
	if r < 3 || c < 3 {
		return 0
	}
	if puzzle[r-1][c-1] == M && puzzle[r-2][c-2] == A && puzzle[r-3][c-3] == S {
		return 1
	}
	return 0
}

func checkUpRight(puzzle [][]string, r, c, rows, cols int) int {
	if r < 3 || c > cols-4 {
		return 0
	}
	if puzzle[r-1][c+1] == M && puzzle[r-2][c+2] == A && puzzle[r-3][c+3] == S {
		return 1
	}
	return 0
}

func checkDownRight(puzzle [][]string, r, c, rows, cols int) int {
	if r > rows-4 || c > cols-4 {
		return 0
	}
	if puzzle[r+1][c+1] == M && puzzle[r+2][c+2] == A && puzzle[r+3][c+3] == S {
		return 1
	}
	return 0
}

func checkDownLeft(puzzle [][]string, r, c, rows, cols int) int {
	if r > rows-4 || c < 3 {
		return 0
	}
	if puzzle[r+1][c-1] == M && puzzle[r+2][c-2] == A && puzzle[r+3][c-3] == S {
		return 1
	}
	return 0
}

func diagDown(puzzle [][]string, r, c int) bool {
	// top left to bottom right
	if (puzzle[r-1][c-1] == M && puzzle[r+1][c+1] == S) || (puzzle[r-1][c-1] == S && puzzle[r+1][c+1] == M) {
		return true
	}

	return false
}

func diagUp(puzzle [][]string, r, c int) bool {
	// bottom left to top right
	if (puzzle[r+1][c-1] == M && puzzle[r-1][c+1] == S) || (puzzle[r+1][c-1] == S && puzzle[r-1][c+1] == M) {
		return true
	}

	return false
}

func countCrossMas(puzzle [][]string) int {
	count := 0

	rows := len(puzzle)
	cols := len(puzzle[0])

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if puzzle[r][c] == A {
				if diagDown(puzzle, r, c) && diagUp(puzzle, r, c) {
					count++
				}
			}
		}
	}

	return count
}

func main() {
	puzzle := setupPuzzle("input.txt")

	xmasCount := countXmas(puzzle)
	fmt.Println("Part 1:", xmasCount)

	crossMasCount := countCrossMas(puzzle)
	fmt.Println("Part 2:", crossMasCount)
}
