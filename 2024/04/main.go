package main

import "fmt"

func run(filepath string) int {
	lines := ReadInput(filepath)
	sum := 0

	fmt.Printf("%+v\n", lines)

	return sum
}

func main() {
	println(run("./inputs/1.txt"))
}
