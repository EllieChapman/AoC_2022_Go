package day7

import (
	"AoC_2022_Go/src/utils"
	"strings"
)

var lines []string = utils.ReadLines("day7/input.txt")

func Day7_part1() int {
	top, _ := make_dir(lines[2:])

	return find_small_dirs(top)
}

func Day7_part2() int {
	top, _ := make_dir(lines[2:])
	space_needed := top.recursive_size - 40000000

	return find_smallest_ok(top, space_needed, top.recursive_size)
}

func find_smallest_ok(d dir, needed int, current_best int) int {
	if d.recursive_size < current_best && d.recursive_size >= needed {
		current_best = d.recursive_size
	}
	for _, subdir := range d.dirs {
		tmp := find_smallest_ok(subdir, needed, current_best)
		if tmp < current_best {
			current_best = tmp
		}
	}
	return current_best
}

func find_small_dirs(d dir) int {
	tmp := 0
	if d.recursive_size < 100000 {
		tmp = d.recursive_size
	}
	return utils.Sum(utils.Map(d.dirs, find_small_dirs)) + tmp
}

func make_dir(lines []string) (dir, []string) {
	// takes in slice starting with stuff in the dict, ie not including the cd into it line or the ls line
	// create blank dir, ie size 0 and empty dirs
	// iterate through lines
	// if starts with cd <dir> (ie cd and then not ..), recusrive call and append returned dir to dirs. make sure to remove the cd and following ls line before passing in
	// if it starts with cd .., return dir and remaining lines (same if run out of lines)
	// if starts with dir, ignore
	// if starts with ls should panic, never should hit
	// otherwise is a number, so add to size

	current := dir{}
	for len(lines) > 0 {
		if strings.HasPrefix(lines[0], "$ cd") {
			if strings.HasPrefix(lines[0], "$ cd ..") {
				return current, lines[1:]
			} else {
				d, remaining := make_dir(lines[2:])
				current.dirs = append(current.dirs, d)
				current.recursive_size += d.recursive_size
				lines = remaining
			}
		} else if strings.HasPrefix(lines[0], "$ ls") {
			panic("Should not hit ls")
		} else if strings.HasPrefix(lines[0], "dir") {
			lines = lines[1:]
		} else {
			current.recursive_size += utils.Atoi(strings.Split(lines[0], " ")[0])
			lines = lines[1:]
		}
	}
	return current, []string{}
}

type dir struct {
	recursive_size int
	dirs           []dir
}
