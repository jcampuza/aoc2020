package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile("(?s)Player 1:\n(?P<one>.*)\n\nPlayer 2:\n(?P<two>.*)")

func getTotal(list []int) int {
	acc := 0
	for i, value := range list {
		acc += value * (len(list) - i)
	}

	return acc
}

func createConfiguration(p1 []int, p2 []int) string {
	str := "p1:"
	for _, value := range p1 {
		str += strconv.Itoa(value)
	}

	str += "p2:"
	for _, value := range p2 {
		str += strconv.Itoa(value)
	}

	return str
}

func partOne(p1 []int, p2 []int) int {
	for len(p1) > 0 && len(p2) > 0 {
		p1Card, p2Card := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]

		if p1Card > p2Card {
			p1 = append(p1, p1Card, p2Card)
		} else {
			p2 = append(p2, p2Card, p1Card)
		}
	}

	if len(p1) > len(p2) {
		return getTotal(p1)
	}

	return getTotal(p2)
}

func partTwo(p1 []int, p2 []int) (int, int) {
	m := map[string]bool{}

	for len(p1) > 0 && len(p2) > 0 {
		// Winner is player one if config already exists
		configuration := createConfiguration(p1, p2)

		if m[configuration] {
			return 1, getTotal(p1)
		}

		m[configuration] = true

		// Play the game
		p1Card, p2Card := p1[0], p2[0]
		p1, p2 = p1[1:], p2[1:]
		winner := 1

		if len(p1) >= p1Card && len(p2) >= p2Card {
			// Play subgame and winner is whoever won subgame
			copy1 := append([]int{}, p1[:p1Card]...)
			copy2 := append([]int{}, p2[:p2Card]...)
			winner, _ = partTwo(copy1, copy2)
		} else {
			// winner is higher values card
			if p1Card > p2Card {
				winner = 1
			} else {
				winner = 2
			}
		}

		if winner == 1 {
			p1 = append(p1, p1Card, p2Card)
		} else {
			p2 = append(p2, p2Card, p1Card)
		}
	}

	if len(p1) > len(p2) {
		return 1, getTotal(p1)
	}

	return 2, getTotal(p2)
}

func main() {
	p := filepath.Join("22", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	res := re.FindStringSubmatch(string(input))
	m := map[string]string{}
	for i, name := range re.SubexpNames() {
		if i != 0 && name != "" {
			m[name] = res[i]
		}
	}

	p1 := []int{}
	for _, value := range strings.Split(m["one"], "\n") {
		v, _ := strconv.Atoi(value)
		p1 = append(p1, v)
	}

	p2 := []int{}
	for _, value := range strings.Split(m["two"], "\n") {
		v, _ := strconv.Atoi(value)
		p2 = append(p2, v)
	}

	fmt.Println("Part one", partOne(p1, p2))

	_, total := partTwo(p1, p2)
	fmt.Println("Part two", total)
}
