package main

import (
	"fmt"
	"strings"
)

type Coordinates struct {
	X int
	Y int
}

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)
	row, sum := 0, 0

	grid := [][]string{}
	antennas := map[string][]Coordinates{}

	for line := range lines {
		grid = append(grid, strings.Split(line, ""))

		for j, ch := range strings.Split(line, "") {
			if ch != "." {
				antennas[ch] = append(antennas[ch], Coordinates{
					X: row,
					Y: j,
				})
			}
		}

		row++
	}

	// Iterate antennas to draw out
	for symbol, antenna := range antennas {
		for i := 0; i < len(antenna)-1; i++ {
			for j := 1; j < len(antenna); j++ {
				if i == j {
					continue
				}

				grid = checkAntenna(symbol, antenna[i], antenna[j], grid, partTwo)
			}
		}
	}

	fmt.Printf("%+v\n", antennas)

	// Count antinodes
	for _, row := range grid {
		for _, col := range row {
			if col == "#" {
				sum++
			}
		}
	}

	if partTwo {
		for _, a := range antennas {
			sum += len(a)
		}
	}

	return sum
}

// Checks two antennas & updates the grid map
func checkAntenna(symbol string, a Coordinates, b Coordinates, grid [][]string, partTwo bool) [][]string {
	breakpoint := 1

	xD := a.X - b.X
	yD := a.Y - b.Y

	if xD == 0 && yD == 0 {
		return grid
	}

	if partTwo {
		breakpoint = len(grid)
	}

	//------------

	// Left check
	for i := 1; i <= breakpoint; i++ {
		if (a.X-(i*xD) < 0 || a.Y-(i*yD) < 0) ||
			(a.X-(i*xD) > len(grid)-1 || a.Y-(i*yD) > len(grid[0])-1) {
			break
		}

		pos := grid[a.X-(i*xD)][a.Y-(i*yD)]

		if pos != symbol && pos != "#" {
			if pos == "." {
				grid[a.X-(i*xD)][a.Y-(i*yD)] = "#"
			}

			printGrid(grid)
		}
	}

	// Right check
	for i := 1; i <= breakpoint; i++ {
		if (a.X+(i*xD) < 0 || a.Y+(i*yD) < 0) ||
			(a.X+(i*xD) > len(grid)-1 || a.Y+(i*yD) > len(grid[0])-1) {
			break
		}

		pos := grid[a.X+(i*xD)][a.Y+(i*yD)]

		if pos != symbol && pos != "#" {
			if pos == "." {
				grid[a.X+(i*xD)][a.Y+(i*yD)] = "#"
			}

			printGrid(grid)
		}
	}

	//------------

	// Left check
	for i := 1; i <= breakpoint; i++ {
		if (b.X-(i*xD) < 0 || b.Y-(i*yD) < 0) ||
			(b.X-(i*xD) > len(grid)-1 || b.Y-(i*yD) > len(grid[0])-1) {
			break
		}

		pos := grid[b.X-(i*xD)][b.Y-(i*yD)]

		if pos != symbol && pos != "#" {
			if pos == "." {
				grid[b.X-(i*xD)][b.Y-(i*yD)] = "#"
			}

			printGrid(grid)
		}
	}

	// Right check
	for i := 1; i <= breakpoint; i++ {
		if (b.X+(i*xD) < 0 || b.Y+(i*yD) < 0) ||
			(b.X+(i*xD) > len(grid)-1 || b.Y+(i*yD) > len(grid[0])-1) {
			break
		}

		pos := grid[b.X+(i*xD)][b.Y+(i*yD)]

		if pos != symbol && pos != "#" {
			if pos == "." {
				grid[b.X+(i*xD)][b.Y+(i*yD)] = "#"
			}

			printGrid(grid)
		}
	}

	//------------

	printGrid(grid)

	fmt.Printf("%+v %+v %+v\n", symbol, xD, yD)

	return grid
}

func printGrid(grid [][]string) {
	for _, line := range grid {
		fmt.Printf("%+v\n", line)
	}

	fmt.Println()
}

func main() {
	// println(run("./inputs/0.txt", false))
	println(run("./inputs/1.txt", true))
}
