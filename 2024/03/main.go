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
		for _, match := range rgx.FindAllStringSubmatch(line, -1) {
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
	doStuff := true

	for line := range lines {
		for _, match := range rgx.FindAllStringSubmatch(line, -1) {
			if match[0] == "do()" {
				doStuff = true
			} else if match[0] == "don't()" {
				doStuff = false
			} else if doStuff {
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
