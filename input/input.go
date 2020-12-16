package input

import (
	"bufio"
	"os"
	"path/filepath"

	"github.com/Ahmad-Ibra/advent-of-code-2020/panicer"
)

// PathToFile generates a path to the given filename
func PathToFile(fName string) string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	panicer.Check(err)

	return filepath.Join(dir, fName)
}

// ByteToStringArray takes a byte array and converts it to a string array
func ByteToStringArray(contents []byte) []string {
	var output []string

	lastChar := 0
	for i := 0; i < len(contents); i++ {
		curChar := string(contents[i])
		if curChar == "\n" {
			line := string(contents[lastChar:i])
			output = append(output, line)
			lastChar = i
		}
	}
	//reached EOF, append remaining chars
	line := string(contents[lastChar:])
	output = append(output, line)
	return output
}

// ReadLines takes a path to a file and returns an array of the lines
func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	panicer.Check(err)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// GroupLines takes the output of readLines and returns an array of grouped lines
// The groups are separated by new lines
func GroupLines(lines []string) [][]string {
	groupedLines := [][]string{}
	currentLine := ""

	row := []string{}

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
	return append(groupedLines, row)
}
