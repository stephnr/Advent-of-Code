package main

import "testing"

func GetGrid(b *testing.B) [][]Octo {
	b.Helper()
	return ParseInput(ReadFile("./inputs/1"))
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		grid := GetGrid(b)
		part_one(grid, 100)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		grid := GetGrid(b)
		part_two(grid)
	}
}
