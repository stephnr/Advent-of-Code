package main

import (
	"log"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/8

func run(filepath string, partTwo bool) int {
	sum := 0

	for line := range ReadInput(filepath) {
		if partTwo {
			// Encoding adds surrounding quotes + one backslash per " and \
			sum += 2 + strings.Count(line, `"`) + strings.Count(line, `\`)
			continue
		}

		// Go's unquote handles the same escapes: \\, \" and \xNN
		mem, err := strconv.Unquote(line)
		if err != nil {
			log.Fatalf("bad string literal %s: %v", line, err)
		}

		sum += len(line) - len(mem)
	}

	return sum
}

func main() {
	println(run("./inputs/0.txt", false)) // 12
	println(run("./inputs/0.txt", true))  // 19
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
