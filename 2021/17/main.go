package main

import (
	"math"
)

type Zone struct {
	X []int
	Y []int
}

type Match struct {
	X, Y int
}

var zone = Zone{
	X: []int{253, 280},
	Y: []int{-46, -73},
}

func solve(zone Zone) (matches []Match, maxHeight int) {
	minXSteps := 0

	tX0 := zone.X[0]

	for tX0 > 0 {
		minXSteps++
		tX0 -= minXSteps
	}

	for x := minXSteps; x <= zone.X[1]*2; x++ {
		for y := zone.Y[1]; y <= int(math.Abs(float64(zone.Y[1])))*2; y++ {
			maxY := 0
			step, posX, posY := 0, 0, 0
			Vx, Vy := x, y

		Loop:
			for {
				// Apply the velocity
				posX += Vx
				posY += Vy

				if posY > maxY {
					maxY = posY
				}

				Vx--
				Vy--
				step++

				// Horizontal Momentum dies
				if Vx <= 0 {
					Vx = 0
				}

				if posY < zone.Y[1] || (posX < zone.X[0] && Vx == 0) || (posX > zone.X[1]) {
					// No more chances // overshoot or no momentum
					break Loop
				}

				if (posX >= zone.X[0] && posX <= zone.X[1]) &&
					(posY <= zone.Y[0] && posY >= zone.Y[1]) {
					matches = append(matches, Match{x, y})

					maxHeight = int(math.Max(float64(maxY), float64(maxHeight)))

					break Loop
				}
			}
		}
	}

	return matches, maxHeight
}

func ParseInput(input string) (out []string) {
	return out
}

func main() {
	matches, maxHeight := solve(zone)
	println(maxHeight, len(matches))
}
