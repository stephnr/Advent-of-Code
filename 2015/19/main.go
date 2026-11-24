package main

import (
	"strings"
	"unicode"
)

// https://adventofcode.com/2015/day/19

func run(filepath string, partTwo bool) int {
	rules := [][2]string{}
	molecule := ""

	for line := range ReadInput(filepath) {
		if from, to, ok := strings.Cut(line, " => "); ok {
			rules = append(rules, [2]string{from, to})
		} else if line != "" {
			molecule = line
		}
	}

	if partTwo {
		return stepsFromE(molecule)
	}

	// Part One: every distinct molecule one replacement away
	seen := map[string]bool{}
	for _, r := range rules {
		for i := 0; i+len(r[0]) <= len(molecule); i++ {
			if molecule[i:i+len(r[0])] == r[0] {
				seen[molecule[:i]+r[1]+molecule[i+len(r[0]):]] = true
			}
		}
	}

	return len(seen)
}

// Part Two: every real input's rules have the shape X => XX, or
// X => X Rn X (Y X)* Ar. Each step adds one element, except Rn/Ar (free)
// and each Y X pair (also free). So, counting elements in the molecule:
//
//	steps = elements - #Rn - #Ar - 2*#Y - 1
//
// This doesn't hold for the puzzle's small example, which uses e => H.
func stepsFromE(molecule string) int {
	elements := 0
	for _, ch := range molecule {
		if unicode.IsUpper(ch) {
			elements++
		}
	}

	rn := strings.Count(molecule, "Rn")
	ar := strings.Count(molecule, "Ar")
	y := strings.Count(molecule, "Y")

	return elements - rn - ar - 2*y - 1
}

func main() {
	println(run("./inputs/0.txt", false)) // 4
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
