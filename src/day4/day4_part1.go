package day4

import (
	"AoC_2022_Go/src/utils"
	"strconv"
	"strings"
)

var lines []string = utils.ReadLines("day4/input.txt")

func Day4_part1() int {

	var total int = 0
	for _, line := range lines {
		if check_overlaps((parse(line))) {
			total++
		}
	}

	return total
}

func parse(line string) [4]int {
	elf1 := strings.Split(line, ",")[0]
	elf2 := strings.Split(line, ",")[1]
	e1a, _ := strconv.Atoi(strings.Split(elf1, "-")[0])
	e1b, _ := strconv.Atoi(strings.Split(elf1, "-")[1])
	e2a, _ := strconv.Atoi(strings.Split(elf2, "-")[0])
	e2b, _ := strconv.Atoi(strings.Split(elf2, "-")[1])
	return [4]int{e1a, e1b, e2a, e2b}
}

func check_overlaps(claims [4]int) bool {
	// ToDo
	return (((claims[0] <= claims[2]) && (claims[1] >= claims[3])) || ((claims[2] <= claims[0]) && (claims[3] >= claims[1])))
}
