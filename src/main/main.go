package main

import (
	"AoC_2022_Go/src/day1"
	"AoC_2022_Go/src/day10"
	"AoC_2022_Go/src/day11"
	"AoC_2022_Go/src/day2"
	"AoC_2022_Go/src/day3"
	"AoC_2022_Go/src/day4"
	"AoC_2022_Go/src/day5"
	"AoC_2022_Go/src/day6"
	"AoC_2022_Go/src/day7"
	"AoC_2022_Go/src/day8"
	"AoC_2022_Go/src/day9"
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
	fmt.Println("Day 4 Part 2")
	utils.Check_answer(day4.Day4_part2(), 845)

	fmt.Println("Day 5 Part 1")
	utils.Check_answer(day5.Day5_part1(), "SBPQRSCDF")
	fmt.Println("Day 5 Part 2")
	utils.Check_answer(day5.Day5_part2(), "RGLVRCQSB")

	fmt.Println("Day 6 Part 1")
	utils.Check_answer(day6.Day6_part1(), 1134)
	fmt.Println("Day 6 Part 2")
	utils.Check_answer(day6.Day6_part2(), 2263)

	fmt.Println("Day 7 Part 1")
	utils.Check_answer(day7.Day7_part1(), 1334506)
	fmt.Println("Day 7 Part 2")
	utils.Check_answer(day7.Day7_part2(), 7421137)

	fmt.Println("Day 8 Part 1")
	utils.Check_answer(day8.Day8_part1(), 1672)
	fmt.Println("Day 8 Part 2")
	utils.Check_answer(day8.Day8_part2(), 327180)

	fmt.Println("Day 9 Part 1")
	utils.Check_answer(day9.Day9_part1(), 6357)
	fmt.Println("Day 9 Part 2")
	utils.Check_answer(day9.Day9_part2(), 2627)

	fmt.Println("Day 10 Part 1")
	utils.Check_answer(day10.Day10_part1(), 12740)
	fmt.Println("Day 10 Part 2")
	utils.Check_answer(day10.Day10_part2(), "RBPARAGF")

	fmt.Println("Day 11 Part 1")
	utils.Check_answer(day11.Day11_part1(), 66802)

	utils.Check_all()
}
