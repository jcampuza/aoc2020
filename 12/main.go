package main

import (
	"aoc/utils"
	"container/ring"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"
)

type position struct {
	x, y int
	dir  string
}

func createCardinalRing(start string) *ring.Ring {
	entries := []string{"N", "E", "S", "W"}
	ring := ring.New(len(entries))

	for _, entry := range entries {
		ring.Value = entry
		ring = ring.Next()
	}

	// Rotate ourselves to our wanted starting direction
	for ring.Value != start {
		ring = ring.Next()
	}

	return ring
}

func executeCardinalDirection(pos position, dir string, value int) position {
	switch dir {
	case "N":
		pos.y += value
	case "E":
		pos.x += value
	case "S":
		pos.y -= value
	case "W":
		pos.x -= value
	}
	return pos
}

// Rotating the position is done by rotating through a circular list (ring) {N,E,S,W}
func rotatePosition(pos position, dir string, value int) position {
	compass := createCardinalRing(pos.dir)

	switch dir {
	case "L":
		pos.dir = compass.Move(-value / 90).Value.(string)
	case "R":
		pos.dir = compass.Move(value / 90).Value.(string)
	}

	return pos
}

// Rotating the waypoint is treated as a series of 90deg vector rotations
func rotateWaypoint(waypoint position, dir string, value int) position {
	updated := waypoint

	for v := value; v > 0; v -= 90 {
		switch dir {
		case "L":
			updated.x, updated.y = -waypoint.y, waypoint.x
		case "R":
			updated.x, updated.y = waypoint.y, -waypoint.x
		}

		waypoint = updated
	}

	return waypoint
}

func chaseWaypoint(pos position, waypoint position, times int) position {
	for i := 0; i < times; i++ {
		pos = executeCardinalDirection(pos, "E", waypoint.x)
		pos = executeCardinalDirection(pos, "N", waypoint.y)
	}

	return pos
}

func partOne(instructions []string) int {
	pos := position{0, 0, "E"}

	for _, instruction := range instructions {
		action := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])

		switch action {
		case "L", "R":
			pos = rotatePosition(pos, action, value)
		case "N", "E", "S", "W":
			pos = executeCardinalDirection(pos, action, value)
		case "F":
			pos = executeCardinalDirection(pos, pos.dir, value)
		}
	}

	return utils.AbsI(pos.x) + utils.AbsI(pos.y)
}

func partTwo(instructions []string) int {
	pos := position{0, 0, "E"}
	waypoint := position{10, 1, ""}

	for _, instruction := range instructions {
		action := string(instruction[0])
		value, _ := strconv.Atoi(instruction[1:])

		switch action {
		case "L", "R":
			waypoint = rotateWaypoint(waypoint, action, value)
		case "N", "E", "S", "W":
			waypoint = executeCardinalDirection(waypoint, action, value)
		case "F":
			pos = chaseWaypoint(pos, waypoint, value)
		}
	}

	return utils.AbsI(pos.x) + utils.AbsI(pos.y)
}

func main() {
	p := filepath.Join("12", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)
	fmt.Println("Part one", partOne(lines))
	fmt.Println("Part two", partTwo(lines))
}
