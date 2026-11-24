package main

import "fmt"

// https://adventofcode.com/2015/day/14

type Reindeer struct {
	Speed int // km/s while flying
	Fly   int // seconds of flying per cycle
	Rest  int // seconds of resting per cycle
}

// Distance flown after t seconds
func (r Reindeer) Distance(t int) int {
	cycle := r.Fly + r.Rest
	flying := (t/cycle)*r.Fly + min(t%cycle, r.Fly)
	return r.Speed * flying
}

func race(filepath string, seconds int, partTwo bool) int {
	deer := []Reindeer{}

	for line := range ReadInput(filepath) {
		var name string
		var r Reindeer
		fmt.Sscanf(line, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&name, &r.Speed, &r.Fly, &r.Rest)
		deer = append(deer, r)
	}

	best := 0

	if !partTwo {
		for _, r := range deer {
			best = max(best, r.Distance(seconds))
		}
		return best
	}

	// Part Two: every second, each reindeer in the lead scores a point
	points := make([]int, len(deer))

	for t := 1; t <= seconds; t++ {
		lead := 0
		for _, r := range deer {
			lead = max(lead, r.Distance(t))
		}

		for i, r := range deer {
			if r.Distance(t) == lead {
				points[i]++
			}
		}
	}

	for _, p := range points {
		best = max(best, p)
	}

	return best
}

func run(filepath string, partTwo bool) int {
	return race(filepath, 2503, partTwo)
}

func main() {
	println(race("./inputs/0.txt", 1000, false)) // 1120
	println(race("./inputs/0.txt", 1000, true))  // 689
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
