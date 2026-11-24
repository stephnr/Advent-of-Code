package main

// https://adventofcode.com/2015/day/18

func animate(filepath string, steps int, partTwo bool) int {
	grid := [][]bool{}
	for line := range ReadInput(filepath) {
		row := make([]bool, len(line))
		for i, ch := range line {
			row[i] = ch == '#'
		}
		grid = append(grid, row)
	}

	rows, cols := len(grid), len(grid[0])

	// Part Two: the four corner lights are stuck on
	stick := func(g [][]bool) {
		if partTwo {
			g[0][0], g[0][cols-1], g[rows-1][0], g[rows-1][cols-1] = true, true, true, true
		}
	}
	stick(grid)

	for s := 0; s < steps; s++ {
		next := make([][]bool, rows)

		for r := 0; r < rows; r++ {
			next[r] = make([]bool, cols)

			for c := 0; c < cols; c++ {
				on := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						nr, nc := r+dr, c+dc
						if (dr != 0 || dc != 0) && nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] {
							on++
						}
					}
				}

				// Stays on with 2-3 neighbours on; turns on with exactly 3
				next[r][c] = on == 3 || (grid[r][c] && on == 2)
			}
		}

		grid = next
		stick(grid)
	}

	count := 0
	for _, row := range grid {
		for _, on := range row {
			if on {
				count++
			}
		}
	}

	return count
}

func run(filepath string, partTwo bool) int {
	return animate(filepath, 100, partTwo)
}

func main() {
	println(animate("./inputs/0.txt", 4, false)) // 4
	println(animate("./inputs/0.txt", 5, true))  // 17
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
