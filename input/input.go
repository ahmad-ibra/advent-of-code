package input

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/Ahmad-Ibra/advent-of-code-2020/errorHandler"
)

// ReadLines takes a path to a file and returns an array of the lines
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	errorHandler.Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// PathToFile generates a path to the given filename
func PathToFile(fName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	errorHandler.Check(err)

	return filepath.Join(dir, fName)
}

// GroupLines takes the output of readLines and returns an array of grouped lines
// The groups are separated by new lines
func GroupLines(lines []string) [][]string {

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
