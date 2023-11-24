package day6

import (
	"AoC_2022_Go/src/utils"
)

var lines []string = utils.ReadLines("day6/input.txt")

func Day6_part1() int {

	for i := 0; i < len(lines[0])-3; i++ {
		if check_duplicate_chars(lines[0][i : i+4]) {
			return i + 4
		}
	}
	panic("No unique 4 char substring found")
}

func Day6_part2() int {

	for i := 0; i < len(lines[0])-13; i++ {
		if check_duplicate_chars(lines[0][i : i+14]) {
			return i + 14
		}
	}
	panic("No unique 14 char substring found")
}

func check_duplicate_chars(s string) bool {
	d := make(map[byte]int)
	for j := 0; j < len(s); j++ {
		_, ok := d[s[j]]
		if ok {
			return false
		}
		d[s[j]] = 1
	}
	return true
}
