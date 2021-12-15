package main

import (
	"fmt"
	"strconv"
)

type Path struct {
	X        int
	Y        int
	Value    int
	BestPath []int
}

func CountPath(path []int) (sum int) {
	for _, v := range path {
		sum += v
	}

	return sum
}

func (p *Path) Key() string {
	return fmt.Sprintf("%d:%d", p.X, p.Y)
}

func (p *Path) Visit(pathTree []int) bool {
	a, b := CountPath(pathTree), CountPath(p.BestPath)

	if a < b || b == 0 {
		p.BestPath = pathTree
		return true
	}

	return false
}

func CopyPos(p []int, value int) []int {
	// Add both as candidates
	tmp := make([]int, len(p))
	copy(tmp, p)
	tmp = append(tmp, value)

	return tmp
}

func part_one(grid [][]Path) int {
	unvisited := make(map[string]*Path)
	unvisited[grid[0][0].Key()] = &grid[0][0]

	finalKey := fmt.Sprintf("%d:%d", len(grid)-1, len(grid)-1)

	for {
		if unvisited[finalKey] != nil {
			return CountPath(unvisited[finalKey].BestPath)
		}

		nextVisited := make(map[string]*Path)

		for _, p := range unvisited {
			// Visit the neighbors
			if p.X+1 < len(grid[0]) {
				next := &grid[p.X+1][p.Y]

				if nextVisited[next.Key()] != nil {
					next = nextVisited[next.Key()]
				}

				// Visit Right
				next.Visit(CopyPos(p.BestPath, next.Value))
				nextVisited[next.Key()] = next
			}

			if p.Y+1 < len(grid) {
				// Visit Down
				next := &grid[p.X][p.Y+1]

				if nextVisited[next.Key()] != nil {
					next = nextVisited[next.Key()]
				}

				// Visit Right
				next.Visit(CopyPos(p.BestPath, next.Value))
				nextVisited[next.Key()] = next
			}
		}

		unvisited = nextVisited
	}
}

func ParseInput(input []string) (grid [][]Path) {
	grid = make([][]Path, len(input))

	for i, line := range input {
		newLine := make([]Path, len(line))

		for j, r := range line {
			val, _ := strconv.Atoi(string(r))
			newLine[j] = Path{
				X:        i,
				Y:        j,
				Value:    val,
				BestPath: make([]int, 0),
			}
		}

		grid[i] = newLine
	}

	return grid
}

func main() {
	grid := ParseInput(ReadFile("./inputs/1"))
	println(part_one(grid))
	// println(part_two())
}
