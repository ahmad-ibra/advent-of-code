package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	START = "mul("
	SEP   = ","
	END   = ")"
	DO    = "do()"
	DONT  = "don't()"
)

func main() {

	bytes, err := os.ReadFile("input.txt")
	//bytes, err := os.ReadFile("sample.txt")
	//bytes, err := os.ReadFile("sample2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(bytes), "\n")

	out := 0
	do := true
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		for i := 0; i < len(line)-8; i++ { // -8 because of the length of "mul(1,2)"
			// part 2 additional requirement --------------
			// comment out for part 1 answer
			if !do {
				// check for do
				if line[i:i+4] == DO {
					do = true
					continue
				}
				continue
			}

			if do {
				// check for dont
				if line[i:i+7] == DONT {
					do = false
					continue
				}
			}
			if line[i:i+4] != START {
				continue
			}
			// part 2 additional requirement end ----------

			var num1, num2 int
			sepDist := 7
			num1, err = strconv.Atoi(line[i+4 : i+7])
			if err != nil {
				sepDist = 6
				num1, err = strconv.Atoi(line[i+4 : i+6])
				if err != nil {
					sepDist = 5
					num1, err = strconv.Atoi(line[i+4 : i+5])
					if err != nil {
						continue
					}
				}
			}

			if line[i+sepDist:i+sepDist+1] != SEP {
				continue
			}

			endDist := sepDist + 4
			num2, err = strconv.Atoi(line[i+sepDist+1 : i+sepDist+4])
			if err != nil {
				endDist = sepDist + 3
				num2, err = strconv.Atoi(line[i+sepDist+1 : i+sepDist+3])
				if err != nil {
					endDist = sepDist + 2
					num2, err = strconv.Atoi(line[i+sepDist+1 : i+sepDist+2])
					if err != nil {
						continue
					}
				}
			}

			if line[i+endDist:i+endDist+1] != END {
				continue
			}

			out += num1 * num2
		}
	}

	fmt.Println(out)
}
