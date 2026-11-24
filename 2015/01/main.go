package main

// https://adventofcode.com/2015/day/1

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)
	floor, pos := 0, 0

	for line := range lines {
		for _, ch := range line {
			pos++

			switch ch {
			case '(':
				floor++
			case ')':
				floor--
			}

			// Part Two: first position that enters the basement
			if partTwo && floor == -1 {
				return pos
			}
		}
	}

	return floor
}

func main() {
	println(run("./inputs/0.txt", false))
	println(run("./inputs/0.txt", true))
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
