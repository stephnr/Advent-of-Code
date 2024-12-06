package main

import (
	"fmt"
	"strings"
)

type Cursor struct {
	X         int
	Y         int
	Direction int // 0 = left, 1 = up, 2 = right, 3 = down
}

func run(filepath string) int {
	lines := ReadInput(filepath)
	area := [][]string{}
	sum := 0

	guard := Cursor{X: 0, Y: 0}

	for line := range lines {
		if strings.Contains(line, "^") {
			guard = Cursor{X: len(area), Y: strings.Index(line, "^"), Direction: 1}
		}

		area = append(area, strings.Split(line, ""))
	}

	area[guard.X][guard.Y] = "X"
	sum++

	// Walk the guard
	for (guard.X >= 0 && guard.X <= len(area)) &&
		(guard.Y >= 0 && guard.Y <= len(area[0])) {
		// printArea(area)
		if guard.Direction == 0 {
			// LEFT
			if guard.Y == 0 {
				// Exit
				break
			}

			if area[guard.X][guard.Y-1] == "." {
				area[guard.X][guard.Y-1] = "X"
				sum++
				guard.Y--
			} else if area[guard.X][guard.Y-1] == "X" {
				guard.Y--
			} else if area[guard.X][guard.Y-1] == "#" {
				// Turn up
				guard.Direction++
			}
		} else if guard.Direction == 1 {
			// UP
			if guard.X == 0 {
				// Exit
				break
			}

			if area[guard.X-1][guard.Y] == "." {
				area[guard.X-1][guard.Y] = "X"
				sum++
				guard.X--
			} else if area[guard.X-1][guard.Y] == "X" {
				guard.X--
			} else if area[guard.X-1][guard.Y] == "#" {
				// Turn right
				guard.Direction++
			}
		} else if guard.Direction == 2 {
			// RIGHT
			if guard.Y+1 >= len(area[0]) {
				// Exit
				break
			}

			if area[guard.X][guard.Y+1] == "." {
				area[guard.X][guard.Y+1] = "X"
				sum++
				guard.Y++
			} else if area[guard.X][guard.Y+1] == "X" {
				guard.Y++
			} else if area[guard.X][guard.Y+1] == "#" {
				// Turn right
				guard.Direction++
			}
		} else if guard.Direction == 3 {
			// DOWN
			if guard.X+1 >= len(area) {
				// Exit
				break
			}

			if area[guard.X+1][guard.Y] == "." {
				area[guard.X+1][guard.Y] = "X"
				sum++
				guard.X++
			} else if area[guard.X+1][guard.Y] == "X" {
				guard.X++
			} else if area[guard.X+1][guard.Y] == "#" {
				// Turn right
				guard.Direction = 0
			}
		}
	}

	return sum
}

func printArea(a [][]string) {
	for _, line := range a {
		fmt.Printf("%+v\n", strings.Join(line, ""))
	}

	fmt.Println()
}

func main() {
	println(run("./inputs/1.txt"))
}
