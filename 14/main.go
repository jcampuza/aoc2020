package main

import (
	"aoc/utils"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

type AddressDecoder = func(string, string) string
type BitmaskDecoder = func(rune, rune) rune

var memoryRegex = regexp.MustCompile("mem\\[(\\d+)\\] = (\\d+)")

func decoderV1(bitmask, bit rune) rune {
	if bitmask == 'X' {
		return bit
	}

	return bitmask
}

func decoderV2(bitmask, bit rune) rune {
	if bitmask == 'X' || bitmask == '1' {
		return bitmask
	}

	return bit
}

func createMaskDecoder(decoder BitmaskDecoder) AddressDecoder {
	return func(mask string, value string) string {
		for len(value) < 36 {
			value = "0" + value
		}

		out := []rune(value)
		for i, v := range mask {
			out[i] = decoder(v, out[i])
		}

		return string(out)
	}
}

func partOne(instructions []string) int64 {
	memory := map[int]string{}
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	decoder := createMaskDecoder(decoderV1)

	for _, instruction := range instructions {
		switch {
		case strings.HasPrefix(instruction, "mask"):
			mask = strings.Split(instruction, " = ")[1]

		case strings.HasPrefix(instruction, "mem"):
			result := memoryRegex.FindAllStringSubmatch(instruction, -1)
			register, valueStr := utils.ParseInt(result[0][1]), utils.ParseInt(result[0][2])
			valueInt := strconv.FormatInt(int64(valueStr), 2)
			memory[register] = decoder(mask, valueInt)
		}
	}

	count := int64(0)
	for _, v := range memory {
		v, _ := strconv.ParseInt(v, 2, 64)
		count += v
	}

	return count
}

func writeToAddresses(addr string, memory map[int64]int64, value int64) map[int64]int64 {
	idx := strings.Index(addr, "X")

	if idx == -1 {
		v, _ := strconv.ParseInt(addr, 2, 64)
		memory[v] = value
		return memory
	}

	memory = writeToAddresses(utils.ReplaceAtIndex(addr, '0', idx), memory, value)
	memory = writeToAddresses(utils.ReplaceAtIndex(addr, '1', idx), memory, value)
	return memory
}

func partTwo(instructions []string) int64 {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	memory := map[int64]int64{}
	decoder := createMaskDecoder(decoderV2)

	for _, instruction := range instructions {
		switch {
		case strings.HasPrefix(instruction, "mask"):
			mask = strings.Split(instruction, " = ")[1]

		case strings.HasPrefix(instruction, "mem"):
			result := memoryRegex.FindAllStringSubmatch(instruction, -1)
			addrInt, valueInt := utils.ParseInt(result[0][1]), int64(utils.ParseInt(result[0][2]))
			addr := decoder(mask, strconv.FormatInt(int64(addrInt), 2))
			memory = writeToAddresses(addr, memory, valueInt)
		}
	}

	count := int64(0)
	for _, v := range memory {
		count += v
	}

	return count
}

func main() {
	p := filepath.Join("14", "input.txt")
	input, err := ioutil.ReadFile(p)
	if err != nil {
		panic(err)
	}

	lines := utils.GetLines(input)
	fmt.Println("Part one", partOne(lines))
	fmt.Println("Part two", partTwo(lines))
}
