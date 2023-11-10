package day1

import (
	"AoC_2022_Go/src/utils"
	"slices"
	"strconv"
)

var lines []string = utils.ReadLines("day1/input.txt")

func Day1_part1() int {

	totals := get_totals(lines)

	return slices.Max(totals)
}

func Day1_part2() int {

	totals := get_totals(lines)

	// ToDo, is there a way to chain these functions?
	slices.Sort(totals)
	slices.Reverse(totals)

	return totals[0] + totals[1] + totals[2]
}

func get_totals(lines []string) []int {
	var totals []int

	var individual_total int = 0
	for _, line := range lines {
		if line != "" {
			i, _ := strconv.Atoi(line)
			individual_total += i
		} else {
			totals = append(totals, individual_total)
			individual_total = 0
		}
	}
	totals = append(totals, individual_total)

	return totals
}
