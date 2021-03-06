package utils

import (
	"log"
	"strconv"
	"strings"
	"time"
)

const MaxUint = ^uint(0)

const MinUint = 0

const MaxInt = int(MaxUint >> 1)

const MinInt = -MaxInt - 1

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

func TrimLine(line string) string {
	return strings.TrimSpace(line)
}

func TrimLines(lines []string) []string {
	for i, line := range lines {
		lines[i] = TrimLine(line)
	}

	return lines
}

func GetLinesInt(input []byte) []int {
	ints := []int{}

	for _, line := range GetLines(input) {
		v, _ := strconv.Atoi(line)
		ints = append(ints, v)
	}

	return ints
}

// Just calculat both since it can be done with one loop
func MinMax(nums ...int) (min, max int) {
	min = nums[0]
	max = nums[0]
	for _, value := range nums {
		if min > value {
			min = value
		}

		if max < value {
			max = value
		}
	}

	return min, max
}

func Min(nums ...int) int {
	min, _ := MinMax(nums...)
	return min
}

func Max(nums ...int) int {
	_, max := MinMax(nums...)
	return max
}

func TrackTime(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func ReplaceAtIndex(str string, replace rune, index int) string {
	out := []rune(str)
	out[index] = replace
	return string(out)
}

func AbsI(v int) int {
	if v < 0 {
		return -v
	}

	return v
}

func ParseInt(str string) int {
	v, _ := strconv.Atoi(str)
	return v
}
