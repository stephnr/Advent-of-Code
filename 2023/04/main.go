package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func partOne(rows []string) int {
	var result int
	return result
}

func partTwo(rows []string) int {
	var result int
	return result
}

func parseRows(in string) []string {
	return strings.Split(strings.ReplaceAll(in, "\r\n", "\n"), "\n")
}

func main() {
	println(partOne(parseRows(input)))
	println(partTwo(parseRows(input)))
}
