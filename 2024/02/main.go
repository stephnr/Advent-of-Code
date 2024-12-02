package main

import (
	"strconv"
	"strings"
)

func countSafeReports(filepath string, dampener bool) int {
	lines := ReadInput(filepath)
	sum := 0

	for line := range lines {
		l := strings.Split(line, " ")
		inc := isInc(l)
		dec := isDec(l)

		if inc || dec {
			sum += 1
			continue
		}

		// Check with slice iterations
		if dampener && checkWithDampener(line) {
			sum += 1
		}
	}

	return sum
}

func checkWithDampener(ss string) bool {
	s := strings.Split(ss, " ")

	for i := 0; i < len(s); i++ {
		if i == 0 {
			if isInc(s[1:]) || isDec(s[1:]) {
				return true
			}
		} else if i == len(s) {
			if isInc(s[:len(s)-1]) || isDec(s[:len(s)-1]) {
				return true
			}
		} else {
			arr := []string{}
			arr = append(arr, s[0:i]...)
			arr = append(arr, s[i+1:]...)
			if isInc(arr) || isDec(arr) {
				return true
			}
		}
	}

	return false
}

func isDec(s []string) bool {
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

func isInc(s []string) bool {
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
	println(countSafeReports("./inputs/1.txt", false))
	println(countSafeReports("./inputs/1.txt", true))
}
