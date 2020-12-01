package utils

import "strconv"

func SliceAtoi(str []string) []int {
	tmp := []int{}

	for _, v := range str {
		val, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}

		tmp = append(tmp, val)
	}

	return tmp
}

func FilterString(str []string, compare func(string) bool) []string {
	tmp := []string{}

	for _, v := range str {
		if compare(v) {
			tmp = append(tmp, v)
		}
	}

	return tmp
}
