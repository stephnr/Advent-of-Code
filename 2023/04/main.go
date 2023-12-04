package main

import (
	_ "embed"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

func partOne(rows []string) float64 {
	result := float64(0)

	scratchRgx := regexp.MustCompile("\\:(.*)\\s\\|\\s(.*)")

	for _, row := range rows {
		numberGroups := scratchRgx.FindAllStringSubmatch(row, -1)
		winningNumbers := map[int]bool{}

		matches := float64(0)

		// Parse winning numbers (numberGroups[1])
		for _, num := range strings.Split(strings.TrimSpace(numberGroups[0][1]), " ") {
			val, _ := strconv.Atoi(num)
			winningNumbers[val] = true
		}

		// Parse card numbers (numberGroups[2])
		for _, num := range strings.Split(strings.TrimSpace(numberGroups[0][2]), " ") {
			if len(num) == 0 {
				continue
			}

			val, _ := strconv.Atoi(num)
			if winningNumbers[val] {
				// Prevent double counting
				winningNumbers[val] = false

				matches++
			}
		}

		if matches > 0 {
			result += math.Pow(2, matches-1)
		}
	}

	return result
}

func partTwo(rows []string) int {
	var result int
	return result
}

func parseRows(in string) []string {
	return strings.Split(strings.ReplaceAll(in, "\r\n", "\n"), "\n")
}

func main() {
	println(strconv.Itoa(int(partOne(parseRows(input)))))
	println(partTwo(parseRows(input)))
}
