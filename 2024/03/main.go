package main

import (
	"regexp"
	"strconv"
)

func partOne(filepath string) int {
	lines := ReadInput(filepath)
	rgx := regexp.MustCompile(`mul\((?P<A>-?\d{1,3}),(?P<B>-?\d{1,3})\)`)
	sum := 0

	for line := range lines {
		matches := rgx.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			a, _ := strconv.Atoi(match[1])
			b, _ := strconv.Atoi(match[2])
			sum += (a * b)
		}
	}

	return sum
}

func partTwo(filepath string) int {
	lines := ReadInput(filepath)
	rgx := regexp.MustCompile(`do\(\)|don't\(\)|mul\((?P<A>-?\d{1,3}),(?P<B>-?\d{1,3})\)`)
	sum := 0
	makeDo := true

	for line := range lines {
		matches := rgx.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			switch match[0] {
			case "do()":
				makeDo = true
				continue
			case "don't()":
				makeDo = false
				continue
			}

			if makeDo {
				a, _ := strconv.Atoi(match[1])
				b, _ := strconv.Atoi(match[2])
				sum += (a * b)
			}

		}
	}

	return sum
}

func main() {
	println(partOne("./inputs/1.txt"))
	println(partTwo("./inputs/1.txt"))
}
