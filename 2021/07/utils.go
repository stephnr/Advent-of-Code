package main

import (
	"bufio"
	"log"
	"math/big"
	"os"
)

func ReadFile(filepath string) []string {
	lines := make([]string, 0)

	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func ReadInput(num string) <-chan string {
	lines := ReadFile(num)

	// Convert it to a channel
	out := make(chan string, len(lines))
	defer close(out)

	for _, line := range lines {
		out <- line
	}

	return out
}

func MinMax(arr []int) (min, max int) {
	min = arr[0]
	max = arr[0]

	for _, el := range arr {
		if el < min {
			min = el
		}

		if el > max {
			max = el
		}
	}

	return min, max
}

func MinOfBigIntMap(m map[int]*big.Int) (key int, val *big.Int) {
	// Set a dummy start
	for k, _ := range m {
		val = m[k]
		break
	}

	for k, v := range m {
		switch v.Cmp(val) {
		case -1:
			key = k
			val = v
		}
	}

	return key, val
}
