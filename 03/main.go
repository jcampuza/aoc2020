package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type TreeGrid = [][]string

type Slope struct {
	x, y int
}

func createTreeGrid(lines []string) TreeGrid {
	grid := make(TreeGrid, len(lines))

	for i, line := range lines {
		grid[i] = make([]string, len(line))
		for j, char := range line {
			grid[i][j] = string(char)
		}
	}

	return grid
}

func findTreesEncountered(lines []string, slope Slope) int {
	grid := createTreeGrid(lines)
	trees := 0
	y, x := 0, 0
	my, mx := slope.y, slope.x

	for {
		if grid[y][x] == "#" {
			trees++
		}

		y, x = my+y, (mx+x)%len(grid[y])
		if y > len(grid)-1 {
			break
		}
	}

	return trees
}

func main() {
	p := filepath.Join("03", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)

	// Part One
	fmt.Println("Part one:", findTreesEncountered(lines, Slope{3, 1}))

	// Part Two
	slopes := []Slope{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	multiple := 1
	for _, slope := range slopes {
		multiple *= findTreesEncountered(lines, slope)
	}

	fmt.Println("part two:", multiple)
}
