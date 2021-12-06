package main

import (
	"fmt"
	"strconv"
	"strings"
)

var (
	// Days -> Fish -> Count
	mnemonicTable = make(map[string]int)
)

func count_fish_growth(fish, days, count int) int {
	fish--
	days--

	if fish < 0 && days >= 0 {
		fish = 6

		key := fmt.Sprintf("%d:%d", days, 8)

		if mnemonicTable[key] != 0 {
			count += mnemonicTable[key]
		} else {
			growth := count_fish_growth(8, days, 1)
			count += growth
			mnemonicTable[key] = growth
		}

	}

	if days < 0 {
		return count
	}

	return count_fish_growth(fish, days, count)
}

func part_one(fishes map[int]int, days int) int {
	var count int

	// Process a fish
	for fish, val := range fishes {
		count += (count_fish_growth(fish, days, 1) * val)
	}

	return count
}

func ParseInput(line string) (out []int) {
	for _, raw := range strings.Split(line, ",") {
		val, _ := strconv.Atoi(raw)
		out = append(out, val)
	}

	return out
}

func main() {
	fishes := ParseInput(ReadFile("./inputs/1")[0])

	// Reduce fish duplicates
	fishMap := make(map[int]int)

	for _, fish := range fishes {
		fishMap[fish]++
	}

	println(part_one(fishMap, 80))
	println(part_one(fishMap, 256))
}
