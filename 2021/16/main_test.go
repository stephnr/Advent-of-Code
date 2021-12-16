package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		part_one()
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		// part_two()
	}
}
