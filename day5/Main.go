package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// IO -------------------------------------------
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

//-----------------------------------------------

func convertRowToInt(row string) int64 {
	// replace with binary
	row = strings.Replace(row, "F", "0", -1)
	row = strings.Replace(row, "B", "1", -1)

	// convert binary to int
	rowInt, err := strconv.ParseInt(row, 2, 64)
	check(err)

	return rowInt
}

func convertColToInt(col string) int64 {
	// replace with binary
	col = strings.Replace(col, "L", "0", -1)
	col = strings.Replace(col, "R", "1", -1)

	// convert binary to int
	colInt, err := strconv.ParseInt(col, 2, 64)
	check(err)

	return colInt
}

func genSeatID(row int64, col int64) int64 {

	return (row * 8) + col
}

func highestSeatID(lines []string) {

	maxID := 0

	mapOfSeats := make(map[int64]bool)

	for _, line := range lines {
		row := line[0:7]
		col := line[7:10]

		rowInt := convertRowToInt(row)
		colInt := convertColToInt(col)

		seatID := genSeatID(rowInt, colInt)

		mapOfSeats[seatID] = true

		if seatID > int64(maxID) {
			maxID = int(seatID)
		}
	}

	fmt.Println("P1 max seat ID: ", maxID)

	for i := maxID; i > 0; i-- {
		if mapOfSeats[int64(i)] {
			// seat is in map
		} else {
			// seat isn't in map, must be my seat
			fmt.Println("P2 my seat ID: ", i)
			break
		}
	}
}

func main() {

	path := pathToFile("input.txt")

	lines, err := readLines(path)
	check(err)

	highestSeatID(lines)
}
