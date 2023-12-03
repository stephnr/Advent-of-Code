package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var input string

func partOne(matrix [][]string) int {
	var res int

	rowCount := len(matrix)
	rowLength := len(matrix[0])

	symbolRx := regexp.MustCompile(`[^0-9\.]`)

	for rowIdx, row := range matrix {
		println()

		digitGroups := map[int]int{}

		startIdx := -1

		for i, c := range row {
			if unicode.IsDigit([]rune(c)[0]) {
				if startIdx == -1 {
					startIdx = i
					continue
				}
			} else if startIdx != -1 {
				digitGroups[startIdx] = i

				startIdx = -1
			}
		}

	DigitScan:
		for leftIdx, rightIdx := range digitGroups {
			digitGroup := row[leftIdx:rightIdx]

			// Form a scanning zone
			scanZone := matrix[int(math.Max(0, float64(rowIdx-1))):int(math.Min(float64(rowCount), float64(rowIdx+2)))]

			// Scan for a symbol in range
			for _, scanRow := range scanZone {
				scanRowZone := scanRow[int(math.Max(0, float64(leftIdx-1))):int(math.Min(float64(rowLength), float64(rightIdx+1)))]

				symbolGroups := symbolRx.FindStringIndex(strings.Join(scanRowZone, ""))

				if len(symbolGroups) > 0 {
					// Valid
					val, _ := strconv.Atoi(strings.Join(digitGroup, ""))
					res += val

					// Print the scan zone
					fmt.Printf("%v\n", rowIdx)
					for _, scr := range scanZone {
						scrz := scr[int(math.Max(0, float64(leftIdx-1))):int(math.Min(float64(rowLength), float64(rightIdx+1)))]

						fmt.Printf("%v\n", strings.Join(scrz, ""))
					}

					fmt.Println()

					continue DigitScan
				}
			}
		}
	}

	return res
}

func partTwo(matrix [][]string) int {
	var res int
	return res
}

func parseRows(in string) [][]string {
	matrix := [][]string{}

	for _, row := range strings.Split(strings.ReplaceAll(in, "\r\n", "\n"), "\n") {
		matrix = append(matrix, strings.Split(row, ""))
	}

	return matrix
}

func main() {
	println(partOne(parseRows(input)))
	println(partTwo(parseRows(input)))
}
