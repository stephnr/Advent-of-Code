package main

import (
	"math"
	"regexp"
	"strconv"
)

type Geyser struct {
	x1, x2 int
	y1, y2 int
}

func part_one(geysers []Geyser, plotDiagonals bool) (multiVents int) {
	ventMap := BuildIntMatrix(999)

	// Plot the vent map
	for _, geyser := range geysers {
		if (geyser.x1 == geyser.x2) || (geyser.y1 == geyser.y2) {
			// Plot verticals
			if geyser.y1 == geyser.y2 {
				low := math.Min(float64(geyser.x1), float64(geyser.x2))
				high := math.Max(float64(geyser.x1), float64(geyser.x2))

				for i := low; i <= high; i++ {
					ventMap[int(i)][geyser.y1] = ventMap[int(i)][geyser.y1] + 1

					if ventMap[int(i)][geyser.y1] == 2 {
						multiVents++
					}
				}
			}

			// Plot horizontals
			if geyser.x1 == geyser.x2 {
				low := math.Min(float64(geyser.y1), float64(geyser.y2))
				high := math.Max(float64(geyser.y1), float64(geyser.y2))

				for i := low; i <= high; i++ {
					ventMap[geyser.x1][int(i)] = ventMap[geyser.x1][int(i)] + 1

					if ventMap[geyser.x1][int(i)] == 2 {
						multiVents++
					}
				}
			}
		} else {
			// Optional: Plot diagonals
			if plotDiagonals {
				xUp, yUp := false, false

				if geyser.x1 > geyser.x2 {
					xUp = true
				}

				if geyser.y1 > geyser.y2 {
					yUp = true
				}

				x, y := geyser.x1, geyser.y1

				for {
					// Plot
					ventMap[x][y] = ventMap[x][y] + 1

					// Check
					if ventMap[x][y] == 2 {
						multiVents++
					}

					if x == geyser.x2 || y == geyser.y2 {
						break
					}

					// Increment
					if xUp {
						x--
					} else {
						x++
					}

					if yUp {
						y--
					} else {
						y++
					}
				}
			}
		}
	}

	return multiVents
}

func ParseInput(rawInput []string) (geysers []Geyser) {
	re := regexp.MustCompile(`^(\d+)\,(\d+) \-\> (\d+)\,(\d+)$`)

	for _, line := range rawInput {
		var groupVals []int
		groups := re.FindAllStringSubmatch(line, -1)

		for i := 1; i < len(groups[0]); i++ {
			val, _ := strconv.Atoi(groups[0][i])
			groupVals = append(groupVals, val)
		}

		geysers = append(geysers, Geyser{groupVals[0], groupVals[2], groupVals[1], groupVals[3]})
	}

	return geysers
}

func main() {
	geysers := ParseInput(ReadFile("./inputs/1"))

	println(part_one(geysers, false))
	println(part_one(geysers, true))
}
