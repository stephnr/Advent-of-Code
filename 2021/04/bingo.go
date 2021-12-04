package main

import (
	"regexp"
	"strconv"
)

type BingoMark struct {
	Value  int
	Marked bool
}

type BingoBoard struct {
	board [][]*BingoMark
}

// ParseBoards reads a series of lines to parse the bingo boards
func ParseBoards(lines []string) []*BingoBoard {
	bingoBoards := make([]*BingoBoard, 0)
	board := &BingoBoard{board: make([][]*BingoMark, boardSize)}
	rowCount := 0
	re := regexp.MustCompile(`(\d+)`)

	for _, line := range lines {
		if len(line) > 0 {
			for _, num := range re.FindAllString(line, -1) {
				val, _ := strconv.Atoi(num)
				board.board[rowCount] = append(board.board[rowCount], &BingoMark{val, false})
			}
		} else {
			bingoBoards = append(bingoBoards, board)
			board = &BingoBoard{board: make([][]*BingoMark, boardSize)}
			rowCount = 0
			continue
		}

		rowCount++
	}

	bingoBoards = append(bingoBoards, board)

	return bingoBoards
}

func (board BingoBoard) HasBingo() (bingo bool) {
	// Check horizontals
	for _, row := range board.board {
		markCount := 0
		for _, mark := range row {
			if mark.Marked {
				markCount++
			}
		}

		if markCount == boardSize {
			return true
		}
	}

	// Check verticals
	for i := 0; i < boardSize; i++ {
		markCount := 0
		for _, row := range board.board {

			if row[i].Marked {
				markCount++
			}
		}

		if markCount == boardSize {
			return true
		}
	}

	return bingo
}

func (board BingoBoard) Mark(num int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if board.board[i][j].Value == num {
				board.board[i][j].Marked = true
			}
		}
	}
}

func (board BingoBoard) Score(lastCallout int) (score int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if !board.board[i][j].Marked {
				score += board.board[i][j].Value
			}
		}
	}

	return score * lastCallout
}
