package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func findTwoWithSum(values []int) int {
	for i, one := range values {
		for _, two := range values[i:] {
			if (one + two) == 2020 {
				return one * two
			}
		}
	}

	return -1
}

func findThreeWithSum(values []int) int {
	for i, one := range values {
		for j, two := range values[i:] {
			for _, three := range values[j:] {
				if one+two+three == 2020 {
					return one * two * three
				}
			}
		}
	}

	return -1
}

func main() {
	p := filepath.Join("01", "01.txt")
	data, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	str := string(data)

	lines := utils.FilterString(strings.Split(str, "\n"), func(s string) bool {
		return s != ""
	})

	nums := utils.SliceAtoi(lines)

	fmt.Println("Part one:", findTwoWithSum(nums))
	fmt.Println("Part two:", findThreeWithSum(nums))
}
