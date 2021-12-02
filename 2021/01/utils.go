package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
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

func StringArrAsChan(strs []string) <-chan string {
	// Convert it to a channel
	out := make(chan string, len(strs))
	defer close(out)

	for _, str := range strs {
		out <- str
	}

	return out
}

func ReadInput(num string) []string {
	root, _ := os.Getwd()
	inputsPath := path.Join(root, "inputs")
	return ReadFile(path.Join(inputsPath, num))
}

func SumOfArr(nums []string) (sum int) {
	for _, num := range nums {
		val, _ := strconv.Atoi(num)
		sum += val
	}

	return sum
}
