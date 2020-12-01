package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
)

func findTwoWithSum(values []int) int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			one := values[i]
			two := values[j]

			if one+two == 2020 {
				return one * two
			}
		}
	}

	return -1
}

func findThreeWithSum(values []int) int {
	for i := 0; i < len(values); i++ {
		for j := i + 1; j < len(values); j++ {
			for k := j + 1; k < len(values); k++ {
				one := values[i]
				two := values[j]
				three := values[k]

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
