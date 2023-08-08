package main

import (
	in "aoc2022/input"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type position2 struct {
	pos     []coordinates
	visited map[string]struct{}
}

type position struct {
	hX int
	hY int
	tX int
	tY int

	visited map[string]struct{}
}

type instruction struct {
	dir   string
	count int
}

func (p *position) print() string {
	return fmt.Sprintf("head is at: %v, tail is at %v\n", posToStr(p.hX, p.hY), posToStr(p.tX, p.tY))
}

func (p *position) moveHead(dir string) {
	switch dir {
	case "R":
		p.hX++
	case "L":
		p.hX--
	case "U":
		p.hY++
	case "D":
		p.hY--
	default:
		fmt.Printf("didnt move for dir %v\n", dir)
	}
}

func (p *position2) moveHead(dir string) {
	switch dir {
	case "R":
		p.pos[0].x++
	case "L":
		p.pos[0].x--
	case "U":
		p.pos[0].y++
	case "D":
		p.pos[0].y--
	default:
		fmt.Printf("didnt move for dir %v\n", dir)
	}

	p.moveRest()
}

func (p *position2) moveRest() {
	iCur := 1
	iPrev := 0

	for iCur < 10 {
		cur := p.pos[iCur]
		prev := p.pos[iPrev]

		if math.Abs(float64(prev.x)-float64(cur.x)) > 1 || math.Abs(float64(prev.y)-float64(cur.y)) > 1 {
			if prev.x > cur.x {
				cur.x++
			} else if prev.x < cur.x {
				cur.x--
			}
			if prev.y > cur.y {
				cur.y++
			} else if prev.y < cur.y {
				cur.y--
			}
		}

		p.pos[iCur] = cur
		iCur++
		iPrev++
	}

	p.visited[posToStr(p.pos[9].x, p.pos[9].y)] = struct{}{}
}

func (p *position) moveTail() {
	if math.Abs(float64(p.hX)-float64(p.tX)) <= 1 && math.Abs(float64(p.hY)-float64(p.tY)) <= 1 {
		return
	} else {
		if p.hX > p.tX {
			p.tX++
		} else if p.hX < p.tX {
			p.tX--
		}
		if p.hY > p.tY {
			p.tY++
		} else if p.hY < p.tY {
			p.tY--
		}
	}
	p.visited[posToStr(p.tX, p.tY)] = struct{}{}
}

func (p *position2) move(line string) {
	parts := strings.Split(line, " ")
	count, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	inst := &instruction{dir: parts[0], count: count}
	for i := 0; i < inst.count; i++ {
		p.moveHead(inst.dir)
	}
}

func (p *position) move(line string) {
	parts := strings.Split(line, " ")
	count, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	inst := &instruction{dir: parts[0], count: count}
	for i := 0; i < inst.count; i++ {
		p.moveHead(inst.dir)
		p.moveTail()
	}
}

func posToStr(x, y int) string {
	var sb strings.Builder

	xPos := strconv.Itoa(x)
	yPos := strconv.Itoa(y)

	sb.WriteString(xPos)
	sb.WriteString(",")
	sb.WriteString(yPos)

	test := sb.String()

	return test
}

func solution(lines []string) {
	p := &position{hX: 0, hY: 0, tX: 0, tY: 0, visited: make(map[string]struct{})}
	p.visited[posToStr(p.tX, p.tY)] = struct{}{}

	for _, line := range lines {
		if line != "" {
			p.move(line)
		}
	}

	fmt.Printf("%v positions visited by the tail\n", len(p.visited))
}

func solution2(lines []string) {
	co := make([]coordinates, 1)
	for i := 0; i < 10; i++ {
		co = append(co, coordinates{0, 0})
	}
	p := &position2{pos: co, visited: make(map[string]struct{})}
	p.visited[posToStr(p.pos[9].x, p.pos[9].y)] = struct{}{}

	for _, line := range lines {
		if line != "" {
			p.move(line)
		}
	}

	fmt.Printf("%v positions visited by the tail\n", len(p.visited))
}

func main() {
	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
	solution2(lines)
}
