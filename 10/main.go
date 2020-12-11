package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"
	"time"
)

func partOne(nums []int) int {
	defer utils.TrackTime(time.Now(), "part one")

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

func _countTotalPaths(adapter *Adapter, memory map[*Adapter]int) int {
	if v, ok := memory[adapter]; ok {
		return v
	}

	if len(adapter.compatible) == 0 {
		return 1
	}

	sum := 0
	for _, adapterCompat := range adapter.compatible {
		sum += _countTotalPaths(adapterCompat, memory)
	}

	memory[adapter] = sum

	return sum
}

func countTotalPaths(adapter *Adapter) int {
	memory := map[*Adapter]int{}
	return _countTotalPaths(adapter, memory)
}

func partTwo(nums []int) int {
	defer utils.TrackTime(time.Now(), "part two")

	// Adapter list with "ground" first
	adapters := []Adapter{
		{
			jolts:      0,
			paths:      -1,
			compatible: []*Adapter{},
		},
	}

	for _, num := range nums {
		adapters = append(adapters, Adapter{
			jolts:      num,
			paths:      -1,
			compatible: []*Adapter{},
		})
	}

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
