package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Segments struct {
	digits []string
	output []string
}

func part_one(input []Segments) (count int) {
	for _, segment := range input {
		for _, digit := range segment.output {
			if len(digit) <= 4 || len(digit) == 7 {
				count++
			}
		}
	}

	return count
}

func part_two(input []Segments) (sum int) {
	for _, segment := range input {
		digitMap := make(map[int]string, 10)
		decoder := make(map[string]int, 10)
		fiveSegments := make([]string, 0)
		sixSegments := make([]string, 0)

		for _, digit := range segment.digits {
			sortedDigit := SortString(digit)

			// 1. Find the 1, 4, 7, 8
			switch len(digit) {
			case 2:
				decoder[sortedDigit] = 1
				digitMap[1] = sortedDigit
			case 4:
				decoder[sortedDigit] = 4
				digitMap[4] = sortedDigit
			case 3:
				decoder[sortedDigit] = 7
				digitMap[7] = sortedDigit
			case 7:
				decoder[sortedDigit] = 8
				digitMap[8] = sortedDigit

			// Leftovers
			case 5:
				fiveSegments = append(fiveSegments, sortedDigit)
			case 6:
				sixSegments = append(sixSegments, sortedDigit)
			}
		}

		for _, digit := range fiveSegments {
			// Find 3 : contains 1
			if DiffUnion(digitMap[1], digit) == 0 {
				digitMap[3] = digit
				decoder[digit] = 3
				continue
			}

			// Find 2 & 5
			switch DiffUnion(digitMap[4], digit) {
			case 2:
				digitMap[2] = digit
				decoder[digit] = 2
			case 1:
				digitMap[5] = digit
				decoder[digit] = 5
			}
		}

		for _, digit := range sixSegments {
			// 6 differs by 1 by 1
			if DiffUnion(digitMap[1], digit) == 1 {
				digitMap[6] = digit
				decoder[digit] = 6
				continue
			}

			switch DiffUnion(digitMap[4], digit) {
			case 1:
				digitMap[0] = digit
				decoder[digit] = 0
			case 0:
				digitMap[9] = digit
				decoder[digit] = 9
			}
		}

		// Decode the output
		rawVal := ""

		for _, line := range segment.output {
			sortedLine := SortString(line)
			rawVal += fmt.Sprintf("%d", decoder[sortedLine])
		}

		val, _ := strconv.Atoi(rawVal)
		sum += val
	}

	return sum
}

func SortString(s string) string {
	split := strings.Split(s, "")
	sort.Strings(split)
	return strings.Join(split, "")
}

func DiffUnion(expect string, check string) (diff int) {

	for _, r := range expect {
		if !strings.ContainsRune(check, r) {
			diff++
		}
	}

	return diff
}

func ParseInput(input string) (out Segments) {
	re := regexp.MustCompile(`([a-g]+)`)
	matches := re.FindAllString(input, -1)

	out.digits = matches[0 : len(matches)-4]
	out.output = matches[len(matches)-4:]

	return out
}

func main() {
	lines := ReadFile("./inputs/1")
	segments := make([]Segments, len(lines))

	for i, line := range lines {
		segments[i] = ParseInput(line)
	}

	println(part_one(segments))
	println(part_two(segments))
}
