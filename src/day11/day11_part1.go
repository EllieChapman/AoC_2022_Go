package day11

import (
	"AoC_2022_Go/src/utils"
	"strings"
)

var lines []string = utils.ReadLines("day11/input.txt")

func Day11_part1() int {

	// monkeys := parse(lines)
	// fmt.Println(monkeys)

	return 0
}

func parse(lines []string) []monkey {
	monkeys := []monkey{}
	for _, line := range lines {
		ss := strings.Split(line, " ")
		if ss[0] == "Monkey" {
			monkeys = append(monkeys, monkey{utils.Atoi(ss[1][0:1]), []int{}, "", 0, 0, 0, 0})
		} else if ss[0] == "Starting" {
			monkeys[len(monkeys)-1].items = utils.Map(ss[2:], func(s string) int { return utils.Atoi(s[0:2]) })
		} else if ss[0] == "Operation" {
			monkeys[len(monkeys)-1].optype = ss[4]
			if ss[5] == "old" {
				monkeys[len(monkeys)-1].opval = -1
			} else {
				monkeys[len(monkeys)-1].opval = utils.Atoi(ss[5])
			}
		} else if ss[0] == "Test" {
			monkeys[len(monkeys)-1].div_val = utils.Atoi(ss[3])
		} else if ss[1] == "true" {
			monkeys[len(monkeys)-1].if_true = utils.Atoi(ss[5])
		} else {
			monkeys[len(monkeys)-1].if_false = utils.Atoi(ss[5])
		}
	}
	return monkeys
}

type monkey struct {
	id       int
	items    []int
	optype   string
	opval    int
	div_val  int
	if_true  int
	if_false int
}
