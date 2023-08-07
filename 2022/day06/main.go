package main

import (
	in "aoc2022/input"
	"fmt"
)

type markMap struct {
	packetMap   map[string]int
	packetCount int

	mesageMap    map[string]int
	messageCount int
}

func (mm *markMap) foundMarkP() bool {
	if len(mm.packetMap) >= mm.packetCount {
		return true
	}
	return false
}

func (mm *markMap) foundMarkM() bool {
	if len(mm.mesageMap) >= mm.messageCount {
		return true
	}
	return false
}

func (mm *markMap) add(s string) {
	if val, ok := mm.packetMap[s]; ok {
		mm.packetMap[s] = val + 1
	} else {
		mm.packetMap[s] = 1
	}

	if val, ok := mm.mesageMap[s]; ok {
		mm.mesageMap[s] = val + 1
	} else {
		mm.mesageMap[s] = 1
	}

}

func (mm *markMap) removeP(s string) {
	if val, ok := mm.packetMap[s]; ok {
		if val > 1 {
			mm.packetMap[s] = val - 1
		} else {
			delete(mm.packetMap, s)
		}
	}
}

func (mm *markMap) removeM(s string) {
	if val, ok := mm.mesageMap[s]; ok {
		if val > 1 {
			mm.mesageMap[s] = val - 1
		} else {
			delete(mm.mesageMap, s)
		}
	}
}

func getMarkerIndex(line string, pIdx int, mIdx int) (int, int) {

	mm := &markMap{packetMap: make(map[string]int), packetCount: pIdx, mesageMap: make(map[string]int), messageCount: mIdx}

	front := 0
	backP := 0 - pIdx
	backM := 0 - mIdx

	ansP := -1
	ansM := -1

	for front < len(line) {
		f := string(line[front])
		mm.add(f)
		if backP >= 0 {
			b := string(line[backP])
			mm.removeP(b)
		}
		if backM >= 0 {
			b := string(line[backM])
			mm.removeM(b)
		}

		if ansP == -1 {
			if mm.foundMarkP() {
				ansP = front
			}
		}

		if ansM == -1 {
			if mm.foundMarkM() {
				ansM = front
				return ansP, ansM
			}
		}
		front++
		backP++
		backM++
	}

	return ansP, ansM
}

func solution(lines []string) {
	for _, line := range lines {
		if line == "" {
			continue
		}

		pCount := 4
		mCount := 14
		packetIdx, msgIdx := getMarkerIndex(line, pCount, mCount)
		fmt.Printf("Start of packet at character %v, Start of message at character %v\n", packetIdx+1, msgIdx+1)
	}
}

func main() {

	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
}
