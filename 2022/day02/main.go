package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Round struct {
	theirHand string
	myHand    string
}

func createRound(hands []string) *Round {
	r := &Round{
		theirHand: hands[0],
		myHand:    hands[1],
	}

	return r
}

func (r *Round) shapeScore() int {
	score := 0

	switch r.myHand {
	case "X":
		score += 1
		break
	case "Y":
		score += 2
		break
	case "Z":
		score += 3
		break
	default:
		panicMsg, _ := fmt.Printf("Unhandled myHand val %v", r.myHand)
		panic(panicMsg)
	}

	fmt.Println("shapescore is " + strconv.Itoa(score))
	return score
}

func (r *Round) outcomeScore() int {
	score := 0
	// outcome score
	if r.isTie() {
		score += 3
	} else if (r.theirHand == "A" && r.myHand == "Y") || (r.theirHand == "B" && r.myHand == "Z") || (r.theirHand == "C" && r.myHand == "X") {
		score += 6
	}

	fmt.Println("outcomeScore is " + strconv.Itoa(score))
	return score
}

func (r *Round) isTie() bool {
	return (r.theirHand == "A" && r.myHand == "X") || (r.theirHand == "B" && r.myHand == "Y") || (r.theirHand == "C" && r.myHand == "Z")
}

func main() {
	// A = Rock = X <- 1
	// B = Paper = Y <- 2
	// C = Scisors = Z <- 3

	// if they have A, i have Y
	// if they have B, i have Z
	// if they have C, i have X

	// totalScore = sum of all roundScore
	// roundScore = shapeScore + outcomeScore
	// outcomeScore = 6 for win, 3 for draw, 0 for loss

	rounds, err := inputSplitByLine("day02/input.txt")
	if err != nil {
		panic(err)
	}

	totalScore := 0
	for _, r := range rounds {
		if r == "" {
			continue
		}
		//fmt.Println("round " + strconv.Itoa(i) + ": " + r)

		round := createRound(strings.Fields(r))
		totalScore += round.shapeScore() + round.outcomeScore()
		fmt.Println()
	}

	fmt.Println("part1 answer: " + strconv.Itoa(totalScore))
}

func inputSplitByLine(fileLoc string) ([]string, error) {

	dat, err := os.ReadFile(fileLoc)
	if err != nil {
		return nil, err
	}

	groups := strings.Split(string(dat), "\n")

	return groups, nil
}
