package main

import (
	"math"
	"strconv"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

type waypoint struct {
	// directions are relative to the ship
	NS int
	EW int
}

type position struct {
	// P1
	NS  int
	EW  int
	dir direction

	// P2
	WP waypoint
}

func (p *position) moveForward(units int) {
	switch p.dir {
	case north:
		p.NS += units
	case east:
		p.EW += units
	case south:
		p.NS -= units
	case west:
		p.EW -= units
	}
}

func (p *position) moveNorth(units int) {
	p.NS += units
}

func (p *position) moveEast(units int) {
	p.EW += units
}

func (p *position) moveSouth(units int) {
	p.NS -= units
}

func (p *position) moveWest(units int) {
	p.EW -= units
}

func (p *position) turnRight(degrees int) {
	units := degrees / 90
	newDir := (int(p.dir) + units) % 4
	p.dir = direction(newDir)
}

func (p *position) turnLeft(degrees int) {
	units := degrees / 90
	newDir := (int(p.dir) - units) % 4
	if newDir == -1 {
		newDir = 3
	}
	if newDir == -2 {
		newDir = 2
	}
	if newDir == -3 {
		newDir = 1
	}
	p.dir = direction(newDir)
}

func (p *position) moveForwardP2(units int) {
	p.NS += p.WP.NS * units
	p.EW += p.WP.EW * units
}

func (p *position) moveNorthP2(units int) {
	p.WP.NS += units
}

func (p *position) moveEastP2(units int) {
	p.WP.EW += units
}

func (p *position) moveSouthP2(units int) {
	p.WP.NS -= units
}

func (p *position) moveWestP2(units int) {
	p.WP.EW -= units
}

func (p *position) turnRightP2(degrees int) {
	units := degrees / 90
	quarterRotations := units % 4

	// rotate 90 degrees clockwise
	for i := 0; i < quarterRotations; i++ {
		tempEW := p.WP.EW
		p.WP.EW = p.WP.NS
		p.WP.NS = tempEW * -1

	}
}

func (p *position) turnLeftP2(degrees int) {
	units := degrees / 90
	quarterRotations := units % 4

	// rotate 90 degrees counter-clockwise
	for i := 0; i < quarterRotations; i++ {
		tempEW := p.WP.EW
		p.WP.EW = p.WP.NS * -1
		p.WP.NS = tempEW
	}
}

func (p *position) manhattenDistance() float64 {
	return math.Abs(float64(p.NS)) + math.Abs(float64(p.EW))
}

// main logic for parsing input
func moveP2(p *position, input string) {
	action := string(input[0])
	units, _ := strconv.Atoi(string(input[1:]))

	switch action {
	case "N":
		p.moveNorthP2(units)
	case "E":
		p.moveEastP2(units)
	case "S":
		p.moveSouthP2(units)
	case "W":
		p.moveWestP2(units)
	case "F":
		p.moveForwardP2(units)
	case "R":
		p.turnRightP2(units)
	case "L":
		p.turnLeftP2(units)
	default:
		panic("unknown action")
	}
}

func move(p *position, input string) {
	action := string(input[0])
	units, _ := strconv.Atoi(string(input[1:]))

	switch action {
	case "N":
		p.moveNorth(units)
	case "E":
		p.moveEast(units)
	case "S":
		p.moveSouth(units)
	case "W":
		p.moveWest(units)
	case "F":
		p.moveForward(units)
	case "R":
		p.turnRight(units)
	case "L":
		p.turnLeft(units)
	default:
		panic("unknown action")
	}
}
