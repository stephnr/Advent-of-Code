package main

import "testing"

func GetSegments(b *testing.B) []Segments {
	b.Helper()

	lines := ReadFile("./inputs/1")
	segments := make([]Segments, len(lines))

	for i, line := range lines {
		segments[i] = ParseInput(line)
	}

	return segments
}

func BenchmarkPartOne(b *testing.B) {
	segments := GetSegments(b)

	for n := 0; n <= b.N; n++ {
		part_one(segments)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	segments := GetSegments(b)

	for n := 0; n <= b.N; n++ {
		part_two(segments)
	}
}
