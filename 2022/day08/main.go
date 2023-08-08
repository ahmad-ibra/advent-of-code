package main

import (
	in "aoc2022/input"
	"fmt"
	"strconv"
)

func isVisible(height, row, col int, grid [][]int) bool {
	if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[0])-1 {
		return true
	}

	visibleLeft := true
	for c := 0; c < col; c++ {
		if height <= grid[row][c] {
			visibleLeft = false
			break
		}
	}

	visibleRight := true
	for c := col + 1; c < len(grid[0]); c++ {
		if height <= grid[row][c] {
			visibleRight = false
			break
		}
	}

	visibleUp := true
	for r := 0; r < row; r++ {
		if height <= grid[r][col] {
			visibleUp = false
			break
		}
	}
	visibleDown := true
	for r := row + 1; r < len(grid); r++ {
		if height <= grid[r][col] {
			visibleDown = false
			break
		}
	}

	return visibleLeft || visibleRight || visibleUp || visibleDown
}

func getScenicScore(height, row, col int, grid [][]int) int {
	if row == 0 || col == 0 || row == len(grid)-1 || col == len(grid[0])-1 {
		return 0
	}

	left := 0
	for c := col - 1; c >= 0; c-- {
		left++
		if height <= grid[row][c] || c == 0 {
			break
		}
	}
	right := 0
	for c := col + 1; c < len(grid[0]); c++ {
		right++
		if height <= grid[row][c] || c == len(grid[0])-1 {
			break
		}
	}
	up := 0
	for r := row - 1; r >= 0; r-- {
		up++
		if height <= grid[r][col] || r == 0 {
			break
		}
	}
	down := 0
	for r := row + 1; r < len(grid); r++ {
		down++
		if height <= grid[r][col] || r == len(grid)-1 {
			break
		}
	}

	return left * right * up * down
}

func solution(lines []string) {

	rows := len(lines)
	cols := len(lines[0])

	grid := make([][]int, len(lines))
	for i := 0; i < cols; i++ {
		grid[i] = make([]int, rows)
	}

	for r := 0; r < rows; r++ {
		line := lines[r]
		if line == "" {
			continue
		}

		for c := 0; c < cols; c++ {
			heightStr := line[c]
			height, err := strconv.Atoi(string(heightStr))
			if err != nil {
				panic(err)
			}
			grid[r][c] = height
		}
	}

	visibleCount := 0
	scenicScore := 0
	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[0]); c++ {

			score := getScenicScore(grid[r][c], r, c, grid)
			if score > scenicScore {
				scenicScore = score
			}

			if isVisible(grid[r][c], r, c, grid) {
				visibleCount++
				continue
			}
		}
	}

	fmt.Printf("%v trees are visible\n", visibleCount)
	fmt.Printf("best scenic score is %v\n", scenicScore)
}

func main() {
	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
}
