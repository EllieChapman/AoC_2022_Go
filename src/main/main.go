package main

import (
	"AoC_2022_Go/src/day1"
	"AoC_2022_Go/src/day2"
	"AoC_2022_Go/src/day3"
	"AoC_2022_Go/src/day4"
	"AoC_2022_Go/src/utils"
	"fmt"
)

func main() {
	fmt.Println("Starting tests")

	fmt.Println("Day 1 Part 1")
	utils.Check_answer(day1.Day1_part1(), 71471)
	fmt.Println("Day 1 Part 2")
	utils.Check_answer(day1.Day1_part2(), 211189)

	fmt.Println("Day 2 Part 1")
	utils.Check_answer(day2.Day2_part1(), 13809)
	fmt.Println("Day 2 Part 2")
	utils.Check_answer(day2.Day2_part2(), 12316)

	fmt.Println("Day 3 Part 1")
	utils.Check_answer(day3.Day3_part1(), 7863)
	fmt.Println("Day 3 Part 2")
	utils.Check_answer(day3.Day3_part2(), 2488)

	fmt.Println("Day 4 Part 1")
	utils.Check_answer(day4.Day4_part1(), 536)

	utils.Check_all()
}
