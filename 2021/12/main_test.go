package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	paths := ParseInput(ReadFile("./inputs/1"))

	for n := 0; n <= b.N; n++ {
		solve(paths, false)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	paths := ParseInput(ReadFile("./inputs/1"))

	for n := 0; n <= b.N; n++ {
		solve(paths, true)
	}
}
