package main

import (
	"fmt"
	"math"
	"regexp"
)

type Polymer struct {
	Pairs        []string
	Rules        map[string]string
	ElementCount map[string]int
}

var (
	// Root Pair => ElementCount after 40
	chainDuplicates map[string]map[string]int = make(map[string]map[string]int)
)

func GetChainKey(pair string, stepsLeft int) string {
	return fmt.Sprintf("%s:%d", pair, stepsLeft)
}

func AddToChain(pair string, stepsLeft int, countTable map[string]int) {
	if stepsLeft < 0 {
		return
	}

	chainDuplicates[GetChainKey(pair, stepsLeft)] = countTable
}

func part_one(polymer Polymer, steps int) int64 {
	// Run each polymer pair N times
	for _, pair := range polymer.Pairs {
		var checkCount map[string]int

		if chainDuplicates[GetChainKey(pair, steps)] == nil {
			checkCount = RunStep(Polymer{
				Pairs:        []string{pair},
				Rules:        polymer.Rules,
				ElementCount: make(map[string]int),
			}, 0, steps).ElementCount

			AddToChain(pair, steps, checkCount)
		} else {
			checkCount = chainDuplicates[GetChainKey(pair, steps)]
		}

		for k, v := range checkCount {
			polymer.ElementCount[k] += v
		}
	}

	var (
		min float64 = math.MaxInt64
		max float64 = 1
	)

	for _, count := range polymer.ElementCount {
		min = math.Min(min, float64(count))
		max = math.Max(max, float64(count))
	}

	return int64(max - min)
}

func RunStep(polymer Polymer, stepCount, steps int) Polymer {
	if stepCount == steps {
		return polymer
	}

	// Iterate the pairs and resolve them
	for _, pair := range polymer.Pairs {
		if polymer.Rules[pair] != "" {
			var (
				newPair    string
				checkCount map[string]int
			)

			polymer.ElementCount[polymer.Rules[pair]]++

			// Left side pair
			newPair = fmt.Sprintf("%s%s", pair[0:1], polymer.Rules[pair])

			if chainDuplicates[GetChainKey(newPair, stepCount-1)] == nil {
				checkCount = RunStep(Polymer{
					Pairs:        []string{newPair},
					Rules:        polymer.Rules,
					ElementCount: make(map[string]int),
				}, stepCount+1, steps).ElementCount

				AddToChain(newPair, stepCount-1, checkCount)
			} else {
				checkCount = chainDuplicates[GetChainKey(newPair, stepCount-1)]
			}

			for k, v := range checkCount {
				polymer.ElementCount[k] += v
			}

			// Right side pair
			newPair = fmt.Sprintf("%s%s", polymer.Rules[pair], pair[1:])

			if chainDuplicates[GetChainKey(newPair, stepCount-1)] == nil {
				checkCount = RunStep(Polymer{
					Pairs:        []string{newPair},
					Rules:        polymer.Rules,
					ElementCount: make(map[string]int),
				}, stepCount+1, steps).ElementCount

				AddToChain(newPair, stepCount-1, checkCount)
			} else {
				checkCount = chainDuplicates[GetChainKey(newPair, stepCount-1)]
			}

			for k, v := range checkCount {
				polymer.ElementCount[k] += v
			}
		}
	}

	return polymer
}

func ParseInput(lines []string) Polymer {
	polymer := Polymer{
		Pairs:        make([]string, 0),
		Rules:        make(map[string]string),
		ElementCount: make(map[string]int),
	}
	reRule := regexp.MustCompile(`([A-Z]{2})\s\-\>\s([A-Z])`)

	for i, line := range lines {
		if i == 0 {
			for i := 1; i < len(line); i++ {
				pair := line[i-1 : i+1]
				polymer.Pairs = append(polymer.Pairs, pair)

				polymer.ElementCount[pair[0:1]]++

				if i == len(line)-1 {
					polymer.ElementCount[pair[1:]]++
				}
			}

			continue
		}

		if len(line) == 0 {
			continue
		}

		// Parse polymer rule
		groups := reRule.FindAllStringSubmatch(line, -1)[0]
		polymer.Rules[groups[1]] = groups[2]
	}

	return polymer
}

func main() {
	polymer := ParseInput(ReadFile("./inputs/1"))
	println(part_one(polymer, 10))
	chainDuplicates = make(map[string]map[string]int)
	println(part_one(polymer, 40))
}
