package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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

// passport ------------------------------
type passport struct {
	byr string // birth year
	iyr string // issue year
	eyr string // expiration year
	hgt string // height
	hcl string // hair colour
	ecl string // eye colour
	pid string // passport id
	cid string // country id
}

func (p *passport) isValidPart1() bool {
	return (p.byr != "") && (p.iyr != "") && (p.eyr != "") && (p.hgt != "") && (p.hcl != "") && (p.ecl != "") && (p.pid != "")
}

// Enforce new validation rules
func (p *passport) isValidPart2() bool {
	return p.isValidByr() && p.isValidIyr() && p.isValidEyr() && p.isValidHgt() && p.isValidHcl() && p.isValidEcl() && p.isValidPid()
}

func (p *passport) isValidByr() bool {
	i, err := strconv.Atoi(p.byr)
	if err != nil {
		return false
	}
	if 1920 <= i {
		if i <= 2002 {
			return true
		}
	}
	return false
}

func (p *passport) isValidIyr() bool {
	i, err := strconv.Atoi(p.iyr)
	if err != nil {
		return false
	}
	if 2010 <= i {
		if i <= 2020 {
			return true
		}
	}
	return false
}

func (p *passport) isValidEyr() bool {
	i, err := strconv.Atoi(p.eyr)
	if err != nil {
		return false
	}
	if 2020 <= i {
		if i <= 2030 {
			return true
		}
	}
	return false
}

func (p *passport) isValidHgt() bool {
	// a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	match, _ := regexp.MatchString("\\d+(cm|in)", p.hgt)

	len := len(p.hgt)
	if match {
		if string(p.hgt[len-2]) == "c" {
			// cm
			sNum := p.hgt[0 : len-2]
			i, err := strconv.Atoi(sNum)
			if err != nil {
				return false
			}
			if 150 <= i {
				if i <= 193 {
					return true
				}
			}
		} else {
			// in
			sNum := p.hgt[0 : len-2]
			i, err := strconv.Atoi(sNum)
			if err != nil {
				return false
			}
			if 59 <= i {
				if i <= 76 {
					return true
				}
			}
		}
	}
	return false
}

func (p *passport) isValidHcl() bool {
	// a # followed by exactly six characters 0-9 or a-f
	match, _ := regexp.MatchString("#[a-f0-9]{6}", p.hcl)
	return match && (len(p.hcl) == 7)
}

func (p *passport) isValidEcl() bool {
	switch p.ecl {
	case "amb":
		return true
	case "blu":
		return true
	case "brn":
		return true
	case "gry":
		return true
	case "grn":
		return true
	case "hzl":
		return true
	case "oth":
		return true
	default:
		return false
	}
}

func (p *passport) isValidPid() bool {
	match, _ := regexp.MatchString("\\d{9}", p.pid)
	return match && (len(p.pid) == 9)
}

func (p *passport) setPassportValue(key string, val string) {
	switch key {
	case "byr":
		p.byr = val
	case "iyr":
		p.iyr = val
	case "eyr":
		p.eyr = val
	case "hgt":
		p.hgt = val
	case "hcl":
		p.hcl = val
	case "ecl":
		p.ecl = val
	case "pid":
		p.pid = val
	case "cid":
		p.cid = val
	default:
		panic("Unhandled case!")
	}

}

//------------------------------

func split(r rune) bool {
	return r == ':' || r == ' '
}

func countValidPassportsP1(passports []passport) int {

	validCount := 0
	for _, p := range passports {

		if p.isValidPart1() {
			validCount++
		}
	}
	return validCount
}

func countValidPassportsP2(passports []passport) int {

	validCount := 0
	for _, p := range passports {

		if p.isValidPart2() {
			validCount++
		}
	}
	return validCount
}

func validPassportCount(lines []string) {

	// init variables
	passports := make([]passport, 0)
	passportNum := -1
	createNew := true

	for _, line := range lines {
		if createNew {
			// creating a passport
			var p passport
			passports = append(passports, p)
			passportNum++
			createNew = false
		}

		// pointer to most recently created passport
		p := &passports[passportNum]

		// passport fields
		fields := strings.FieldsFunc(line, split)

		for i := 0; i < len(fields); i = i + 2 {
			p.setPassportValue(fields[i], fields[i+1])
		}

		// if end of current passport
		if line == "" {
			// create a new passport on next line read
			createNew = true
		}
	}

	fmt.Println("P1 valid passports: ", countValidPassportsP1(passports))
	fmt.Println("P2 valid passports: ", countValidPassportsP2(passports))
}

func main() {

	path := pathToFile("input.txt")

	lines, err := readLines(path)
	check(err)

	validPassportCount(lines)
}
