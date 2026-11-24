package main

import (
	"slices"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/2

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)
	sum := 0

	for line := range lines {
		dims := []int{}
		for _, s := range strings.Split(line, "x") {
			n, _ := strconv.Atoi(s)
			dims = append(dims, n)
		}

		// Sorted so dims[0] and dims[1] are the two smallest sides
		slices.Sort(dims)
		l, w, h := dims[0], dims[1], dims[2]

		if partTwo {
			// Ribbon: smallest perimeter + volume for the bow
			sum += 2*(l+w) + l*w*h
		} else {
			// Paper: surface area + area of smallest side
			sum += 2*(l*w+w*h+h*l) + l*w
		}
	}

	return sum
}

func main() {
	println(run("./inputs/0.txt", false))
	println(run("./inputs/0.txt", true))
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
