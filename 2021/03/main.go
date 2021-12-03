package main

import (
	"strconv"
	"strings"
)

func CalculatePartTwo(bitCol int, rows []string, useGamma bool) string {
	remaining := make([]string, 0)

	gamma, epsilon := calculateGammaAndEpsilon(rows)
	mask := gamma

	if !useGamma {
		mask = epsilon
	}

	if len(rows) == 1 {
		return rows[0]
	}

	for _, row := range rows {
	RowCheck:
		for col, bit := range row {
			check := []rune(mask)[bitCol]

			if col != bitCol {
				continue
			}

			// Check case for even parity
			if check == '-' {
				if useGamma {
					check = '1'
				} else {
					check = '0'
				}
			}

			if bit == check {
				remaining = append(remaining, row)
				break RowCheck
			}
		}
	}

	return CalculatePartTwo(bitCol+1, remaining, useGamma)
}

func calculateGammaAndEpsilon(lines []string) (gammaStr string, epsilonStr string) {
	rotatedBits := make([]string, len(lines[0]))

	// Rotate the input where columns = rows
	for _, line := range lines {
		for j, bit := range line {
			rotatedBits[j] = rotatedBits[j] + string(bit)
		}
	}

	// Perform a simple parity count
	for _, bits := range rotatedBits {
		count0 := strings.Count(bits, "0")
		count1 := strings.Count(bits, "1")

		if count0 == count1 {
			gammaStr += "-"
			epsilonStr += "-"
		} else if count1 > count0 {
			gammaStr += "1"
			epsilonStr += "0"
		} else {
			gammaStr += "0"
			epsilonStr += "1"
		}
	}

	return
}

func part_one() (string, string) {
	lines := ReadFile("./inputs/1")
	return calculateGammaAndEpsilon(lines)
}

func part_two() int64 {
	lines := ReadFile("./inputs/1")

	oxygenRatingStr := CalculatePartTwo(0, lines, true)
	co2RatingStr := CalculatePartTwo(0, lines, false)

	// Convert binary values
	oxygenVal, _ := strconv.ParseInt(oxygenRatingStr, 2, 32)
	co2Val, _ := strconv.ParseInt(co2RatingStr, 2, 32)

	return oxygenVal * co2Val
}

func main() {
	gamma, epsilon := part_one()

	// Convert binary values
	gammaVal, _ := strconv.ParseInt(gamma, 2, 32)
	epsilonVal, _ := strconv.ParseInt(epsilon, 2, 32)

	println(gammaVal * epsilonVal)

	println(part_two())
}
