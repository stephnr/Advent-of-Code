package main

import (
	"regexp"
	"strconv"
	"strings"
)

func part_two(filepath string) int {
	lines := ReadInput(filepath)
	sum := 0

	for line := range lines {
		left := findDigit(line, true)
		right := findDigit(line, false)

		digit, _ := strconv.Atoi(left + right)
		sum += digit
	}

	return sum
}

func findDigit(line string, left bool) string {
	digits := [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	digitPos := -1
	digit := ""

	rgx := regexp.MustCompile("\\d{1}")
	firstDigits := rgx.FindAllStringIndex(line, -1)

	if len(firstDigits) > 0 {
		if left {
			digitPos = firstDigits[0][0]
			digit = line[firstDigits[0][0]:firstDigits[0][1]]
		} else {
			digitPos = firstDigits[len(firstDigits)-1][0]
			digit = line[firstDigits[len(firstDigits)-1][0]:firstDigits[len(firstDigits)-1][1]]
		}
	}

	for idx, digitName := range digits {
		i := -1

		if left {
			i = strings.Index(line, digitName)

			if i >= 0 && i < digitPos {
				digitPos = i
				digit = strconv.Itoa(idx + 1)
			}
		} else {
			i = strings.LastIndex(line, digitName)

			if i >= 0 && i > digitPos {
				digitPos = i
				digit = strconv.Itoa(idx + 1)
			}
		}
	}

	return digit
}

func main() {
	println(part_two("./inputs/1.txt"))
}
