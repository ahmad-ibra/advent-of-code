package main

import (
	in "aoc2022/input"
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func findPaths(lines []string) {
	S := ""
	for r, line := range lines {
		for c, char := range line {
			if char == 'S' {
				S = toString(r, c)
				break
			}
		}
	}

	visited := make(map[string]int)
	r, c := toPos(S)

	steps := findE(lines, r, c, 0, visited)
	fmt.Printf("Par1: Took %v steps to reach desitnation\n", steps)

	// better solution for part 2 would be BFS from dest to any a or S, but wanted to play with goroutines
	// plus i got to fully re-use part 1's solution
	ch := make(chan int)
	for r, line := range lines {
		for c, char := range line {
			if char == 'S' || char == 'a' {
				wg.Add(1)
				go findEConcurrently(lines, r, c, ch)
			}
		}
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	ans := math.MaxInt
	for val := range ch {
		if val < ans {
			ans = val
		}
	}
	fmt.Printf("Par2: Took %v steps to reach desitnation from any lowest point\n", ans)
}

func findEConcurrently(grid []string, r, c int, ch chan<- int) {
	defer wg.Done()
	visited := make(map[string]int)
	ans := findE(grid, r, c, 0, visited)
	ch <- ans
}

func findE(grid []string, r, c, steps int, visited map[string]int) int {
	if grid[r][c] == 'E' {
		return steps
	}

	visited[toString(r, c)] = steps

	s := math.MaxInt

	// check right
	if c < len(grid[0])-1 && canMove(grid[r][c], grid[r][c+1]) {
		if pastSteps, ok := visited[toString(r, c+1)]; ok {
			if pastSteps > steps+1 {
				st := findE(grid, r, c+1, steps+1, visited)
				if st < s {
					s = st
				}
			}
		} else {
			st := findE(grid, r, c+1, steps+1, visited)
			if st < s {
				s = st
			}
		}
	}
	// check left
	if c > 0 && canMove(grid[r][c], grid[r][c-1]) {
		if pastSteps, ok := visited[toString(r, c-1)]; ok {
			if pastSteps > steps+1 {
				st := findE(grid, r, c-1, steps+1, visited)
				if st < s {
					s = st
				}
			}
		} else {
			st := findE(grid, r, c-1, steps+1, visited)
			if st < s {
				s = st
			}
		}
	}
	// check up
	if r > 0 && canMove(grid[r][c], grid[r-1][c]) {
		if pastSteps, ok := visited[toString(r-1, c)]; ok {
			if pastSteps > steps+1 {
				st := findE(grid, r-1, c, steps+1, visited)
				if st < s {
					s = st
				}
			}
		} else {
			st := findE(grid, r-1, c, steps+1, visited)
			if st < s {
				s = st
			}
		}
	}
	// check down
	if r < len(grid)-1 && canMove(grid[r][c], grid[r+1][c]) {
		if pastSteps, ok := visited[toString(r+1, c)]; ok {
			if pastSteps > steps+1 {
				st := findE(grid, r+1, c, steps+1, visited)
				if st < s {
					s = st
				}
			}
		} else {
			st := findE(grid, r+1, c, steps+1, visited)
			if st < s {
				s = st
			}
		}
	}

	return s
}

func canMove(src, dest byte) bool {
	if dest == 'E' {
		return src == 'z' || src == 'y'
	}

	if src == 'S' {
		return dest == 'a' || dest == 'b'
	}

	s := int(src - 'a')
	d := int(dest - 'a')

	return d-s <= 1
}

func toPos(s string) (int, int) {
	pos := strings.Split(s, "|")
	r, _ := strconv.Atoi(pos[0])
	c, _ := strconv.Atoi(pos[1])
	return r, c
}

func toString(r, c int) string {
	var sb strings.Builder
	sb.WriteString(strconv.Itoa(r))
	sb.WriteString("|")
	sb.WriteString(strconv.Itoa(c))
	return sb.String()
}

func main() {
	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	findPaths(lines)
}
