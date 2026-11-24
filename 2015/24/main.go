package main

import (
	"slices"
	"strconv"
)

// https://adventofcode.com/2015/day/24

func run(filepath string, partTwo bool) int {
	weights := []int{}
	for line := range ReadInput(filepath) {
		n, _ := strconv.Atoi(line)
		weights = append(weights, n)
	}

	groups := 3
	if partTwo {
		groups = 4 // Part Two adds the trunk
	}

	return bestEntanglement(weights, groups)
}

// Finds the smallest first group (then lowest product) whose weight is
// total/groups and whose leftover packages still split evenly
func bestEntanglement(weights []int, groups int) int {
	total := 0
	for _, w := range weights {
		total += w
	}
	target := total / groups

	for size := 1; size <= len(weights); size++ {
		// All first groups of this size, lowest entanglement first
		candidates := [][]int{}
		subsets(weights, size, target, func(pick []int) {
			candidates = append(candidates, slices.Clone(pick))
		})

		slices.SortFunc(candidates, func(a, b []int) int {
			return product(a) - product(b)
		})

		for _, c := range candidates {
			if canSplit(without(weights, c), groups-1, target) {
				return product(c)
			}
		}
	}

	return -1
}

// Calls visit with every size-k subset of weights summing to target
func subsets(weights []int, k, target int, visit func([]int)) {
	pick := []int{}

	var walk func(i, sum int)
	walk = func(i, sum int) {
		if len(pick) == k {
			if sum == target {
				visit(pick)
			}
			return
		}

		for j := i; j < len(weights); j++ {
			if sum+weights[j] > target {
				continue
			}

			pick = append(pick, weights[j])
			walk(j+1, sum+weights[j])
			pick = pick[:len(pick)-1]
		}
	}

	walk(0, 0)
}

// Reports whether weights can be split into n groups of target each
func canSplit(weights []int, n, target int) bool {
	if n == 1 {
		return true // total is n*target, so the rest is exactly one group
	}

	for k := 1; k <= len(weights); k++ {
		found := false
		subsets(weights, k, target, func(pick []int) {
			if !found && canSplit(without(weights, pick), n-1, target) {
				found = true
			}
		})

		if found {
			return true
		}
	}

	return false
}

// Returns weights minus one copy of each value in remove
func without(weights, remove []int) []int {
	rest := slices.Clone(weights)
	for _, r := range remove {
		i := slices.Index(rest, r)
		rest = slices.Delete(rest, i, i+1)
	}
	return rest
}

func product(a []int) int {
	p := 1
	for _, n := range a {
		p *= n
	}
	return p
}

func main() {
	println(run("./inputs/0.txt", false)) // 99
	println(run("./inputs/0.txt", true))  // 44
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
