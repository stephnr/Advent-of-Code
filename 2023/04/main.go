package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

// ID => Score
var scratchTable = []float64{}

// Maps Part 2 Scratchcard ID to final total
var deepCounter = map[float64]float64{}

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

		// Track scratchcard win amount
		scratchTable = append(scratchTable, math.Floor(math.Pow(2, matches-1)))
	}

	return result
}

func partTwo(rows []string) float64 {
	var result float64

	for i := 0; i < len(scratchTable)-1; i++ {
		result += countCards(float64(i))
	}

	return result
}

func countCards(idx float64) float64 {
	res := float64(1)
	score := scratchTable[int(idx)]

	if score > 0 {
		for i := float64(idx + 1); i <= float64(idx)+score && i < float64(len(scratchTable)); i++ {
			res += countCards(i)
		}
	}

	return res
}

func parseRows(in string) []string {
	return strings.Split(strings.ReplaceAll(in, "\r\n", "\n"), "\n")
}

func main() {
	println(strconv.Itoa(int(partOne(parseRows(input)))))
	println(fmt.Sprintf("%f\n", partTwo(parseRows(input))))
	println()
}
