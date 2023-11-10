package day2

import (
	"AoC_2022_Go/src/utils"
)

var lines []string = utils.ReadLines("day2/input.txt")

func Day2_part1() int {
	var total int = 0

	for _, line := range lines {
		total += score(interpret_move(line[0]), interpret_move(line[2]))
		total += interpret_move(line[2]) + 1
	}

	return total
}

func Day2_part2() int {
	var total int = 0

	for _, line := range lines {
		total += calc_move(interpret_move(line[0]), interpret_move(line[2]))
		total += interpret_move(line[2]) * 3
	}

	return total
}

// result receives  0=lose, 1=draw, 2=win
// result: need to convert to 0=draw, 1=win, 2=lose, so add 2 then mod 3
func calc_move(move int, result int) int {
	// slice represents {R, P, S} as their scores
	moves := []int{1, 2, 3}
	index := (move + ((result + 2) % 3)) % 3
	return moves[index]
}

func interpret_move(move byte) int {
	switch {
	case move == 'A' || move == 'X':
		return 0 // rock or lose
	case move == 'B' || move == 'Y':
		return 1 // paper or draw
	case move == 'C' || move == 'Z':
		return 2 // scissors or win
	default:
		panic("Unexpected move")
	}
}

func score(opponant int, me int) int {
	switch {
	case opponant == me:
		return 3
	case (3+me-opponant)%3 == 1:
		return 6
	default:
		return 0
	}
}
