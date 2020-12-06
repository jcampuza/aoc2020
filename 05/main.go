package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"sort"
)

func binarySearch(str string, min float64, max float64, lower rune, upper rune) float64 {
	for _, char := range str {
		switch char {
		case lower:
			max = min + math.Floor((max-min)/2.0)
		case upper:
			min = min + math.Ceil((max-min)/2.0)
		}
	}

	return min
}

func getSeatId(str string) float64 {
	row := binarySearch(str[:7], 0., 127., 'F', 'B')
	col := binarySearch(str[7:], 0., 7., 'L', 'R')

	return row*8 + col
}

func getMaxSeatId(strings []string) float64 {
	max := 0.

	for _, string := range strings {
		if m := getSeatId(string); m > max {
			max = m
		}
	}

	return max
}

func getMySeatId(strings []string) float64 {
	// Generate a sorted list of seatIds
	seatIds := []float64{}

	for _, string := range strings {
		seatIds = append(seatIds, getSeatId(string))
	}

	sort.Float64s(seatIds)

	// Search list for an entry where the next item is item+2
	for i := 0; i < len(seatIds)-1; i++ {
		if seatIds[i+1] == seatIds[i]+2 {
			return seatIds[i] + 1
		}
	}

	return 0
}

func main() {
	p := filepath.Join("05", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)

	fmt.Println("Part one:", getMaxSeatId(lines))
	fmt.Println("Part two:", getMySeatId(lines))
}
