package main

// https://adventofcode.com/2015/day/3

type Coordinates struct {
	X int
	Y int
}

func run(filepath string, partTwo bool) int {
	lines := ReadInput(filepath)

	// Part One has only Santa; Part Two adds Robo-Santa taking alternate moves
	santas := make([]Coordinates, 1)
	if partTwo {
		santas = make([]Coordinates, 2)
	}

	visited := map[Coordinates]bool{{X: 0, Y: 0}: true}
	turn := 0

	for line := range lines {
		for _, ch := range line {
			s := &santas[turn%len(santas)]
			turn++

			switch ch {
			case '^':
				s.Y++
			case 'v':
				s.Y--
			case '>':
				s.X++
			case '<':
				s.X--
			}

			visited[*s] = true
		}
	}

	return len(visited)
}

func main() {
	println(run("./inputs/0.txt", false))
	println(run("./inputs/0.txt", true))
	println(run("./inputs/1.txt", false))
	println(run("./inputs/1.txt", true))
}
