package main

import (
	in "aoc2022/input"
	"fmt"
	"strconv"
	"strings"
)

type monkey struct {
	inspectionCount uint64
	items           []uint64
	operation       string
	test            uint64
	trueMonkey      uint64
	falseMonkey     uint64
}

func (m *monkey) testItem(itm uint64, superMod uint64, divisor uint64) (uint64, uint64) {
	ans := calc(itm, m.operation, divisor)

	if ans%m.test == 0 {
		return ans % superMod, m.trueMonkey
	}
	return ans % superMod, m.falseMonkey
}

func calc(val1 uint64, op string, divisor uint64) uint64 {
	parts := strings.Split(op, " ")
	var val2 uint64
	v, err := strconv.Atoi(parts[1])
	if err != nil {
		val2 = val1
	} else {
		val2 = uint64(v)
	}

	var ans uint64
	switch parts[0] {
	case "+":
		ans = val1 + val2
		break
	case "*":
		ans = val1 * val2
	default:
		fmt.Println("UNEXPECTED CASE IN CALC FUNCTION")
	}

	return ans / divisor
}

func solution(lines []string) {
	monkeys1 := []*monkey{}
	monkeys2 := []*monkey{}

	m1 := &monkey{inspectionCount: 0, items: []uint64{}}
	m2 := &monkey{inspectionCount: 0, items: []uint64{}}
	for _, line := range lines {
		if line == "" {
			monkeys1 = append(monkeys1, m1)
			monkeys2 = append(monkeys2, m2)
			m1 = &monkey{inspectionCount: 0, items: []uint64{}}
			m2 = &monkey{inspectionCount: 0, items: []uint64{}}
		} else if strings.Contains(line, "  Starting items: ") {
			l := strings.TrimLeft(line, "  Starting items: ")
			strItems := strings.Split(l, ", ")
			intItems := []uint64{}
			for _, it := range strItems {
				it, _ := strconv.Atoi(it)
				intItems = append(intItems, uint64(it))
			}

			m1.items = append(m1.items, intItems...)
			m2.items = append(m2.items, intItems...)

		} else if strings.Contains(line, "  Operation: new = old ") {
			m1.operation = strings.TrimLeft(line, "  Operation: new = old ")
			m2.operation = strings.TrimLeft(line, "  Operation: new = old ")

		} else if strings.Contains(line, "  Test: divisible by ") {
			l := strings.TrimLeft(line, "  Test: divisible by ")
			t, _ := strconv.Atoi(l)
			m1.test = uint64(t)
			m2.test = uint64(t)

		} else if strings.Contains(line, "    If true: throw to monkey ") {
			l := strings.TrimLeft(line, "    If true: throw to monkey ")
			mNumber, _ := strconv.Atoi(l)
			m1.trueMonkey = uint64(mNumber)
			m2.trueMonkey = uint64(mNumber)

		} else if strings.Contains(line, "    If false: throw to monkey ") {
			l := strings.TrimLeft(line, "    If false: throw to monkey ")
			mNumber, _ := strconv.Atoi(l)
			m1.falseMonkey = uint64(mNumber)
			m2.falseMonkey = uint64(mNumber)

		} else if strings.Contains(line, "Monkey ") {
			continue
		} else {
			fmt.Printf("FOUND UNEXPECTED LINE:\n%v\n", line)
		}
	}
	monkeys1 = monkeys1[:len(monkeys1)-1]
	monkeys2 = monkeys2[:len(monkeys2)-1]

	superMod := uint64(1)
	for _, m := range monkeys1 {
		superMod *= m.test
	}

	// go through 20 rounds
	for i := 0; i < 20; i++ { // part 1
		for _, m := range monkeys1 {

			for j := 0; j < len(m.items); j++ {
				itm := m.items[j]
				newVal, nextMonkey := m.testItem(itm, superMod, 3)
				monkeys1[nextMonkey].items = append(monkeys1[nextMonkey].items, newVal)
				m.inspectionCount++
			}
			m.items = []uint64{}
		}
	}

	inspections := []uint64{0, 0}

	for _, m := range monkeys1 {
		ic := m.inspectionCount
		if ic > inspections[0] && ic > inspections[1] {
			inspections[0] = inspections[1]
			inspections[1] = ic
		} else if ic > inspections[0] {
			inspections[0] = ic
		}
	}
	fmt.Printf("PART1: %v, monkey business = %v\n", inspections, inspections[0]*inspections[1])

	for i := 0; i < 10000; i++ { // part 2
		for _, m := range monkeys2 {

			for j := 0; j < len(m.items); j++ {
				itm := m.items[j]
				newVal, nextMonkey := m.testItem(itm, superMod, 1)
				monkeys2[nextMonkey].items = append(monkeys2[nextMonkey].items, newVal)
				m.inspectionCount++
			}
			m.items = []uint64{}
		}
	}

	inspections = []uint64{0, 0}

	for _, m := range monkeys2 {
		ic := m.inspectionCount
		if ic > inspections[0] && ic > inspections[1] {
			inspections[0] = inspections[1]
			inspections[1] = ic
		} else if ic > inspections[0] {
			inspections[0] = ic
		}
	}
	fmt.Printf("PART2: %v, monkey business = %v\n", inspections, inspections[0]*inspections[1])
}

func main() {
	lines, err := in.InputSplitByLine("input.txt")
	if err != nil {
		panic(err)
	}

	solution(lines)
}
