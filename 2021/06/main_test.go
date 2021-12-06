package main

import "testing"

func helpInput(b *testing.B) map[int]int {
	b.Helper()

	fishes := ParseInput(ReadFile("./inputs/1")[0])

	// Reduce fish duplicates
	fishMap := make(map[int]int)

	for _, fish := range fishes {
		fishMap[fish]++
	}

	return fishMap
}

func BenchmarkPartOne(b *testing.B) {
	fishes := helpInput(b)

	for n := 0; n <= b.N; n++ {
		part_one(fishes, 80)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	fishes := helpInput(b)

	for n := 0; n <= b.N; n++ {
		part_one(fishes, 256)
	}
}
