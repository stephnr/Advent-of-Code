package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	grid := ParseInput(ReadFile("./inputs/1"), 1)
	for n := 0; n <= b.N; n++ {
		solve(grid, false)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	grid := ParseInput(ReadFile("./inputs/1"), 5)
	for n := 0; n <= b.N; n++ {
		solve(grid, true)
	}
}
