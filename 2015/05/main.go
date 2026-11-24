package main

import "strings"

// https://adventofcode.com/2015/day/5

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)
	sum := 0

	for line := range lines {
		if partTwo && isNiceTwo(line) {
			sum++
		}

		if !partTwo && isNiceOne(line) {
			sum++
		}
	}

	return sum
}

func isNiceOne(s string) bool {
	for _, bad := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, bad) {
			return false
		}
	}

	vowels, double := 0, false

	for i := 0; i < len(s); i++ {
		if strings.IndexByte("aeiou", s[i]) >= 0 {
			vowels++
		}

		if i > 0 && s[i] == s[i-1] {
			double = true
		}
	}

	return vowels >= 3 && double
}

func isNiceTwo(s string) bool {
	pair, repeat := false, false

	for i := 0; i+2 < len(s); i++ {
		// Same pair appears again later without overlapping this one
		if strings.Contains(s[i+2:], s[i:i+2]) {
			pair = true
		}

		// Letter repeats with exactly one letter between (xyx)
		if s[i] == s[i+2] {
			repeat = true
		}
	}

	return pair && repeat
}

func main() {
	println(run("./inputs/0.txt", false)) // Part One examples
	println(run("./inputs/2.txt", true))  // Part Two examples
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
