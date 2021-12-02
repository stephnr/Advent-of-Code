package main

import (
	"bufio"
	"log"
	"os"
	"path"
)

func ReadFileAsChan(filepath string) <-chan string {
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

	// Convert it to a channel
	out := make(chan string, len(lines))
	defer close(out)

	for _, line := range lines {
		out <- line
	}

	return out
}

func ReadInput(num string) <-chan string {
	root, _ := os.Getwd()
	inputsPath := path.Join(root, "inputs")
	return ReadFileAsChan(path.Join(inputsPath, num))
}
