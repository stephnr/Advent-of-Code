package main

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func calculateDistance(filepath string) int {
	lines := ReadInput(filepath)
	l, r, sum := []int{}, []int{}, float64(0)

	for line := range lines {
		out := strings.Split(line, "   ")
		li, _ := strconv.Atoi(out[0])
		ri, _ := strconv.Atoi(out[1])
		l = append(l, li)
		r = append(r, ri)
	}

	sort.Ints(l)
	sort.Ints(r)

	for i := 0; i < len(l); i++ {
		sum += math.Max(float64(l[i]), float64(r[i])) - math.Min(float64(l[i]), float64(r[i]))
	}

	return int(sum)
}

func calculateSimilarity(filepath string) int {
	lines := ReadInput(filepath)
	l, r, sum := map[int]int{}, map[int]int{}, 0

	// Read in as a map
	for line := range lines {
		out := strings.Split(line, "   ")
		li, _ := strconv.Atoi(out[0])
		ri, _ := strconv.Atoi(out[1])
		l[li]++
		r[ri]++
	}

	// Iterate left and multiplay by O(1) on set
	for i := range l {
		sum += (i * r[i])
	}

	return sum
}

func main() {
	println(calculateDistance("./inputs/1.txt"))
	println(calculateSimilarity("./inputs/1.txt"))
}
