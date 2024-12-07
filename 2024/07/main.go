package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Equation struct {
	Result int
	Nums   []int
}

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)
	sum := 0

	equations := readEquations(lines)

	for _, eq := range equations {
		if permuteEquations(eq, partTwo) > 0 {
			sum += eq.Result
		}
	}

	return sum
}

func permuteEquations(eq Equation, partTwo bool) int {
	steps := 0

	if eq.Nums[0] > eq.Result {
		return 0
	}

	if len(eq.Nums) == 2 {

		// Check with addition
		if eq.Nums[0]+eq.Nums[1] == eq.Result {
			steps++
		}

		// Check with multiplication
		if eq.Nums[0]*eq.Nums[1] == eq.Result {
			steps++
		}

		// Check with || concat operator
		if partTwo {
			a := fmt.Sprintf("%d%d", eq.Nums[0], eq.Nums[1])
			b, _ := strconv.Atoi(a)

			if b == eq.Result {
				steps++
			}
		}

		return steps
	}

	// Flatten & continue on Addition
	steps += permuteEquations(Equation{
		Result: eq.Result,
		Nums:   append([]int{eq.Nums[0] + eq.Nums[1]}, eq.Nums[2:]...),
	}, partTwo)

	// Flatten & continue on Multiply
	steps += permuteEquations(Equation{
		Result: eq.Result,
		Nums:   append([]int{eq.Nums[0] * eq.Nums[1]}, eq.Nums[2:]...),
	}, partTwo)

	if partTwo {
		// Flatten & continue if part two
		a := fmt.Sprintf("%d%d", eq.Nums[0], eq.Nums[1])
		b, _ := strconv.Atoi(a)

		steps += permuteEquations(Equation{
			Result: eq.Result,
			Nums:   append([]int{b}, eq.Nums[2:]...),
		}, partTwo)
	}

	return steps
}

func readEquations(lines <-chan string) []Equation {
	equations := []Equation{}

	for line := range lines {
		// add empty line on input to allow for regex to work
		if len(line) == 0 {
			break
		}

		a := strings.Split(line, ":")
		b := strings.Split(a[1], " ")

		aa, _ := strconv.Atoi(a[0])

		eq := Equation{
			Result: aa,
			Nums:   []int{},
		}

		for _, val := range b[1:] {
			vv, _ := strconv.Atoi(val)
			eq.Nums = append(eq.Nums, vv)
		}

		equations = append(equations, eq)
	}

	return equations
}

func main() {
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
