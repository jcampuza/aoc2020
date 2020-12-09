package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func FindTwoWithSum(values []int, sum int) int {
	m := make(map[int]bool)

	for _, v := range values {
		rem := sum - v

		if _, ok := m[rem]; ok {
			return v * rem
		}

		m[v] = true
	}

	return -1
}

func FindThreeWithSum(values []int, sum int) int {
	m := make(map[int]bool)

	for i, v1 := range values {
		for _, v2 := range values[:i] {
			rem := sum - v1 - v2

			if _, ok := m[rem]; ok {
				return v1 * v2 * rem
			}
		}

		m[v1] = true
	}

	return -1
}

func main() {
	p := filepath.Join("01", "input.txt")
	data, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	str := string(data)

	lines := utils.FilterString(strings.Split(str, "\n"), func(s string) bool {
		return s != ""
	})

	nums := utils.SliceAtoi(lines)

	fmt.Println("Part one:", FindTwoWithSum(nums, 2020))
	fmt.Println("Part two:", FindThreeWithSum(nums, 2020))
}
