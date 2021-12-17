package main

import "testing"

func BenchmarkTest(b *testing.B) {
	for n := 0; n <= b.N; n++ {
		solve(zone)
	}
}
