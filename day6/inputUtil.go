package main

import (
	"bufio"
	"os"
	"path/filepath"
)

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

/*
This function takes the output of readLines and returns an array of grouped lines.
The groups are separated by new lines
*/
func groupLines(lines []string) [][]string {

	groupedLines := [][]string{}
	currentLine := ""

	row := []string{}
	groupedLines = append(groupedLines, row)

	for _, line := range lines {
		currentLine = line
		if currentLine == "" {
			groupedLines = append(groupedLines, row)
			row = nil
		} else {
			row = append(row, currentLine)
		}
	}

	// the inputs have tended to not end with an empty line, so we need to handle adding
	row = append(row, currentLine)
	groupedLines = append(groupedLines, row)

	return groupedLines
}
