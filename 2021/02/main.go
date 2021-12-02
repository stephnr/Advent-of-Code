package main

import (
	"strconv"
	"strings"
)

func part_one() {
	x, y := 0, 0
	input_one := ReadInput("1")

	for line := range input_one {
		parts := strings.Split(line, " ")
		direction, val_str := parts[0], parts[1]
		val, _ := strconv.Atoi(val_str)

		switch direction {
		case ("forward"):
			x += val
		case ("up"):
			y -= val
		case ("down"):
			y += val
		}
	}

	println(x * y)
}

func part_two() {
	x, y, aim := 0, 0, 0
	input_one := ReadInput("1")

	for line := range input_one {
		parts := strings.Split(line, " ")
		direction, val_str := parts[0], parts[1]
		val, _ := strconv.Atoi(val_str)

		switch direction {
		case ("forward"):
			x += val
			y += (aim * val)
		case ("up"):
			aim -= val
		case ("down"):
			aim += val
		}

		if aim < 0 {
			aim = 0
		}
	}

	println(x * y)
}

func main() {
	part_one()
	part_two()
}
