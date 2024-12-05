package main

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func partOne(filepath string) int {
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

func partTwo(filepath string) int {
	lines := ReadInput(filepath)
	rules := map[int][]int{}
	sum := 0

Line:
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
						// Reorder the bag chain
						newSS := []int{}

						for _, sEl := range ss {
							// Place on end
							sEli, _ := strconv.Atoi(sEl)
							if !sliceContains(newSS, rules[sEli]) {
								newSS = append(newSS, sEli)
								continue
							}

							// Check for best position from back
							for j := len(newSS); j >= 0; j-- {
								if j == 0 {
									newSS = append([]int{sEli}, newSS...)
									break
								} else if !sliceContains(newSS[0:j], rules[sEli]) {
									newSS = slices.Insert(newSS, j, sEli)
									break
								}
							}
						}

						// Count the middle of the corrected sort string
						sum += newSS[int(math.Ceil(float64(len(newSS)/2)))]
						continue Line
					}
				}
			}
		}
	}

	return sum
}

// Checks if a contains anything from b
func sliceContains(a []int, b []int) bool {
	for _, el := range b {
		if slices.Contains(a, el) {
			return true
		}
	}

	return false
}

func main() {
	println(partOne("./inputs/0.txt"))
	println(partTwo("./inputs/1.txt"))
}
