package main

import (
	"fmt"
	"strconv"
)

type Octo struct {
	ID      string
	Charge  int
	Flashes map[int]bool
}

func part_one(octoGrid [][]Octo, steps int) (sum int) {
	// Charge grows

	for i := 0; i < steps; i++ {
		_, flashes := RunStep(octoGrid, i, 0)
		sum += flashes
	}

	return sum
}

func part_two(octoGrid [][]Octo) int {
	// Charge grows
	i := 0

	for {
		i++
		_, count := RunStep(octoGrid, i, 0)

		if count == 100 {
			return i
		}
	}
}

func RunStep(octoGrid [][]Octo, step, flashCount int) ([][]Octo, int) {
	// Handle pulses
	for row := 0; row < len(octoGrid); row++ {
		for col := 0; col < len(octoGrid[row]); col++ {
			// Charge grows only if they haven't flashed
			if !octoGrid[row][col].Flashes[step] {
				octoGrid[row][col].Charge = octoGrid[row][col].Charge + 1
			}

			// Check if pulse ready but hasn't flashed
			if octoGrid[row][col].Charge > 9 {
				octoGrid, flashCount = HandleFlash(octoGrid, row, col, step, flashCount)
			}
		}
	}

	return octoGrid, flashCount
}

func HandleFlash(octoGrid [][]Octo, x, y, step, flashCount int) ([][]Octo, int) {
	// Run step if this octo hasn't flashed yet but was pulsed
	if !octoGrid[x][y].Flashes[step] && (x < 0 || x > len(octoGrid)) || (y < 0 || y > len(octoGrid[0])) {
		return octoGrid, flashCount
	}

	flashCount++
	octoGrid[x][y].Flashes[step] = true
	octoGrid[x][y].Charge = 0

	// Begin the loop
	checkForFlash := func(a, b int) {
		if octoGrid[a][b].Flashes[step] {
			// No action since its flashing this round
			return
		}

		octoGrid[a][b].Charge = octoGrid[a][b].Charge + 1

		if octoGrid[a][b].Charge > 9 {
			octoGrid, flashCount = HandleFlash(octoGrid, a, b, step, flashCount)
		}
	}

	// Prev Row
	if x-1 >= 0 {
		checkForFlash(x-1, y)

		// Prev Row Diagonals
		if y-1 >= 0 {
			checkForFlash(x-1, y-1)
		}

		// Next Col
		if y+1 <= len(octoGrid[0])-1 {
			checkForFlash(x-1, y+1)
		}
	}

	// Prev Col
	if y-1 >= 0 {
		checkForFlash(x, y-1)
	}

	// Next Col
	if y+1 <= len(octoGrid[0])-1 {
		checkForFlash(x, y+1)
	}

	// Next Row
	if x+1 <= len(octoGrid)-1 {
		checkForFlash(x+1, y)

		// Prev Row Diagonals
		if y-1 >= 0 {
			checkForFlash(x+1, y-1)
		}

		// Next Col
		if y+1 <= len(octoGrid[0])-1 {
			checkForFlash(x+1, y+1)
		}
	}

	return octoGrid, flashCount
}

func PrintGrid(octoGrid [][]Octo) {
	for _, row := range octoGrid {
		for _, octo := range row {
			fmt.Printf("%d", octo.Charge)
		}
		fmt.Printf("\n")
	}
}

func ParseInput(input []string) (octos [][]Octo) {
	for row, line := range input {
		octos = append(octos, make([]Octo, len(line)))

		for col, valRaw := range line {
			val, _ := strconv.Atoi(string(valRaw))
			octos[row][col] = Octo{
				ID:      fmt.Sprintf("%d:%d", row, col),
				Charge:  val,
				Flashes: make(map[int]bool),
			}
		}
	}

	return octos
}

func main() {
	octoGridA := ParseInput(ReadFile("./inputs/1"))
	octoGridB := ParseInput(ReadFile("./inputs/1"))
	println(part_one(octoGridA, 100))
	println(part_two(octoGridB))
}
