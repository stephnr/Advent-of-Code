package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	lines := ReadFile("./inputs/1")

	for n := 0; n <= b.N; n++ {
		part_one(lines, false)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	lines := ReadFile("./inputs/1")

	for n := 0; n <= b.N; n++ {
		part_one(lines, true)
	}
}
