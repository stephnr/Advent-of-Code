package main

import "testing"

func BenchmarkPartOne(b *testing.B) {
	polymer := ParseInput(ReadFile("./inputs/1"))
	chainDuplicates = make(map[string]map[string]int)

	for n := 0; n <= b.N; n++ {
		part_one(polymer, 10)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	polymer := ParseInput(ReadFile("./inputs/1"))
	chainDuplicates = make(map[string]map[string]int)

	for n := 0; n <= b.N; n++ {
		part_one(polymer, 40)
	}
}
