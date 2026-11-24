package main

import "fmt"

// https://adventofcode.com/2015/day/9

func run(filepath string, partTwo bool) int {
	dist := map[[2]string]int{}
	seen := map[string]bool{}
	cities := []string{}

	for line := range ReadInput(filepath) {
		var a, b string
		var d int
		fmt.Sscanf(line, "%s to %s = %d", &a, &b, &d)

		dist[[2]string{a, b}] = d
		dist[[2]string{b, a}] = d

		for _, c := range []string{a, b} {
			if !seen[c] {
				seen[c] = true
				cities = append(cities, c)
			}
		}
	}

	// Brute force every route: shortest for Part One, longest for Part Two
	best := -1
	permute(cities, 0, func(route []string) {
		total := 0
		for i := 1; i < len(route); i++ {
			total += dist[[2]string{route[i-1], route[i]}]
		}

		if best == -1 || (!partTwo && total < best) || (partTwo && total > best) {
			best = total
		}
	})

	return best
}

// Calls visit with every ordering of a (Heap-style swap permutation)
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
	println(run("./inputs/0.txt", false)) // 605
	println(run("./inputs/0.txt", true))  // 982
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
