package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func parsePassports(lines []string) []string {
	passports, curr := []string{}, ""

	for _, line := range lines {
		if line == "" {
			passports, curr = append(passports, curr), ""
		}

		curr = curr + " " + line
	}

	if curr != "" {
		passports = append(passports, curr)
	}

	return passports
}

func hasAllFields(m map[string]bool) bool {
	for _, value := range m {
		if !value {
			return false
		}
	}

	return true
}

func strNumberBetween(str string, min int, max int) bool {
	num, err := strconv.Atoi(str)
	if err != nil {
		return false
	}

	return num >= min && num <= max
}

func validatePassport(passport string) bool {
	fields := map[string]bool{
		"byr": false,
		"iyr": false,
		"eyr": false,
		"hgt": false,
		"hcl": false,
		"ecl": false,
		"pid": false,
		// "cid": false,
	}

	for key := range fields {
		if strings.Contains(passport, key) {
			fields[key] = true
		}
	}

	if hasAllFields(fields) {
		return true
	}

	return false
}

func validatePassportStrict(passport string) bool {
	if !validatePassport(passport) {
		return false
	}

	fields := strings.Split(passport, " ")

	for _, field := range fields {
		split := strings.Split(field, ":")
		left, right := split[0], split[1]

		switch left {
		case "byr":
			if !strNumberBetween(right, 1920, 2002) {
				return false
			}
		case "iyr":
			if !strNumberBetween(right, 2010, 2020) {
				return false
			}
		case "eyr":
			if !strNumberBetween(right, 2020, 2030) {
				return false
			}
		case "hgt":
			isInches, isCm := strings.HasSuffix(right, "in"), strings.HasSuffix(right, "cm")

			if !isInches && !isCm {
				return false
			}

			if isInches && !strNumberBetween(right[:len(right)-2], 59, 76) {
				return false
			}

			if isCm && !strNumberBetween(right[:len(right)-2], 150, 193) {
				return false
			}
		case "hcl":
			if match, _ := regexp.MatchString("^#([0-9a-f]){6}$", right); !match {
				return false
			}
		case "ecl":
			if match, _ := regexp.MatchString("^(amb|blu|brn|gry|grn|hzl|oth)$", right); !match {
				return false
			}
		case "pid":
			if match, _ := regexp.MatchString("^(\\d){9}$", right); !match {
				return false
			}
		case "cid":
			// Ignore
		}
	}

	return true
}

func countValidPassports(passports []string, validator func(password string) bool) int {
	valid := 0

	for _, passport := range passports {
		if validator(passport) {
			valid++
		}
	}

	return valid
}

func main() {
	p := filepath.Join("04", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLinesWithEmptyLines(input)
	passports := utils.TrimLines(parsePassports(lines))

	fmt.Printf("Part one: %d\n", countValidPassports(passports, validatePassport))
	fmt.Printf("Part two: %d\n", countValidPassports(passports, validatePassportStrict))
}
