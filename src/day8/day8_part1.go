package day8

import (
	"AoC_2022_Go/src/utils"
	"strings"
)

var lines []string = utils.ReadLines("day8/input.txt")

func Day8_part1() int {

	heights := utils.Map(lines, func(s string) []int { return utils.Map(strings.Split(s, ""), utils.Atoi) })

	total := 0
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			total += check_visible(x, y, heights)
		}
	}
	return total
}

func Day8_part2() int {

	heights := utils.Map(lines, func(s string) []int { return utils.Map(strings.Split(s, ""), utils.Atoi) })

	total := 0
	for x := 0; x < len(lines[0]); x++ {
		for y := 0; y < len(lines); y++ {
			tmp := check_score(x, y, heights)
			if tmp > total {
				total = tmp
			}
		}
	}
	return total
}

// func check_visible(x int, y int, heights [][]int) int {
// 	if check_column(x, y, heights) || check_row(x, heights[y]) {
// 		return 1
// 	}
// 	return 0
// }

// func check_row(x int, row []int) bool {
// 	for i := x + 1; i < len(row); i++ {
// 		if row[i] >= row[x] {
// 			// blocked on right
// 			return false
// 		}
// 	}
// 	// visible
// 	return true
// }

// func check_column(x int, y int, heights [][]int) bool {
// 	for i := y + 1; i < len(heights); i++ {
// 		if heights[i][x] >= heights[y][x] {
// 			// blocked below, check up
// 			for j := 0; j < y; j++ {
// 				if heights[j][x] >= heights[y][x] {
// 					// also blocked up
// 					return false
// 				}
// 			}
// 		}
// 	}
// 	// visible
// 	return true
// }

func check_visible(x int, y int, heights [][]int) int {
	if check_column_down(x, y, heights, false) ||
		check_row_to_right(x, heights[y], false) ||
		check_column_down(x, len(heights)-(y+1), utils.Reverse(heights), true) ||
		check_row_to_right(len(heights[0])-(x+1), utils.Reverse(heights[y]), true) {
		return 1
	}
	return 0
}

func check_score(x int, y int, heights [][]int) int {
	total := 1
	total *= observable_right(x, heights[y]) * observable_right(len(heights[0])-(x+1), utils.Reverse(heights[y]))
	utils.Reverse(heights[y])
	total *= observable_down(x, y, heights) * observable_down(x, len(heights)-(y+1), utils.Reverse(heights))
	utils.Reverse(heights)
	return total
}

func check_row_to_right(x int, row []int, reverse bool) bool {
	if reverse {
		defer utils.Reverse(row)
	}
	for i := x + 1; i < len(row); i++ {
		if row[i] >= row[x] {
			// blocked on right
			return false
		}
	}
	// visible
	return true
}

func check_column_down(x int, y int, heights [][]int, reverse bool) bool {
	if reverse {
		defer utils.Reverse(heights)
	}
	for i := y + 1; i < len(heights); i++ {
		if heights[i][x] >= heights[y][x] {
			// blocked below
			return false
		}
	}
	// visible
	return true
}

func observable_right(x int, row []int) int {
	total := 0
	for i := x + 1; i < len(row); i++ {
		total++
		if row[i] >= row[x] {
			return total
		}
	}
	return total
}

func observable_down(x int, y int, heights [][]int) int {
	total := 0
	for i := y + 1; i < len(heights); i++ {
		total++
		if heights[i][x] >= heights[y][x] {
			return total
		}
	}
	return total
}
