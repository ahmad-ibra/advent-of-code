package main

import (
	in "aoc2022/input"
	"fmt"
	"strconv"
	"strings"
)

const (
	COMMAND_LS         = "$ ls"
	PREFIX_DIR         = "dir "
	COMMAND_CD_BACK    = "$ cd .."
	COMMAND_CD_FORWARD = "$ cd "
)

type dirNode struct {
	name        string
	fileSize    int
	subDirNodes map[string]*dirNode
	parent      *dirNode
}

func isCommand(line string) bool {
	return strings.HasPrefix(line, "$ ")
}

func getSizes(dir *dirNode, memo map[*dirNode]int) int {
	if val, ok := memo[dir]; ok {
		return val
	}

	size := 0
	for _, val := range dir.subDirNodes {
		size += getSizes(val, memo)
	}
	size += dir.fileSize
	memo[dir] = size

	return size
}

func solution(lines []string) {

	root := &dirNode{name: "/", parent: nil, subDirNodes: make(map[string]*dirNode)}
	cur := root

	for i := 1; i < len(lines); i++ {
		line := lines[i]

		if line == "" || line == COMMAND_LS {
			continue
		} else if strings.HasPrefix(line, PREFIX_DIR) {
			parts := strings.Split(line, " ")
			child := &dirNode{name: parts[1], parent: cur, subDirNodes: make(map[string]*dirNode)}
			cur.subDirNodes[parts[1]] = child

		} else if line == COMMAND_CD_BACK {
			cur = cur.parent

		} else if strings.HasPrefix(line, COMMAND_CD_FORWARD) {
			// make cur the child thats defined in the line
			parts := strings.Split(line, " ")
			cur = cur.subDirNodes[parts[2]]

		} else {
			// its a file, add size to current node
			parts := strings.Split(line, " ")
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				fmt.Printf("FAILED TO GET FILE SIZE AT LINE %v\n", line)
				panic(err)
			}
			cur.fileSize += size
		}
	}

	sizes := make(map[*dirNode]int)
	usedSpace := getSizes(root, sizes)

	sizeLimit := 100000
	total := 0
	for _, size := range sizes {
		if size <= sizeLimit {
			total += size
		}

	}
	fmt.Printf("PART 1: SIZE OF DIRECTORIES AT MOST %v is %v\n", sizeLimit, total)

	totalDiskSpace := 70000000
	requiredUnusedDiskSPace := 30000000
	freeSpace := totalDiskSpace - usedSpace
	spaceToDelete := requiredUnusedDiskSPace - freeSpace

	minSize := totalDiskSpace
	for _, size := range sizes {
		if size >= spaceToDelete && size < minSize {
			minSize = size
		}
	}

	fmt.Printf("PART 2: DELETING FOLDER WITH %v DISK SPACE\n", minSize)

}

func main() {

	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
}
