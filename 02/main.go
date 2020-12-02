package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

type Password struct {
	Low, High int
	Letter    byte
	Value     string
}

func isValidPolicyOne(password Password) bool {
	count := 0

	for _, char := range password.Value {
		if byte(char) == password.Letter {
			count++
		}
	}

	return count >= password.Low && count <= password.High
}

func isValidPolicyTwo(password Password) bool {
	positionOne, positionTwo := password.Value[password.Low-1], password.Value[password.High-1]

	if positionOne == password.Letter {
		return positionTwo != password.Letter
	}

	return positionTwo == password.Letter
}

func parsePassword(str string) Password {
	parsed := strings.Split(str, " ")
	minMaxWithDash, letter, value := parsed[0], parsed[1][0], parsed[2]
	minMax := strings.Split(minMaxWithDash, "-")
	min, _ := strconv.Atoi(minMax[0])
	max, _ := strconv.Atoi(minMax[1])

	return Password{Low: min, High: max, Letter: letter, Value: value}
}

func countValidPasswords(passwords []string, isValidFunc func(password Password) bool) int {
	invalidPasswords := 0

	for _, password := range passwords {
		if isValidFunc(parsePassword(password)) {
			invalidPasswords++
		}
	}

	return invalidPasswords
}

func main() {
	p := filepath.Join("02", "02.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	inputStr := string(input)
	lines := utils.FilterString(utils.SplitNewlines(inputStr), func(s string) bool {
		return s != ""
	})

	fmt.Println("Part one:", countValidPasswords(lines, isValidPolicyOne))
	fmt.Println("Part two:", countValidPasswords(lines, isValidPolicyTwo))
}
