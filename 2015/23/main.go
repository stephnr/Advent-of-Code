package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/23

func execute(filepath string, a int) map[string]int {
	// "jio a, +19" -> ["jio", "a", "+19"]
	prog := [][]string{}
	for line := range ReadInput(filepath) {
		prog = append(prog, strings.Fields(strings.ReplaceAll(line, ",", "")))
	}

	regs := map[string]int{"a": a, "b": 0}

	for pc := 0; pc >= 0 && pc < len(prog); {
		in := prog[pc]

		switch in[0] {
		case "hlf":
			regs[in[1]] /= 2
			pc++
		case "tpl":
			regs[in[1]] *= 3
			pc++
		case "inc":
			regs[in[1]]++
			pc++
		case "jmp":
			pc += offset(in[1])
		case "jie": // jump if even
			if regs[in[1]]%2 == 0 {
				pc += offset(in[2])
			} else {
				pc++
			}
		case "jio": // jump if ONE (not odd)
			if regs[in[1]] == 1 {
				pc += offset(in[2])
			} else {
				pc++
			}
		}
	}

	return regs
}

func offset(s string) int {
	n, _ := strconv.Atoi(s) // handles "+19" and "-7"
	return n
}

func run(filepath string, partTwo bool) int {
	// Part Two: register a starts at 1
	a := 0
	if partTwo {
		a = 1
	}

	return execute(filepath, a)["b"]
}

func main() {
	println(execute("./inputs/0.txt", 0)["a"]) // 2
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
