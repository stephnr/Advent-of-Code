package main

import (
	"strings"
)

type PathTree map[string][]string

func solve(paths PathTree, partTwo bool) (uniquePaths int) {
	result := Traverse(paths, partTwo, "start", []string{}, make(map[string]bool))

	// -- Enable the following for debugging --

	// for path := range result {
	// 	fmt.Println(path)
	// }

	return len(result)
}

func Traverse(paths PathTree, partTwo bool, from string, steps []string, uniquePaths map[string]bool) map[string]bool {
	steps = append(steps, from)

	if from == "end" {
		uniquePaths[strings.Join(steps, ",")] = true
		return uniquePaths
	}

	for _, route := range paths[from] {
		var smallRevisits int
		smallVisits := make(map[string]int)

		if route == "start" {
			continue
		}

		possibleRouteTree := append(steps, route)

		for _, step := range possibleRouteTree {
			// Check if we visited this small cave before?
			if strings.ToLower(step) == step {
				smallVisits[step]++

				if smallVisits[step] >= 2 {
					smallRevisits++
				}
			}
		}

		if !partTwo && smallRevisits > 0 {
			// No revisits allowed
			continue
		}

		if partTwo && smallRevisits > 1 {
			// The next route would be invalid due to small cave revisit twice
			continue
		}

		// New Small Cave or previous Big Cave
		uniquePaths = Traverse(paths, partTwo, route, steps, uniquePaths)
	}

	return uniquePaths
}

func ParseInput(input []string) (paths PathTree) {
	paths = make(PathTree)

	for _, line := range input {
		split := strings.Split(line, "-")
		a, b := split[0], split[1]
		paths[a] = append(paths[a], b)
		paths[b] = append(paths[b], a)
	}

	return paths
}

func main() {
	paths := ParseInput(ReadFile("./inputs/1"))
	println(solve(paths, false))
	println(solve(paths, true))
}
