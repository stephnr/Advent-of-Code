package main

import (
	"math"
)

func part_one(lines []string, part_two bool) (score int) {
	starts := make([]rune, 0)
	mismatches := make([]rune, 0)

	for _, line := range lines {
	CHAR_LOOP:
		for _, char := range line {
			switch char {
			case '[':
				fallthrough
			case '(':
				fallthrough
			case '{':
				fallthrough
			case '<':
				starts = append(starts, char)
				continue
			// Closer found
			default:
				if len(starts) == 0 {
					mismatches = append(mismatches, char)
					break CHAR_LOOP
				}

				if math.Abs(float64(char-starts[len(starts)-1])) > 2 {
					mismatches = append(mismatches, char)
					break CHAR_LOOP
				} else {
					if len(starts) >= 2 {
						starts = starts[0 : len(starts)-1]
					} else {
						starts = make([]rune, 0)
					}
				}
			}
		}
	}

	// Count score
	for _, mismatch := range mismatches {
		switch mismatch {
		case ')':
			score += 3
		case ']':
			score += 57
		case '}':
			score += 1197
		case '>':
			score += 25137
		}
	}

	return score
}

func main() {
	lines := ReadFile("./inputs/1")
	println(part_one(lines))
	// println(part_two())
}
