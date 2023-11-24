package day5

import (
	"AoC_2022_Go/src/utils"
	"slices"
	"strconv"
	"strings"
)

var lines []string = utils.ReadLines("day5/input.txt")

func Day5_part1() string {
	positions_unparsed, instructions_unparsed := separate_two_inputs(lines)

	var positions [][]string = parse_positions(positions_unparsed)
	var instructions []instruction = flatten(parse_instructions(instructions_unparsed))

	final_pos := apply(instructions, positions)

	return collect_answer(final_pos)
}

func Day5_part2() string {
	positions_unparsed, instructions_unparsed := separate_two_inputs(lines)

	var positions [][]string = parse_positions(positions_unparsed)
	var instructions []instruction = parse_instructions(instructions_unparsed)

	final_pos := apply(instructions, positions)

	return collect_answer(final_pos)
}

func apply(instructions []instruction, positions [][]string) [][]string {
	i := instructions[0]
	for _, block := range positions[i.src][len(positions[i.src])-i.count:] {
		positions[i.dst] = append(positions[i.dst], block)
	}
	positions[i.src] = positions[i.src][0 : len(positions[i.src])-i.count]

	if len(instructions) != 1 {
		return apply(instructions[1:], positions)
	} else {
		return positions
	}
}

type instruction struct {
	count, src, dst int
}

func flatten(instructions []instruction) []instruction {
	new_instructions := []instruction{}
	for _, i := range instructions {
		for n := 0; n < i.count; n++ {
			new_instructions = append(new_instructions, instruction{count: 1, src: i.src, dst: i.dst})
		}
	}
	return new_instructions
}

func collect_answer(positions [][]string) string {
	var ans string
	for _, position := range positions {
		ans += position[len(position)-1]
	}
	return ans
}

func separate_two_inputs(lines []string) ([]string, []string) {
	for i, line := range lines {
		if line == "" {
			return lines[0 : i-1], lines[i+1:]
		}
	}
	panic("No blank line found")
}

func parse_instructions(instructions_unparsed []string) []instruction {
	var instructions []instruction
	for _, line := range instructions_unparsed {
		// EBC this is super dumb and relies on Atoi returning 0 on error
		ss := utils.Map(strings.Split(line, " "), func(s string) int { i, _ := strconv.Atoi(s); return i })
		instructions = append(instructions, instruction{count: ss[1], src: ss[3] - 1, dst: ss[5] - 1})
	}
	return instructions
}

func parse_positions(positions_unparsed []string) [][]string {
	var positions [][]string

	for _, line := range positions_unparsed {
		column := 0
		for i := 1; i < len(line); i += 4 {
			if len(positions) < column+1 {
				positions = append(positions, []string{})
			}
			if line[i] != ' ' {
				positions[column] = append(positions[column], string(line[i]))
			}
			column++
		}
	}

	return utils.Map(positions, func(ss []string) []string { slices.Reverse(ss); return ss })
}

// func apply(instructions []instruction, positions [][]string) [][]string {
// 	i := &instructions[0]
// 	positions[i.dst] = append(positions[i.dst], positions[i.src][len(positions[i.src])-1])
// 	positions[i.src] = positions[i.src][0 : len(positions[i.src])-1]
// 	i.count--
// 	if i.count == 0 {
// 		instructions = instructions[1:]
// 	}
// 	if len(instructions) != 0 {
// 		return apply(instructions, positions)
// 	} else {
// 		return positions
// 	}
// }
