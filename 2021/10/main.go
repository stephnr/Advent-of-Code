package main

import (
	"math"
	"sort"
	"strings"
)

var (
	leftAssignScorecard = "([{<"
)

func part_one(lines []string, part_two bool) (result uint64) {
	scores := make([]uint64, 0)

	for _, line := range lines {
		var mismatch rune

		starts := make([]rune, 0)
		score := uint64(0)

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
					mismatch = char
					break CHAR_LOOP
				}

				if math.Abs(float64(char-starts[len(starts)-1])) > 2 {
					mismatch = char
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

		// Part Two
		if part_two && len(starts) > 0 && mismatch == 0 {
			// Score the missing incomplete parts
			for i := len(starts) - 1; i >= 0; i-- {
				// println(string(starts[i]), score)
				// println(int(strings.IndexRune(leftAssignScorecard, starts[i])) + 1)
				score = uint64(score*5 + (uint64(strings.IndexRune(leftAssignScorecard, starts[i])) + 1))
			}

			scores = append(scores, score)
		} else {
			// Part One - score mismatches
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

			result += score
		}
	}

	if part_two {
		// Find median of scores
		sort.Slice(scores, func(i, j int) bool { return scores[i] < scores[j] })
		return scores[len(scores)/2]
	}

	// Part One: Just return the sum of mismatches
	return result
}

func main() {
	lines := ReadFile("./inputs/1")
	println(part_one(lines, false))
	println(part_one(lines, true))
}
