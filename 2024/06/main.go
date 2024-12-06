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

func partOne(filepath string) int {
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

func partTwo(filepath string) int {
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

	// Loop through entire grid & try to form a loop
	for i, row := range area {
		for j, _ := range row {
			aa := [][]string{}

			for _, row := range area {
				aa = append(aa, strings.Split(strings.Clone(strings.Join(row, "")), ""))
			}

			if aa[i][j] == "." {
				aa[i][j] = "#"
				if loopCheck(Cursor{X: guard.X, Y: guard.Y, Direction: guard.Direction}, aa) {
					sum++
				}
			}
		}
	}

	return sum
}

func loopCheck(g Cursor, a [][]string) bool {
	// Loop found if we reach the position + direction again
	steps := 0

	// Walk the guard
	for (g.X >= 0 && g.X <= len(a)) &&
		(g.Y >= 0 && g.Y <= len(a[0])) {
		// Check if loop
		if steps >= 100000 {
			return true
		}

		steps++

		if g.Direction == 0 {
			// LEFT
			if g.Y == 0 {
				// Exit
				break
			}

			if a[g.X][g.Y-1] == "." {
				a[g.X][g.Y-1] = "X"
				// steps++
				g.Y--
			} else if a[g.X][g.Y-1] == "X" {
				g.Y--
			} else if a[g.X][g.Y-1] == "#" {
				// Turn up
				g.Direction++
			}
		} else if g.Direction == 1 {
			// UP
			if g.X == 0 {
				// Exit
				break
			}

			if a[g.X-1][g.Y] == "." {
				a[g.X-1][g.Y] = "X"
				// steps++
				g.X--
			} else if a[g.X-1][g.Y] == "X" {
				g.X--
			} else if a[g.X-1][g.Y] == "#" {
				// Turn right
				g.Direction++
			}
		} else if g.Direction == 2 {
			// RIGHT
			if g.Y+1 >= len(a[0]) {
				// Exit
				break
			}

			if a[g.X][g.Y+1] == "." {
				a[g.X][g.Y+1] = "X"
				// steps++
				g.Y++
			} else if a[g.X][g.Y+1] == "X" {
				g.Y++
			} else if a[g.X][g.Y+1] == "#" {
				// Turn right
				g.Direction++
			}
		} else if g.Direction == 3 {
			// DOWN
			if g.X+1 >= len(a) {
				// Exit
				break
			}

			if a[g.X+1][g.Y] == "." {
				a[g.X+1][g.Y] = "X"
				// steps++
				g.X++
			} else if a[g.X+1][g.Y] == "X" {
				g.X++
			} else if a[g.X+1][g.Y] == "#" {
				// Turn right
				g.Direction = 0
			}
		}
	}

	return false
}

func printArea(a [][]string) {
	for _, line := range a {
		fmt.Printf("%+v\n", strings.Join(line, ""))
	}

	fmt.Println()
}

func main() {
	println(partOne("./inputs/1.txt"))
	println(partTwo("./inputs/1.txt"))
}
