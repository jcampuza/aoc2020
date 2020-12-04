package utils

import (
	"strconv"
	"strings"
)

func SliceAtoi(str []string) []int {
	tmp := []int{}

	for _, v := range str {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		tmp = append(tmp, val)
	}

	return tmp
}

func FilterString(str []string, compare func(string) bool) []string {
	tmp := []string{}

	for _, v := range str {
		if compare(v) {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

func SplitNewlines(str string) []string {
	return strings.Split(str, "\n")
}

func GetLines(input []byte) []string {
	return FilterString(SplitNewlines(string(input)), func(s string) bool {
		return s != ""
	})
}

func GetLinesWithEmptyLines(input []byte) []string {
	return SplitNewlines(string(input))
}

func TrimLines(lines []string) []string {
	for i, line := range lines {
		lines[i] = strings.TrimSpace(line)
	}

	return lines
}
