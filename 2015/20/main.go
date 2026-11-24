package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/20

func run(filepath string, partTwo bool) int {
	target := 0
	for line := range ReadInput(filepath) {
		target, _ = strconv.Atoi(strings.TrimSpace(line))
	}

	return firstHouse(target, partTwo)
}

// Sieve: elf e drops presents at houses e, 2e, 3e, ...
// Part One: 10 presents per house, forever. Part Two: 11 presents, 50 houses each.
func firstHouse(target int, partTwo bool) int {
	per, limit := 10, 0 // limit 0 = no limit
	if partTwo {
		per, limit = 11, 50
	}

	// House n gets at least per*n presents from elf n alone,
	// so the answer is never above target/per
	n := target / per
	houses := make([]int, n+1)

	for elf := 1; elf <= n; elf++ {
		visits := 0
		for h := elf; h <= n && (limit == 0 || visits < limit); h += elf {
			houses[h] += elf * per
			visits++
		}
	}

	for h := 1; h <= n; h++ {
		if houses[h] >= target {
			return h
		}
	}

	return -1
}

func main() {
	println(firstHouse(70, false))  // 4
	println(firstHouse(120, false)) // 6
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
