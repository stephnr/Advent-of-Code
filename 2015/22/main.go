package main

import (
	"math"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/22

type State struct {
	PlayerHP, Mana, BossHP   int
	Shield, Poison, Recharge int // turns left on each effect
	Spent                    int
}

const (
	MagicMissile = iota
	Drain
	Shield
	Poison
	Recharge
)

var costs = []int{53, 73, 113, 173, 229}

// Ticks active effects at the start of either side's turn; returns player armor
func applyEffects(s *State) int {
	armor := 0

	if s.Shield > 0 {
		armor = 7
		s.Shield--
	}

	if s.Poison > 0 {
		s.BossHP -= 3
		s.Poison--
	}

	if s.Recharge > 0 {
		s.Mana += 101
		s.Recharge--
	}

	return armor
}

// Depth-first search over every spell order, pruning anything already
// costlier than the cheapest win found so far
func fight(s State, bossDamage int, hard bool, best *int) {
	// Player turn. Part Two (hard): lose 1 HP first.
	if hard {
		s.PlayerHP--
		if s.PlayerHP <= 0 {
			return
		}
	}

	applyEffects(&s)
	if s.BossHP <= 0 {
		*best = min(*best, s.Spent)
		return
	}

	for spell, cost := range costs {
		if cost > s.Mana || s.Spent+cost >= *best {
			continue
		}

		n := s
		n.Mana -= cost
		n.Spent += cost

		// Effects can't be recast while still active
		switch spell {
		case MagicMissile:
			n.BossHP -= 4
		case Drain:
			n.BossHP -= 2
			n.PlayerHP += 2
		case Shield:
			if n.Shield > 0 {
				continue
			}
			n.Shield = 6
		case Poison:
			if n.Poison > 0 {
				continue
			}
			n.Poison = 6
		case Recharge:
			if n.Recharge > 0 {
				continue
			}
			n.Recharge = 5
		}

		// Boss turn
		armor := applyEffects(&n)
		if n.BossHP <= 0 {
			*best = min(*best, n.Spent)
			continue
		}

		n.PlayerHP -= max(1, bossDamage-armor)
		if n.PlayerHP > 0 {
			fight(n, bossDamage, hard, best)
		}
	}
}

func minMana(playerHP, mana, bossHP, bossDamage int, hard bool) int {
	best := math.MaxInt
	fight(State{PlayerHP: playerHP, Mana: mana, BossHP: bossHP}, bossDamage, hard, &best)
	return best
}

func run(filepath string, partTwo bool) int {
	// "Hit Points: 71" / "Damage: 10"
	stats := []int{}
	for line := range ReadInput(filepath) {
		_, v, _ := strings.Cut(line, ": ")
		n, _ := strconv.Atoi(v)
		stats = append(stats, n)
	}

	return minMana(50, 500, stats[0], stats[1], partTwo)
}

func main() {
	println(minMana(10, 250, 13, 8, false)) // 226
	println(minMana(10, 250, 14, 8, false)) // 641
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
