package main

import (
	"strconv"
	"strings"
)

func countSafeReports(filepath string) int {
	lines := ReadInput(filepath)
	sum := 0

	for line := range lines {
		inc := isInc(line)
		dec := isDec(line)

		if inc || dec {
			sum += 1
		}
	}

	return sum
}

func isDec(ss string) bool {
	s := strings.Split(ss, " ")

	for i := 1; i < len(s); i++ {
		l, _ := strconv.Atoi(s[i-1])
		r, _ := strconv.Atoi(s[i])
		diff := (r - l)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func isInc(ss string) bool {
	s := strings.Split(ss, " ")

	for i := 1; i < len(s); i++ {
		l, _ := strconv.Atoi(s[i-1])
		r, _ := strconv.Atoi(s[i])
		diff := (l - r)
		if diff < 1 || diff > 3 {
			return false
		}
	}

	return true
}

func main() {
	println(countSafeReports("./inputs/1.txt"))
}
