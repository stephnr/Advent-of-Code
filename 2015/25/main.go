package main

import (
	"regexp"
	"strconv"
)

// https://adventofcode.com/2015/day/25

func run(filepath string) int {
	// "... Enter the code at row 2981, column 3075."
	nums := []int{}
	for line := range ReadInput(filepath) {
		for _, m := range regexp.MustCompile(`\d+`).FindAllString(line, -1) {
			n, _ := strconv.Atoi(m)
			nums = append(nums, n)
		}
	}

	return code(nums[0], nums[1])
}

// Codes fill the grid diagonal by diagonal (bottom-left to top-right),
// each one = previous * 252533 % 33554393, starting from 20151125
func code(row, col int) int {
	diag := row + col - 1
	n := diag*(diag-1)/2 + col // 1-based position in fill order

	c := 20151125
	for i := 1; i < n; i++ {
		c = c * 252533 % 33554393
	}

	return c
}

func main() {
	println(code(4, 2)) // 32451966
	println(code(6, 6)) // 27995004
	println(run("./inputs/1.txt"))
}
