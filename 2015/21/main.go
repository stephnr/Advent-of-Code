package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/21

type Item struct{ Cost, Damage, Armor int }

type Fighter struct{ HP, Damage, Armor int }

var weapons = []Item{{8, 4, 0}, {10, 5, 0}, {25, 6, 0}, {40, 7, 0}, {74, 8, 0}}

// The first entry is "no armor"
var armors = []Item{{0, 0, 0}, {13, 0, 1}, {31, 0, 2}, {53, 0, 3}, {75, 0, 4}, {102, 0, 5}}

var rings = []Item{{25, 1, 0}, {50, 2, 0}, {100, 3, 0}, {20, 0, 1}, {40, 0, 2}, {80, 0, 3}}

// The player attacks first. Each side needs ceil(HP / hit) turns to win.
func wins(player, boss Fighter) bool {
	pHit := max(1, player.Damage-boss.Armor)
	bHit := max(1, boss.Damage-player.Armor)

	return (boss.HP+pHit-1)/pHit <= (player.HP+bHit-1)/bHit
}

func run(filepath string, partTwo bool) int {
	// "Hit Points: 104" / "Damage: 8" / "Armor: 1"
	stats := []int{}
	for line := range ReadInput(filepath) {
		_, v, _ := strings.Cut(line, ": ")
		n, _ := strconv.Atoi(v)
		stats = append(stats, n)
	}
	boss := Fighter{stats[0], stats[1], stats[2]}

	// Zero, one or two different rings
	ringSets := []Item{{0, 0, 0}}
	for i, a := range rings {
		ringSets = append(ringSets, a)
		for _, b := range rings[i+1:] {
			ringSets = append(ringSets, Item{a.Cost + b.Cost, a.Damage + b.Damage, a.Armor + b.Armor})
		}
	}

	// Part One: cheapest win. Part Two: most expensive loss.
	best := -1
	for _, w := range weapons {
		for _, a := range armors {
			for _, r := range ringSets {
				cost := w.Cost + a.Cost + r.Cost
				won := wins(Fighter{100, w.Damage + a.Damage + r.Damage, w.Armor + a.Armor + r.Armor}, boss)

				if !partTwo && won && (best == -1 || cost < best) {
					best = cost
				}

				if partTwo && !won && cost > best {
					best = cost
				}
			}
		}
	}

	return best
}

func main() {
	println(wins(Fighter{8, 5, 5}, Fighter{12, 7, 2})) // true
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
