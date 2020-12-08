package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/Ahmad-Ibra/advent-of-code-2020/input"
	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

func testStartup(line []string) {

	for j := 0; j < len(line); j++ {
		testLine := line
		checkInstruction := strings.Split(testLine[j], " ")

		// if its an instruction we can switch
		if checkInstruction[0] != "acc" {
			if checkInstruction[0] == "nop" {
				testLine[j] = strings.Replace(testLine[j], "nop", "jmp", -1)
			} else {
				testLine[j] = strings.Replace(testLine[j], "jmp", "nop", -1)
			}
			// we have now replaced opperation, run the startup2
			success := runStartupP2(testLine)
			if success {
				break
			}
			// reverting testLine
			// not sure why i need this, should be getting reset on line 15
			if checkInstruction[0] == "nop" {
				testLine[j] = strings.Replace(testLine[j], "jmp", "nop", -1)
			} else {
				testLine[j] = strings.Replace(testLine[j], "nop", "jmp", -1)
			}
		}
	}
}

func runStartupP2(line []string) bool {
	goal := len(line)
	isSuccessfull := false

	// current instruction
	i := 0
	acc := 0

	visited := make(map[int]bool)

	for {
		if visited[i] {
			break
		}
		visited[i] = true
		if visited[goal] {
			fmt.Println("success!")
			isSuccessfull = true
			break
		}

		instruction := strings.Split(line[i], " ")
		opp := instruction[0]
		dir := string(instruction[1][0])
		delta, _ := strconv.Atoi(instruction[1][1:])

		switch opp {
		case "acc":
			if dir == "+" {
				acc += delta
				i++
			} else {
				acc -= delta
				i++
			}
		case "jmp":
			if dir == "+" {
				i += delta
			} else {
				i -= delta
			}
		case "nop":
			i++
		default:
			panic("unhandled case")
		}
	}

	if isSuccessfull {
		fmt.Println("acc: ", acc)
	}

	return isSuccessfull
}

func runStartup(line []string) {
	// current instruction
	i := 0
	acc := 0

	visited := make(map[int]bool)

	for {
		if visited[i] {
			fmt.Println("loop")
			break
		}
		visited[i] = true

		instruction := strings.Split(line[i], " ")
		opp := instruction[0]
		dir := string(instruction[1][0])
		delta, _ := strconv.Atoi(instruction[1][1:])

		switch opp {
		case "acc":
			if dir == "+" {
				acc += delta
				i++
			} else {
				acc -= delta
				i++
			}
		case "jmp":
			if dir == "+" {
				i += delta
			} else {
				i -= delta
			}
		case "nop":
			i++
		default:
			panic("unhandled case")
		}
	}

	fmt.Println("acc: ", acc)

}

func main() {

	path := input.PathToFile("input.txt")

	lines, err := input.ReadLines(path)
	panicer.Check(err)

	runStartup(lines)
	testStartup(lines)
}
