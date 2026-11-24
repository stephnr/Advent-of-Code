package main

import "strings"

// https://adventofcode.com/2015/day/11

func run(filepath string, partTwo bool) string {
	pw := ""
	for line := range ReadInput(filepath) {
		pw += strings.TrimSpace(line)
	}

	// Part Two is simply the next valid password after Part One's
	pw = next(pw)
	if partTwo {
		pw = next(pw)
	}

	return pw
}

func next(s string) string {
	b := []byte(s)

	for {
		increment(b)
		if valid(b) {
			return string(b)
		}
	}
}

// Increments like a base-26 number: "xz" -> "ya"
func increment(b []byte) {
	for i := len(b) - 1; i >= 0; i-- {
		if b[i] == 'z' {
			b[i] = 'a'
			continue
		}

		b[i]++

		// Skip straight past forbidden letters
		if b[i] == 'i' || b[i] == 'o' || b[i] == 'l' {
			b[i]++
		}

		return
	}
}

func valid(b []byte) bool {
	straight := false
	pairs := map[byte]bool{}

	for i := 0; i < len(b); i++ {
		if b[i] == 'i' || b[i] == 'o' || b[i] == 'l' {
			return false
		}

		if i+2 < len(b) && b[i+1] == b[i]+1 && b[i+2] == b[i]+2 {
			straight = true
		}
	}

	// Count different, non-overlapping pairs like "aa" and "bb"
	for i := 0; i+1 < len(b); i++ {
		if b[i] == b[i+1] {
			pairs[b[i]] = true
			i++
		}
	}

	return straight && len(pairs) >= 2
}

func main() {
	println(next("abcdefgh")) // abcdffaa
	println(next("ghijklmn")) // ghjaabcc
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
