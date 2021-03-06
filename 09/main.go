package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

func isValid(nums []int, num int) bool {
	for i, v1 := range nums {
		for _, v2 := range nums[:i] {
			if (v1 + v2) == num {
				return true
			}
		}
	}

	return false
}

func _partOne(nums []int, preambleSize int) (int, int) {
	for i := range nums {
		preamble, value := nums[i:i+preambleSize], nums[i+preambleSize]

		if !isValid(preamble, value) {
			return value, i
		}
	}

	return -1, -1
}

func partOne(nums []int, preambleSize int) int {
	val, _ := _partOne(nums, preambleSize)
	return val
}

func findContigousSubarray(arr []int, val int) []int {
	start, end := 0, 1
	sum := arr[start]

	for end < len(arr) {
		for sum > val {
			sum -= arr[start]
			start++
		}

		if sum == val {
			return arr[start:end]
		}

		sum += arr[end]
		end++
	}

	return []int{}
}

func partTwo(nums []int, preambleSize int) int {
	val, _ := _partOne(nums, preambleSize)
	subarr := findContigousSubarray(nums, val)
	min, max := utils.MinMax(subarr...)

	return min + max
}

func main() {
	p := filepath.Join("09", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)

	// Turn into numbers
	nums := []int{}
	for _, line := range lines {
		v, _ := strconv.Atoi(line)
		nums = append(nums, v)
	}

	fmt.Println("Part one", partOne(nums, 25))
	fmt.Println("Part two", partTwo(nums, 25))
}
