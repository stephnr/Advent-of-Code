package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	input := ParseInput(ReadFile("./inputs/1"))

	for n := 0; n <= b.N; n++ {
		part_one(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	input := ParseInput(ReadFile("./inputs/1"))

	for n := 0; n <= b.N; n++ {
		part_two(input)
	}
}
