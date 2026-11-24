package main

import (
	"crypto/md5"
	"strconv"
	"strings"
)

// https://adventofcode.com/2015/day/4

func run(filepath string, partTwo bool) int {
	key := ""
	for line := range ReadInput(filepath) {
		key += strings.TrimSpace(line)
	}

	buf := []byte(key)

	for n := 1; ; n++ {
		// Reuse the key prefix and only rewrite the number each attempt
		sum := md5.Sum(strconv.AppendInt(buf[:len(key)], int64(n), 10))

		// Each byte is 2 hex chars: 5 zeros = 2 zero bytes + a zero high nibble,
		// 6 zeros = 3 zero bytes
		if sum[0] != 0 || sum[1] != 0 {
			continue
		}

		if partTwo && sum[2] == 0 {
			return n
		}

		if !partTwo && sum[2] < 0x10 {
			return n
		}
	}
}

func main() {
	println(run("./inputs/0.txt", false))
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
