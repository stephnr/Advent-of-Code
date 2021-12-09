package main

import (
	"sort"
	"strconv"
	"strings"
)

func part_one(grid [][]int) (sum int) {
	for x, row := range grid {
		for y, currHeight := range row {
			if isDeep(grid, x, y) {
				sum += currHeight + 1
			}
		}
	}

	return sum
}

func isDeep(grid [][]int, x, y int) bool {
	isDeep := true

	checkRowBefore := x > 0
	checkRowNext := x < (len(grid) - 1)

	currHeight := grid[x][y]

	if checkRowBefore {
		isDeep = isDeep && (currHeight < grid[x-1][y])
	}

	if checkRowNext {
		isDeep = isDeep && (currHeight < grid[x+1][y])
	}

	// Check left & right
	if y > 0 {
		isDeep = isDeep && (currHeight < grid[x][y-1])
	}

	if y < len(grid[0])-1 {
		isDeep = isDeep && (currHeight < grid[x][y+1])
	}

	return isDeep
}

func part_two(grid [][]int) (sum int) {
	basins := make([]int, 0)
	sum = 1

	for x, row := range grid {
		for y := range row {
			// Basin = expand from lowest point until 9s
			if isDeep(grid, x, y) {
				basins = append(basins, getBasinSize(grid, x, y, 0))
			}
		}
	}

	sort.Ints(basins)

	for _, basin := range basins[len(basins)-3:] {
		sum *= basin
	}

	return sum
}

func getBasinSize(grid [][]int, x, y, size int) int {
	// Perform basin size check
	checkRowBefore := x > 0
	checkRowNext := x < (len(grid) - 1)

	checkPos := func(a, b int) {
		pos := grid[a][b]

		if pos == 9 {
			return
		}

		grid[a][b] = 9
		size = getBasinSize(grid, a, b, size+1)
	}

	// Check x-1
	if checkRowBefore {
		checkPos(x-1, y)
	}

	// Check x+1
	if checkRowNext {
		checkPos(x+1, y)
	}

	// Check left & right
	if y > 0 {
		checkPos(x, y-1)
	}

	if y < len(grid[0])-1 {
		checkPos(x, y+1)
	}

	return size
}

func ParseInput(input []string) (grid [][]int) {
	// Use a pointer map so we can check against nil
	grid = make([][]int, len(input))

	for i, line := range input {
		for _, rawHeight := range strings.Split(line, "") {
			height, _ := strconv.Atoi(rawHeight)
			grid[i] = append(grid[i], height)
		}
	}

	return grid
}

func main() {
	grid := ParseInput(ReadFile("./inputs/1"))
	println(part_one(grid))
	println(part_two(grid))
}
