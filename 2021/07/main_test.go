package main

import "testing"

func GetCrabs(b *testing.B) []int {
	b.Helper()
	return ParseCrabs(ReadFile("./inputs/1")[0])
}

func BenchmarkPartOne(b *testing.B) {
	crabs := GetCrabs(b)
	for n := 0; n <= b.N; n++ {
		part_one(crabs, true)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	crabs := GetCrabs(b)
	for n := 0; n <= b.N; n++ {
		part_one(crabs, false)
	}
}
