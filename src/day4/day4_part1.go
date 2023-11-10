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
		if check_total_overlaps((parse(line))) {
			total++
		}
	}

	return total
}

func Day4_part2() int {

	var total int = 0
	for _, line := range lines {
		if check_any_overlaps((parse(line))) {
			total++
		}
	}

	return total
}

func parse(line string) []int {
	ss := strings.Split(strings.Replace(line, ",", "-", -1), "-")
	return utils.Map(ss, func(s string) int { i, _ := strconv.Atoi(s); return i })
}

func check_total_overlaps(claims []int) bool {
	return (((claims[0] <= claims[2]) && (claims[1] >= claims[3])) || ((claims[2] <= claims[0]) && (claims[3] >= claims[1])))
}

func check_any_overlaps(claims []int) bool {
	return !((claims[1] < claims[2]) || (claims[0] > claims[3]))
}
