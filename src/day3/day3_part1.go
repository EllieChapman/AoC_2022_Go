package day3

import (
	"AoC_2022_Go/src/utils"
	"strings"
)

// ToDo questions
// when when use range on a string iterate through runes, but when do string[i] get a byte?

var lines []string = utils.ReadLines("day3/input.txt")

func Day3_part1() int {

	var total int = 0
	for _, line := range lines {
		// total += score(find_duplicate(line))
		total += score(find_variadic(line[0:len(line)/2], line[len(line)/2:]))
	}

	return total
}

func Day3_part2() int {

	var total int = 0
	for i := 0; i < len(lines); i += 3 {
		// total += score(find_badge(lines[i], lines[i+1], lines[i+2]))
		total += score(find_variadic(lines[i], lines[i+1], lines[i+2]))
	}

	return total
}

func find_variadic(ss ...string) byte {
	// EBC what I actually wanted to do
	// empty set
	// for s in ss
	// // create new empty set
	// // for char in s add to new set
	// // intersection of set and origianl set
	// assert lenght of set is one
	// return single char in set

	for i := 0; i < len(ss[0]); i++ {
		b := ss[0][i]
		for _, s := range ss {
			if !strings.Contains(s, string(b)) {
				break
			} else if s == ss[len(ss)-1] {
				return b
			}
		}
	}

	panic("No duplicate found")
}

// 'A', 'Z', 'a', 'z' = 65, 90, 97, 122
// lower case minus 96 for score, upper case minus 38 for score
func score(b byte) int {
	if 96 < b && b < 123 {
		return int(b - 96)
	} else if 64 < b && b < 91 {
		return int(b - 38)
	} else {
		panic("Unexpected character")
	}
}

// func find_duplicate(line string) byte {
// 	for i := 0; i < len(line)/2; i++ {
// 		for j := len(line) / 2; j < len(line); j++ {
// 			if line[i] == line[j] {
// 				return line[i]
// 			}
// 		}
// 	}
// 	panic("No duplicate found")
// }

// func find_badge(e1 string, e2 string, e3 string) byte {
// 	for _, i := range e1 {
// 		for _, j := range e2 {
// 			for _, k := range e3 {
// 				if i == j && j == k {
// 					return byte(i) // ToDo use bytes all way through
// 				}
// 			}
// 		}
// 	}
// 	panic("No badge found")
// }
