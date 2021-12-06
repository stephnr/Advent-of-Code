package main

import (
	"strconv"
	"strings"
)

func part_one(fishes map[int]int, days int) int {
	processedFish := make(map[int]int)

	if days == 0 {
		count := 0

		for _, val := range fishes {
			count += val
		}
		return count
	}

	// Process a day
	for fish, val := range fishes {
		fish--

		if fish < 0 {
			// Gen new fish
			fish = 6
			processedFish[8] += val
		}

		processedFish[fish] += val
	}

	return part_one(processedFish, days-1)
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
