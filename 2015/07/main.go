package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/7

// Circuit maps each wire to the expression that drives it, e.g. "x AND y"
type Circuit map[string]string

func parse(filepath string) Circuit {
	c := Circuit{}

	for line := range ReadInput(filepath) {
		expr, wire, _ := strings.Cut(line, " -> ")
		c[wire] = expr
	}

	return c
}

// Resolves a wire (or literal number) recursively, caching each wire's signal
func eval(c Circuit, cache map[string]uint16, w string) uint16 {
	if n, err := strconv.Atoi(w); err == nil {
		return uint16(n)
	}

	if v, ok := cache[w]; ok {
		return v
	}

	var v uint16
	f := strings.Fields(c[w])

	switch len(f) {
	case 1: // "x"
		v = eval(c, cache, f[0])
	case 2: // "NOT x"
		v = ^eval(c, cache, f[1])
	case 3: // "x OP y"
		a, b := eval(c, cache, f[0]), eval(c, cache, f[2])

		switch f[1] {
		case "AND":
			v = a & b
		case "OR":
			v = a | b
		case "LSHIFT":
			v = a << b
		case "RSHIFT":
			v = a >> b
		}
	}

	cache[w] = v
	return v
}

func signal(filepath string, wire string) int {
	return int(eval(parse(filepath), map[string]uint16{}, wire))
}

func run(filepath string, partTwo bool) int {
	c := parse(filepath)
	a := eval(c, map[string]uint16{}, "a")

	if !partTwo {
		return int(a)
	}

	// Override wire b with a's signal and reset all other wires
	c["b"] = strconv.Itoa(int(a))
	return int(eval(c, map[string]uint16{}, "a"))
}

func main() {
	println(signal("./inputs/0.txt", "d")) // 72
	println(signal("./inputs/0.txt", "h")) // 65412
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
