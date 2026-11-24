package main

import (
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/10

func run(filepath string, partTwo bool) int {
	times := 40
	if partTwo {
		times = 50
	}

	s := ""
	for line := range ReadInput(filepath) {
		s += strings.TrimSpace(line)
	}

	return len(expand(s, times))
}

func expand(s string, times int) string {
	for i := 0; i < times; i++ {
		s = lookAndSay(s)
	}

	return s
}

// "111221" -> "312211": each run of a digit becomes <count><digit>
func lookAndSay(s string) string {
	var b strings.Builder

	for i := 0; i < len(s); {
		j := i
		for j < len(s) && s[j] == s[i] {
			j++
		}

		b.WriteString(strconv.Itoa(j - i))
		b.WriteByte(s[i])
		i = j
	}

	return b.String()
}

func main() {
	println(expand("1", 5)) // 312211
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
