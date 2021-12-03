package main

import (
	"bufio"
	"log"
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
