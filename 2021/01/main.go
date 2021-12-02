package main

import "strconv"

func part_one() {
	input_one := ReadInput("1")

	depth_increases := 0

	for i := 0; i < len(input_one)-1; i++ {
		a, _ := strconv.Atoi(input_one[i])
		b, _ := strconv.Atoi(input_one[i+1])

		if b > a {
			depth_increases++
		}
	}

	println(depth_increases)
}

func part_two() {
	in := ReadInput("1")

	depth_increases := 0

	for i := 0; i < len(in)-3; i += 1 {
		a := SumOfArr([]string{in[i], in[i+1], in[i+2]})
		b := SumOfArr([]string{in[i+1], in[i+2], in[i+3]})

		println(a, b)

		if b > a {
			depth_increases++
		}
	}

	println(depth_increases)
}

func main() {
	// part_one()
	part_two()
}
