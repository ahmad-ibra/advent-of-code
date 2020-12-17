package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func applyMask(mask string, addr uint64, num uint64, mem map[uint64]uint64) map[uint64]uint64 {

	andMask := strings.Replace(mask, "X", "1", -1)
	andMaskUint, _ := strconv.ParseUint(andMask, 2, 64)
	orMask := strings.Replace(mask, "X", "0", -1)
	orMaskUint, _ := strconv.ParseUint(orMask, 2, 64)

	maskedNum := (num & andMaskUint) | orMaskUint
	mem[addr] = maskedNum

	return mem
}

func part1(lines []string) {
	mask := ""
	mem := make(map[uint64]uint64)

	for _, line := range lines {
		lineParts := strings.Split(line, " = ")

		// set to mask
		if lineParts[0] == "mask" {
			mask = lineParts[1]
		} else {
			// memory address
			addrLoc := lineParts[0][4 : len(lineParts[0])-1]
			addrInt, _ := strconv.Atoi(addrLoc)
			addr := uint64(addrInt)
			// value
			numInt, _ := strconv.Atoi(lineParts[1])
			num := uint64(numInt)

			mem = applyMask(mask, addr, num, mem)
		}
	}

	output := uint64(0)
	for _, val := range mem {
		output += val
	}

	fmt.Println("P1: ", output)
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	part1(lines)
}
