package main

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"
)

func InitBenchTest(b *testing.B) []*BingoBoard {
	b.Helper()

	lines := ReadFile("./inputs/1")
	rawBoards := lines[2:]

	// Parse the bingo # callouts
	for _, num := range strings.Split(lines[0], ",") {
		callout, _ := strconv.Atoi(num)
		callouts = append(callouts, callout)
	}

	return ParseBoards(rawBoards)
}

func BenchmarkPartOne(b *testing.B) {
	boards := InitBenchTest(b)

	for n := 0; n <= b.N; n++ {
		part_one(boards)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	boards := InitBenchTest(b)

	for n := 0; n <= b.N; n++ {
		part_two(boards)
	}
}

func TestHasBingo_Vertical(t *testing.T) {
	board := BingoBoard{board: make([][]*BingoMark, boardSize)}
	boardPassMarks := make([]*BingoMark, 0)

	for i := 0; i < boardSize; i++ {
		boardPassMarks = append(boardPassMarks, &BingoMark{rand.Intn(99), i == 0})
	}

	for i := 0; i < boardSize; i++ {
		board.board[i] = boardPassMarks
	}

	if !board.HasBingo() {
		t.Fail()
	}
}

func TestHasBingo_Horizontal(t *testing.T) {
	board := BingoBoard{board: make([][]*BingoMark, boardSize)}
	boardPassMarks := make([]*BingoMark, 0)

	for i := 0; i < boardSize; i++ {
		boardPassMarks = append(boardPassMarks, &BingoMark{rand.Intn(99), true})
	}

	boardMarks := make([]*BingoMark, 0)

	for i := 0; i < boardSize; i++ {
		boardMarks = append(boardMarks, &BingoMark{rand.Intn(99), false})
	}

	for i := 0; i < boardSize; i++ {
		if i == 0 {
			board.board[i] = boardPassMarks
		} else {
			board.board[i] = boardMarks
		}
	}

	if !board.HasBingo() {
		t.Fail()
	}
}

func TestHasBingo_MarkBingo(t *testing.T) {
	board := BingoBoard{board: make([][]*BingoMark, boardSize)}
	boardMarks := make([]*BingoMark, 0)

	for i := 0; i < boardSize; i++ {
		boardMarks = append(boardMarks, &BingoMark{99, false})
	}

	for i := 0; i < boardSize; i++ {
		board.board[i] = boardMarks
	}

	// Call a number
	board.Mark(99)

	if !board.HasBingo() {
		t.Fail()
	}
}
