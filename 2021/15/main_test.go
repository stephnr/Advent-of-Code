package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	grid := ParseInput(ReadFile("./inputs/1"))
	for n := 0; n <= b.N; n++ {
		part_one(grid)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		// part_two()
	}
}
