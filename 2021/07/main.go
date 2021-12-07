package main

import (
	"math"
	"math/big"
	"strconv"
	"strings"
)

func part_one(crabs []int, partOne bool) (pos int, fuel *big.Int) {
	_, max := MinMax(crabs)

	// target => sum of movements
	distMap := make(map[int]*big.Int)

	for i := 1; i <= max; i++ {
		for _, crab := range crabs {
			if partOne {
				bigInt := new(big.Int)
				fuelCost := bigInt.SetInt64(int64(math.Abs(float64(crab) - float64(i))))

				if distMap[i] == nil {
					distMap[i] = fuelCost
				} else {
					distMap[i] = bigInt.Add(fuelCost, distMap[i])
				}
			} else {
				// Compute the binomial coefficient
				bigInt := new(big.Int)
				fuelCost := bigInt.Binomial(int64(math.Abs(float64(crab)-float64(i)))+1, 2)

				if distMap[i] == nil {
					distMap[i] = fuelCost
				} else {
					distMap[i] = bigInt.Add(fuelCost, distMap[i])
				}
			}
		}
	}

	k, v := MinOfBigIntMap(distMap)
	return k, v
}

func ParseCrabs(input string) (out []int) {
	for _, val := range strings.Split(input, ",") {
		crab, _ := strconv.Atoi(val)
		out = append(out, crab)
	}

	return out
}

func main() {
	crabs := ParseCrabs(ReadFile("./inputs/1")[0])

	k, v := part_one(crabs, true)
	println(k, v.String())

	k, v = part_one(crabs, false)
	println(k, v.String())
}
