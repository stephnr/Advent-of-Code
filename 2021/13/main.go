package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

type Instruction struct {
	Axis  string
	Value int
}

type DotGrid struct {
	RowSize int
	ColSize int

	Grid [][]int

	Instructions []Instruction
}

func part_one(dotGrid *DotGrid) int {
	dotGrid = FoldGrid(dotGrid)
	// dotGrid.Print()

	return dotGrid.DotCount()
}

func part_two(dotGrid *DotGrid) *DotGrid {
	instrCount := len(dotGrid.Instructions)

	for i := 0; i < instrCount; i++ {
		dotGrid = FoldGrid(dotGrid)
	}

	return dotGrid
}

func (dotGrid *DotGrid) DotCount() (sum int) {
	for _, row := range dotGrid.Grid {
		for _, col := range row {
			if col > 0 {
				sum++
			}
		}
	}

	return sum
}

func (dotGrid *DotGrid) Print() {
	fmt.Println()
	for _, row := range dotGrid.Grid {
		for _, col := range row {
			if col == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Printf("\n")
	}
}

func FoldGrid(dotGrid *DotGrid) *DotGrid {
	var newX, newY int

	instr := dotGrid.Instructions[0]

	newGrid := &DotGrid{
		Grid:         make([][]int, 0),
		Instructions: dotGrid.Instructions[1:],
	}

	switch instr.Axis {
	case "x":
		// Flip along columns
		newY = dotGrid.RowSize
		newX = instr.Value

		// Set up the grid
		for row := 0; row < len(dotGrid.Grid); row++ {
			colMap := make([]int, newX)

			for col := 0; col < len(dotGrid.Grid[0]); col++ {
				if col == newX {
					continue
				}

				if col > newX*2 {
					break
				}

				if col < newX {
					// Copy value
					colMap[col] = dotGrid.Grid[row][col]
				} else {
					// Overlay value
					sourceVal := dotGrid.Grid[row][col]
					colMap[newX-(col-newX)] += sourceVal
				}
			}

			newGrid.Grid = append(newGrid.Grid, colMap)
		}

	case "y":
		// Rows slashed
		newY = instr.Value
		newX = dotGrid.ColSize

		// Set up the grid
		for row := 0; row < len(dotGrid.Grid); row++ {
			if row == newY {
				continue
			}

			if row > newY*2 {
				break
			}

			if row <= instr.Value {
				// Copy exact values
				newGrid.Grid = append(newGrid.Grid, dotGrid.Grid[row])
			} else {
				for col := 0; col < len(dotGrid.Grid[0]); col++ {
					// Flipped values (write to copied row)
					sourceVal := dotGrid.Grid[row][col]
					newGrid.Grid[newY-(row-newY)][col] += sourceVal
				}
			}
		}

		// Slash the grid
		newGrid.Grid = newGrid.Grid[0:instr.Value]
	default:
		panic("Unknown Axis Value")
	}

	newGrid.RowSize = newY
	newGrid.ColSize = newX

	return newGrid
}

func ParseInput(input []string) (out *DotGrid) {
	dotGrid := &DotGrid{
		Grid: make([][]int, 2000),
	}

	reFold := regexp.MustCompile(`([x,y])\=(\d+)`)

	for _, line := range input {
		if line == "" {
			continue
		}

		if strings.HasPrefix(line, "fold") {
			groups := reFold.FindAllStringSubmatch(line, -1)
			size, _ := strconv.Atoi(groups[0][2])
			dotGrid.Instructions = append(dotGrid.Instructions, Instruction{groups[0][1], size})
			continue
		}

		// Parse dot
		vals := strings.Split(line, ",")
		y, _ := strconv.Atoi(vals[0])
		x, _ := strconv.Atoi(vals[1])

		if dotGrid.Grid[x] == nil {
			dotGrid.Grid[x] = make([]int, 2000)
		}

		dotGrid.Grid[x][y] = 1

		dotGrid.RowSize = int(math.Max(float64(dotGrid.RowSize), float64(x)))
		dotGrid.ColSize = int(math.Max(float64(dotGrid.ColSize), float64(y)))
	}

	// Shrink the int matrix
	dotGrid.Grid = dotGrid.Grid[:dotGrid.RowSize+1]

	for i := range dotGrid.Grid {
		if len(dotGrid.Grid[i]) == 0 {
			dotGrid.Grid[i] = make([]int, dotGrid.ColSize+1)
		} else {
			dotGrid.Grid[i] = dotGrid.Grid[i][:dotGrid.ColSize+1]
		}
	}

	return dotGrid
}

func main() {
	input := ParseInput(ReadFile("./inputs/1"))
	println(part_one(input))
	out := part_two(input)
	out.Print()
}
