package input

import (
	"os"
	"strings"
)

func InputSplitByLine(fileLoc string) ([]string, error) {

	dat, err := os.ReadFile(fileLoc)
	if err != nil {
		return nil, err
	}

	groups := strings.Split(string(dat), "\n")

	return groups, nil
}
