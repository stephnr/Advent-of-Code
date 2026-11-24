package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/6

func run(filepath string, partTwo bool) int {
	grid := make([]int, 1000*1000)

	for line := range ReadInput(filepath) {
		// "turn on 0,0 through 9,9" / "toggle 0,0 through 9,9"
		f := strings.Fields(line)
		action := f[len(f)-4] // "on", "off" or "toggle"
		x1, y1 := parsePoint(f[len(f)-3])
		x2, y2 := parsePoint(f[len(f)-1])

		for x := x1; x <= x2; x++ {
			for y := y1; y <= y2; y++ {
				i := x*1000 + y
				grid[i] = apply(action, grid[i], partTwo)
			}
		}
	}

	sum := 0
	for _, v := range grid {
		sum += v
	}

	return sum
}

// Part One: lights are on/off. Part Two: lights have a brightness.
func apply(action string, v int, partTwo bool) int {
	switch action {
	case "on":
		if partTwo {
			return v + 1
		}
		return 1
	case "off":
		if partTwo {
			return max(v-1, 0)
		}
		return 0
	default: // toggle
		if partTwo {
			return v + 2
		}
		return 1 - v
	}
}

func parsePoint(s string) (int, int) {
	a, b, _ := strings.Cut(s, ",")
	x, _ := strconv.Atoi(a)
	y, _ := strconv.Atoi(b)
	return x, y
}

func main() {
	println(run("./inputs/0.txt", false)) // 998996
	println(run("./inputs/0.txt", true))  // 1001996
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
