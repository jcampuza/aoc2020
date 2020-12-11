package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"math"
	"path/filepath"
	"strings"
)

// Matrix of unit vectors
var positions = [][]int{
	{-1, 1}, {0, 1}, {1, 1},
	{-1, 0}, {1, 0},
	{-1, -1}, {0, -1}, {1, -1},
}

func printGrid(grid []string) {
	for _, line := range grid {
		fmt.Printf("%v\n", line)
	}
}

func areGridsEqual(a []string, b []string) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func countSeats(grid []string) int {
	count := 0
	for _, row := range grid {
		count += strings.Count(row, "#")
	}

	return count
}

func getSurroundingOccupiedSeats(grid []string, pX, pY, sight int) int {
	count := 0

	for _, position := range positions {
		for s := 1; s <= sight; s++ {
			x := position[0] * s
			y := position[1] * s
			isDiagonal := math.Abs(float64(x)) == math.Abs(float64(x))
			isVertical, isHorizontal := y == 0, x == 0

			if !isDiagonal && !isHorizontal && !isVertical {
				continue
			}

			posY := pY + y
			posX := pX + x
			isYInBounds := posY >= 0 && posY < len(grid)
			isXInBounds := posX >= 0 && posX < len(grid[0])

			if !isYInBounds || !isXInBounds {
				continue
			}

			if grid[posY][posX] == 'L' {
				break
			}

			if grid[posY][posX] == '#' {
				count++
				break
			}
		}
	}

	return count
}

func runSimulation(grid []string, tolerance, sight int) []string {
	newGrid := append([]string{}, grid...)

	for y := range grid {
		for x, v := range grid[y] {
			switch v {
			case '.':
				continue
			case 'L':
				if getSurroundingOccupiedSeats(grid, x, y, sight) == 0 {
					newGrid[y] = utils.ReplaceAtIndex(newGrid[y], '#', x)
				}
			case '#':
				if getSurroundingOccupiedSeats(grid, x, y, sight) >= tolerance {
					newGrid[y] = utils.ReplaceAtIndex(newGrid[y], 'L', x)
				}
			}
		}
	}

	return newGrid
}

func partOne(grid []string) int {
	for {
		nextGrid := runSimulation(grid, 4, 1)
		if areGridsEqual(nextGrid, grid) {
			return countSeats(nextGrid)
		}

		grid = nextGrid
	}
}

func partTwo(grid []string) int {
	for {
		nextGrid := runSimulation(grid, 5, utils.Max(len(grid), len(grid[0])))
		if areGridsEqual(nextGrid, grid) {
			return countSeats(nextGrid)
		}

		grid = nextGrid
	}
}

func main() {
	p := filepath.Join("11", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)

	fmt.Println("Part one", partOne(lines))
	fmt.Println("Part two", partTwo(lines))
}
