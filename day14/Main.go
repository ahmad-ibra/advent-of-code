package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func applyMaskP1(mask string, addr uint64, num uint64, mem map[uint64]uint64) map[uint64]uint64 {

	andMask := strings.Replace(mask, "X", "1", -1)
	andMaskUint, _ := strconv.ParseUint(andMask, 2, 64)
	orMask := strings.Replace(mask, "X", "0", -1)
	orMaskUint, _ := strconv.ParseUint(orMask, 2, 64)

	maskedNum := (num & andMaskUint) | orMaskUint
	mem[addr] = maskedNum

	return mem
}

func setVal(address string, num uint64, mem map[uint64]uint64) {
	for i := 0; i < len(address); i++ {
		if string(address[i]) == "X" {
			// replace the X with 0 and call setVal on the new address
			var sb0 strings.Builder
			sb0.WriteString(address[0:i])
			sb0.WriteString("0")
			sb0.WriteString(address[i+1:])
			setVal(sb0.String(), num, mem)

			// replace the X with 0 and call setVal on the new address
			var sb1 strings.Builder
			sb1.WriteString(address[0:i])
			sb1.WriteString("1")
			sb1.WriteString(address[i+1:])
			setVal(sb1.String(), num, mem)
			return
		}
	}
	add, _ := strconv.ParseUint(address, 2, 64)
	mem[add] = num
}

func applyMaskP2(mask string, addr uint64, num uint64, mem map[uint64]uint64) map[uint64]uint64 {

	binAddr := strconv.FormatUint(addr, 2)
	lenAddr := len(binAddr)
	diffLen := 36 - lenAddr

	var sb strings.Builder
	for i, str := range mask {

		// address is 0
		if i < diffLen {
			if string(str) == "X" {
				sb.WriteString("X")
			} else {
				sb.WriteRune(str)
			}

		} else {
			//compare address value and map value
			if string(str) == "0" {
				sb.WriteByte(binAddr[i-diffLen])
			} else {
				sb.WriteRune(str)
			}
		}
	}
	newAddr := sb.String()
	setVal(newAddr, num, mem)
	return mem
}

func readInstructions(lines []string) {
	mask := ""
	memP1 := make(map[uint64]uint64)
	memP2 := make(map[uint64]uint64)

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

			memP1 = applyMaskP1(mask, addr, num, memP1)
			memP2 = applyMaskP2(mask, addr, num, memP2)
		}
	}

	outputP1 := uint64(0)
	for _, val := range memP1 {
		outputP1 += val
	}
	outputP2 := uint64(0)
	for _, val := range memP2 {
		outputP2 += val
	}

	fmt.Println("P1: ", outputP1)
	fmt.Println("P2: ", outputP2)
}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	readInstructions(lines)
}
