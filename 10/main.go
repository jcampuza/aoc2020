package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
)

func partOne(nums []int) int {
	sort.Ints(nums)
	m := map[int]int{}

	jolts := 0
	for _, v := range nums {
		diff := v - jolts
		m[diff]++
		jolts = v
	}

	m[3]++

	return m[1] * m[3]
}

type Adapter struct {
	jolts      int
	paths      int
	compatible []*Adapter
}

func countTotalPaths(adapter *Adapter) int {
	// Only if we haven't already, calculated the number of paths from this adapter
	if adapter.paths == -1 {
		// Last in chain, return 1 compatible path
		if len(adapter.compatible) == 0 {
			return 1
		}

		adapter.paths = 0
		for _, adapterCompat := range adapter.compatible {
			adapter.paths += countTotalPaths(adapterCompat)
		}

	}

	return adapter.paths
}

func partTwo(nums []int) int {
	// Adapter list with "ground" first
	adapters := []Adapter{
		{jolts: 0, paths: -1, compatible: []*Adapter{}},
	}

	for _, num := range nums {
		adapters = append(adapters, Adapter{
			jolts:      num,
			paths:      -1,
			compatible: []*Adapter{},
		})
	}

	// Build list of adapters w/ compatible paths
	for i := range adapters {
		for j := range adapters {
			if (adapters[j].jolts > adapters[i].jolts) && (adapters[j].jolts-adapters[i].jolts) <= 3 {
				adapters[i].compatible = append(adapters[i].compatible, &adapters[j])
			}
		}
	}

	return countTotalPaths(&adapters[0])
}

func main() {
	p := filepath.Join("10", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	nums := utils.GetLinesInt(input)

	fmt.Println("Part one", partOne(nums))
	fmt.Println("Part two", partTwo(nums))
}
