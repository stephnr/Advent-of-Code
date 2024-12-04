package main

func partOne(filepath string) int {
	lines := ReadInput(filepath)
	sum := 0

	xmasMap := []string{}

	for line := range lines {
		xmasMap = append(xmasMap, line)
	}

	// Search left to right
	for _, line := range xmasMap {
		for i := 0; i < len(xmasMap[0])-3; i++ {
			if compare(line[i : i+4]) {
				sum += 1
			}
		}
	}

	// Search top down
	for i := 0; i < len(xmasMap)-3; i++ {
		for j := 0; j < len(xmasMap[i]); j++ {
			if compare(string([]byte{xmasMap[i][j], xmasMap[i+1][j], xmasMap[i+2][j], xmasMap[i+3][j]})) {
				sum += 1
			}
		}
	}

	// Search diagonally (Left to Right)
	for i := 0; i < len(xmasMap)-3; i++ {
		for j := 0; j < len(xmasMap[i])-3; j++ {
			if compare(string([]byte{xmasMap[i][j], xmasMap[i+1][j+1], xmasMap[i+2][j+2], xmasMap[i+3][j+3]})) {
				sum += 1
			}
		}
	}

	// Search diagonally (Right to Left)
	for i := 0; i < len(xmasMap)-3; i++ {
		for j := 0; j < len(xmasMap[i])-3; j++ {
			if compare(string([]byte{xmasMap[i][j+3], xmasMap[i+1][j+2], xmasMap[i+2][j+1], xmasMap[i+3][j]})) {
				sum++
			}
		}
	}

	return sum
}

func partTwo(filepath string) int {
	lines := ReadInput(filepath)
	sum := 0

	xmasMap := []string{}

	for line := range lines {
		xmasMap = append(xmasMap, line)
	}

	// Skip first/last rows
	for i := 1; i < len(xmasMap)-1; i++ {
		for j := 1; j < len(xmasMap[i])-1; j++ {
			// Check in a 3x3 grid only if curr scan = 'A'
			if xmasMap[i][j] == 'A' {
				if (xmasMap[i-1][j-1] == 'M' && xmasMap[i+1][j+1] == 'S' || xmasMap[i-1][j-1] == 'S' && xmasMap[i+1][j+1] == 'M') &&
					(xmasMap[i-1][j+1] == 'M' && xmasMap[i+1][j-1] == 'S' || xmasMap[i-1][j+1] == 'S' && xmasMap[i+1][j-1] == 'M') {
					sum++
				}
			}
		}
	}

	return sum
}

func compare(s string) bool {
	if s == "XMAS" || s == "SAMX" {
		return true
	}

	return false
}

func main() {
	println(partOne("./inputs/1.txt"))
	println(partTwo("./inputs/1.txt"))
}
