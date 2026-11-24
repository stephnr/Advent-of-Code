package main

// https://adventofcode.com/2016/day/8

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)
	sum := 0

	for line := range lines {
		_ = line
	}

	if partTwo {
		return sum
	}

	return sum
}

func main() {
	println(run("./inputs/0.txt", false))
	// println(run("./inputs/1.txt", false))
	// println(run("./inputs/1.txt", true))
}
