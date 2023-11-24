package day9

import (
	"AoC_2022_Go/src/utils"
	"strings"
)

var lines []string = utils.ReadLines("day9/input.txt")

func Day9_part1() int {

	instructions := utils.Map(lines, parse)

	start := pos{0, 0}
	m := map[pos]int{start: 1}
	knots := []pos{start, start}

	// This works because maps are effectively passed as reference
	run(instructions, knots, m)
	return len(m)
}

func Day9_part2() int {

	instructions := utils.Map(lines, parse)

	start := pos{0, 0}
	knots := []pos{start, start, start, start, start, start, start, start, start, start}
	m := map[pos]int{start: 1}

	run(instructions, knots, m)
	return len(m)
}

func parse(line string) instruction {
	return instruction{strings.Split(line, " ")[0], utils.Atoi(strings.Split(line, " ")[1])}
}

func run(is []instruction, knots []pos, m map[pos]int) {
	knots[0] = move(is[0].dir, knots[0])

	if is[0].count > 1 {
		is[0].count--
	} else {
		is = is[1:]
	}

	for i := 0; i < len(knots)-1; i++ {
		head := knots[i]
		tail := knots[i+1]

		if !adjacent(head, tail) {
			knots[i+1] = calc_tail_pos(head, tail)
		}
	}

	// Create or update entry
	m[knots[len(knots)-1]] = 1

	if len(is) > 0 {
		run(is, knots, m)
		return
	}
	return
}

func calc_tail_pos(head pos, tail pos) pos {
	return pos{resolve(head.x, tail.x), resolve(head.y, tail.y)}
}

func resolve(head int, tail int) int {
	// if not the same, make tail one closer
	if head > tail {
		return tail + 1
	} else if head < tail {
		return tail - 1
	} else {
		return tail
	}
}

func adjacent(p1 pos, p2 pos) bool {
	if utils.AbsDiffInt(p1.x, p2.x) <= 1 && utils.AbsDiffInt(p1.y, p2.y) <= 1 {
		return true
	}
	return false
}

func move(dir string, p pos) pos {
	switch dir {
	case "U":
		return pos{p.x, p.y + 1}
	case "D":
		return pos{p.x, p.y - 1}
	case "R":
		return pos{p.x + 1, p.y}
	case "L":
		return pos{p.x - 1, p.y}
	default:
		panic("Invalid direction")
	}
}

type instruction struct {
	dir   string
	count int
}

type pos struct {
	x, y int
}
