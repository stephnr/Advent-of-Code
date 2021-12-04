package main

import (
	"strconv"
	"strings"
)

var (
	boardSize = 5
	callouts  = make([]int, 0)
)

func part_one(boards []*BingoBoard) int {
	for _, callout := range callouts {
		for _, board := range boards {
			board.Mark(callout)

			if board.HasBingo() {
				return board.Score(callout)
			}
		}
	}

	return 0
}

func part_two(boards []*BingoBoard) int {
	boardsLeft := make(map[int]int)

	for _, callout := range callouts {
		for boardID, board := range boards {
			board.Mark(callout)

			if board.HasBingo() && boardsLeft[boardID] == 0 {
				boardsLeft[boardID] = callout
			}

			if len(boardsLeft) == len(boards) {
				return board.Score(callout)
			}
		}
	}

	return 0
}

func main() {
	lines := ReadFile("./inputs/1")
	rawBoards := lines[2:]

	// Parse the bingo # callouts
	for _, num := range strings.Split(lines[0], ",") {
		callout, _ := strconv.Atoi(num)
		callouts = append(callouts, callout)
	}

	boards := ParseBoards(rawBoards)

	println(part_one(boards))
	println(part_two(boards))
}
