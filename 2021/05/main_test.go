package main

import "testing"

func GetGeysers(b *testing.B) []Geyser {
	b.Helper()
	return ParseInput(ReadFile("./inputs/1"))
}

func BenchmarkPartOne(b *testing.B) {
	geysers := GetGeysers(b)

	for n := 0; n <= b.N; n++ {
		part_one(geysers, false)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	geysers := GetGeysers(b)

	for n := 0; n <= b.N; n++ {
		part_one(geysers, true)
	}
}
