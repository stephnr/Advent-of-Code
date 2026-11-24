package main

import (
	"encoding/json"
	"log"
)

// https://adventofcode.com/2015/day/12

func run(filepath string, partTwo bool) int {
	sum := 0

	for line := range ReadInput(filepath) {
		var doc any
		if err := json.Unmarshal([]byte(line), &doc); err != nil {
			log.Fatal(err)
		}

		sum += total(doc, partTwo)
	}

	return sum
}

// Sums every number in a decoded JSON value. Part Two skips any object
// (and everything inside it) that has a "red" value.
func total(v any, partTwo bool) int {
	sum := 0

	switch t := v.(type) {
	case float64:
		return int(t)
	case []any:
		for _, x := range t {
			sum += total(x, partTwo)
		}
	case map[string]any:
		if partTwo {
			for _, x := range t {
				if x == "red" {
					return 0
				}
			}
		}

		for _, x := range t {
			sum += total(x, partTwo)
		}
	}

	return sum
}

func main() {
	println(run("./inputs/0.txt", false)) // 33
	println(run("./inputs/0.txt", true))  // 16
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
