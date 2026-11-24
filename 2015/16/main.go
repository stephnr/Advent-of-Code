package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/16

// What the MFCSAM detected on the gift
var ticker = map[string]int{
	"children":    3,
	"cats":        7,
	"samoyeds":    2,
	"pomeranians": 3,
	"akitas":      0,
	"vizslas":     0,
	"goldfish":    5,
	"trees":       3,
	"cars":        2,
	"perfumes":    1,
}

func run(filepath string, partTwo bool) int {
	for line := range ReadInput(filepath) {
		// "Sue 1: cars: 9, akitas: 3, goldfish: 0"
		head, rest, _ := strings.Cut(line, ": ")

		var sue int
		fmt.Sscanf(head, "Sue %d", &sue)

		match := true
		for _, kv := range strings.Split(rest, ", ") {
			k, v, _ := strings.Cut(kv, ": ")
			n, _ := strconv.Atoi(v)

			if !matches(k, n, partTwo) {
				match = false
				break
			}
		}

		if match {
			return sue
		}
	}

	return -1
}

// Part Two: cats/trees readings are lower bounds, pomeranians/goldfish upper bounds
func matches(k string, v int, partTwo bool) bool {
	want := ticker[k]

	if partTwo {
		switch k {
		case "cats", "trees":
			return v > want
		case "pomeranians", "goldfish":
			return v < want
		}
	}

	return v == want
}

func main() {
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
