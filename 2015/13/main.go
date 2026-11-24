package main

import (
	"fmt"
	"strings"
)

// https://adventofcode.com/2015/day/13

func run(filepath string, partTwo bool) int {
	happy := map[[2]string]int{}
	seen := map[string]bool{}
	people := []string{}

	for line := range ReadInput(filepath) {
		var a, b, dir string
		var n int
		fmt.Sscanf(line, "%s would %s %d happiness units by sitting next to %s", &a, &dir, &n, &b)
		b = strings.TrimSuffix(b, ".")

		if dir == "lose" {
			n = -n
		}
		happy[[2]string{a, b}] = n

		if !seen[a] {
			seen[a] = true
			people = append(people, a)
		}
	}

	// Part Two: add yourself, neutral (0) towards everyone
	if partTwo {
		people = append(people, "Me")
	}

	// The table is circular, so fix the first seat and permute the rest
	best := 0
	permute(people, 1, func(seats []string) {
		total := 0
		for i := range seats {
			a, b := seats[i], seats[(i+1)%len(seats)]
			total += happy[[2]string{a, b}] + happy[[2]string{b, a}]
		}

		best = max(best, total)
	})

	return best
}

// Calls visit with every ordering of a[k:]
func permute(a []string, k int, visit func([]string)) {
	if k == len(a) {
		visit(a)
		return
	}

	for i := k; i < len(a); i++ {
		a[k], a[i] = a[i], a[k]
		permute(a, k+1, visit)
		a[k], a[i] = a[i], a[k]
	}
}

func main() {
	println(run("./inputs/0.txt", false)) // 330
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
