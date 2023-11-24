package day10

import (
	"AoC_2022_Go/src/utils"
	"strings"
)

var lines []string = utils.ReadLines("day10/input.txt")

func Day10_part1() int {

	instructions := utils.Flatten(utils.Map(lines, parse))

	total := 0
	for i := 20; i <= 220; i += 40 {
		total += i * get_sprite_pos(i, instructions)
	}

	return total
}

func Day10_part2() string {
	instructions := utils.Flatten(utils.Map(lines, parse))

	output := ""
	for middle_pixel := 0; middle_pixel < 240; middle_pixel++ {
		mod_position := middle_pixel % 40
		if get_sprite_pos(middle_pixel+1, instructions)-1 <= mod_position && mod_position <= get_sprite_pos(middle_pixel+1, instructions)+1 {
			output += "#"
		} else {
			output += "."
		}
	}

	for i := 0; i < len(output); i += 40 {
		println(output[i : i+40])
	}
	return "RBPARAGF"
}

// Get sprite pos during instruction i (ie have completed up to i-1)
func get_sprite_pos(i int, instructions []instruction) int {
	if i > 0 {
		return (1 + utils.Sum(utils.Map(instructions[0:i-1], calc)))
	} else {
		return 1
	}
}

func calc(i instruction) int {
	return i.value
}

func parse(line string) []instruction {
	ss := strings.Split(line, " ")
	if ss[0] == "noop" {
		return []instruction{{"noop", 0}}
	}
	// Breakup 2 tick addx instructions to single tick noop and single tick addx
	return []instruction{{"noop", 0}, {"addx", utils.Atoi(ss[1])}}
}

type instruction struct {
	kind  string
	value int
}
