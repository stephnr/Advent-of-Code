package main

import "strconv"

// https://adventofcode.com/2015/day/17

func combos(filepath string, target int, partTwo bool) int {
	sizes := []int{}
	for line := range ReadInput(filepath) {
		n, _ := strconv.Atoi(line)
		sizes = append(sizes, n)
	}

	// ways[k] = number of k-container combinations that hold exactly target
	ways := map[int]int{}

	// Every subset of containers as a bitmask (20 containers = ~1M subsets)
	for mask := 0; mask < 1<<len(sizes); mask++ {
		sum, k := 0, 0
		for i, s := range sizes {
			if mask&(1<<i) != 0 {
				sum += s
				k++
			}
		}

		if sum == target {
			ways[k]++
		}
	}

	if !partTwo {
		total := 0
		for _, w := range ways {
			total += w
		}
		return total
	}

	// Part Two: only combinations using the fewest containers
	for k := 0; k <= len(sizes); k++ {
		if ways[k] > 0 {
			return ways[k]
		}
	}

	return 0
}

func run(filepath string, partTwo bool) int {
	return combos(filepath, 150, partTwo)
}

func main() {
	println(combos("./inputs/0.txt", 25, false)) // 4
	println(combos("./inputs/0.txt", 25, true))  // 3
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
