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

var rowIdxRegex = regexp.MustCompile(`Game\s(\d+)`)
var gameSplitRegex = regexp.MustCompile(`((\d{1,2})\s(r|g|b))`)

func partOne(rows []string) int {
	var result int

	max := map[string]int{"r": 12, "g": 13, "b": 14}

RowStep:
	for _, row := range rows {
		rowID, _ := strconv.Atoi(rowIdxRegex.FindStringSubmatch(row)[1])
		gameResults := gameSplitRegex.FindAllStringSubmatch(row, -1)

		for _, cubeCaptures := range gameResults {
			cubeCount, _ := strconv.Atoi(cubeCaptures[2])
			if cubeCount > max[cubeCaptures[3]] {
				// Invalid game state - skip
				continue RowStep
			}
		}

		result += rowID
	}

	return result
}

func partTwo(rows []string) float64 {
	var result float64

	for _, row := range rows {
		gameResults := gameSplitRegex.FindAllStringSubmatch(row, -1)
		maxPowerSet := map[string]float64{"r": 0, "g": 0, "b": 0}

		for _, cubeCaptures := range gameResults {
			cubeCount, _ := strconv.Atoi(cubeCaptures[2])
			maxPowerSet[cubeCaptures[3]] = math.Max(float64(cubeCount), float64(maxPowerSet[cubeCaptures[3]]))
		}

		result += (maxPowerSet["r"] * maxPowerSet["g"] * maxPowerSet["b"])
	}

	return result
}

func parseRows(in string) []string {
	return strings.Split(strings.ReplaceAll(in, "\r\n", "\n"), "\n")
}

func main() {
	println(partOne(parseRows(input)))
	println(strconv.Itoa(int(partTwo(parseRows(input)))))
}
