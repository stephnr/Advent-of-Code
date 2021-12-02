package main

import (
	"strconv"
	"strings"
)

func part_one() int {
	x, y := 0, 0
	input_one := ReadInput("1")

	for line := range input_one {
		parts := strings.Split(line, " ")
		direction, val_str := parts[0][0], parts[1]
		val, _ := strconv.Atoi(val_str)

		switch direction {
		case ('f'):
			x += val
		case ('u'):
			y -= val
		case ('d'):
			y += val
		}
	}

	return x * y
}

func part_two() int {
	x, y, aim := 0, 0, 0
	input_one := ReadInput("1")

	for line := range input_one {
		parts := strings.Split(line, " ")
		direction, val_str := parts[0][0], parts[1]
		val, _ := strconv.Atoi(val_str)

		switch direction {
		case ('f'):
			x += val
			y += (aim * val)
		case ('u'):
			aim -= val
		case ('d'):
			aim += val
		}

		if aim < 0 {
			aim = 0
		}
	}

	return x * y
}

func main() {
	println(part_one())
	println(part_two())
}
