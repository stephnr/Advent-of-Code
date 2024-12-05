package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func run(filepath string) int {
	lines := ReadInput(filepath)
	rules := map[int][]int{}
	sum := 0

Check:
	for line := range lines {
		if strings.Contains(line, "|") {
			ss := strings.Split(line, "|")
			a, _ := strconv.Atoi(ss[0])
			b, _ := strconv.Atoi(ss[1])
			rules[a] = append(rules[a], b)
		}

		// Line check
		if strings.Contains(line, ",") {
			ss := strings.Split(line, ",")

			for i := 1; i < len(ss); i++ {
				p := ss[i]
				pp, _ := strconv.Atoi(string(p))

				for _, el := range ss[:i] {
					ele, _ := strconv.Atoi(string(el))
					if slices.Contains(rules[pp], ele) {
						continue Check
					}
				}
			}

			el, _ := strconv.Atoi(ss[int(math.Ceil(float64(len(ss)/2)))])

			sum += el
		}
	}

	return sum
}

func main() {
	println(run("./inputs/1.txt"))
}
