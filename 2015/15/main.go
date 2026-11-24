package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/15

// Ingredient properties: capacity, durability, flavor, texture, calories
type Ingredient [5]int

func run(filepath string, partTwo bool) int {
	ings := []Ingredient{}

	for line := range ReadInput(filepath) {
		// "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8"
		_, props, _ := strings.Cut(line, ": ")

		var ing Ingredient
		for i, p := range strings.Split(props, ", ") {
			_, v, _ := strings.Cut(p, " ")
			ing[i], _ = strconv.Atoi(v)
		}

		ings = append(ings, ing)
	}

	best := 0
	amounts := make([]int, len(ings))

	// Try every way to split 100 teaspoons between the ingredients
	var mix func(i, left int)
	mix = func(i, left int) {
		if i == len(ings)-1 {
			amounts[i] = left
			best = max(best, score(ings, amounts, partTwo))
			return
		}

		for n := 0; n <= left; n++ {
			amounts[i] = n
			mix(i+1, left-n)
		}
	}
	mix(0, 100)

	return best
}

func score(ings []Ingredient, amounts []int, partTwo bool) int {
	var totals Ingredient
	for i, ing := range ings {
		for p := range ing {
			totals[p] += amounts[i] * ing[p]
		}
	}

	// Part Two: only cookies with exactly 500 calories count
	if partTwo && totals[4] != 500 {
		return 0
	}

	product := 1
	for _, t := range totals[:4] {
		product *= max(t, 0)
	}

	return product
}

func main() {
	println(run("./inputs/0.txt", false)) // 62842880
	println(run("./inputs/0.txt", true))  // 57600000
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
