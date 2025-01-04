package test

import (
	"testing"
	"tictactoe/board"
)

func TestInitializeBoard(t *testing.T) {
	b := board.InitializeBoard()
	if len(b.Cells) != 3 {
		t.Errorf("Expected board to have 3 rows, got %d", len(b.Cells))
	}
	for i := 0; i < 3; i++ {
		if len(b.Cells[i]) != 3 {
			t.Errorf("Expected row to have 3 Cells, got %d", len(b.Cells[i]))
		}
		for j := 0; j < 3; j++ {
			if b.Cells[i][j] != "." {
				t.Errorf("Expected cell to be empty, got %s", b.Cells[i][j])
			}
		}
	}
}

func TestDisplayBoard(t *testing.T) {
	b := board.InitializeBoard()
	board.DisplayBoard(b)
}

func TestIsValidMove(t *testing.T) {

	b := board.InitializeBoard()
	if !board.IsValidMove(b, 0) {
		t.Errorf("Expected move to be valid, got invalid")
	}
	b = board.MakeMove(b, 0, "X")
	if board.IsValidMove(b, 0) {
		t.Errorf("Expected move to be invalid, got valid")
	}
	if board.IsValidMove(b, 9) {
		t.Errorf("Expected move to be invalid, got valid")
	}
}

func TestCheckWinner(t *testing.T) {
	b := board.InitializeBoard()
	b = board.MakeMove(b, 0, "X")
	b = board.MakeMove(b, 1, "X")
	b = board.MakeMove(b, 2, "X")
	player, is_winner := board.CheckWinner(b)
	if !is_winner {
		t.Errorf("Expected winner, got no winner")
	}
	if player != "X" {
		t.Errorf("Expected player X to win, got %s", player)
	}
	// test for column
	b = board.InitializeBoard()
	b = board.MakeMove(b, 0, "O")
	b = board.MakeMove(b, 3, "O")
	b = board.MakeMove(b, 6, "O")
	player, is_winner = board.CheckWinner(b)
	if !is_winner {
		t.Errorf("Expected winner, got no winner")
	}
	if player != "O" {
		t.Errorf("Expected player O to win, got %s", player)
	}
	// test for diagonal
	b = board.InitializeBoard()
	b = board.MakeMove(b, 0, "X")
	b = board.MakeMove(b, 4, "X")
	b = board.MakeMove(b, 8, "X")
	player, is_winner = board.CheckWinner(b)
	if !is_winner {
		t.Errorf("Expected winner, got no winner")
	}
	if player != "X" {
		t.Errorf("Expected player X to win, got %s", player)
	}
}

func TestMakeMove(t *testing.T) {
	b := board.InitializeBoard()
	b = board.MakeMove(b, 0, "X")
	if b.Cells[0][0] != "X" {
		t.Errorf("Expected cell to be X, got %s", b.Cells[0][0])
	}
}

func TestDrawConditions(t *testing.T) {
	b := board.InitializeBoard()
	b = board.MakeMove(b, 0, "X")
	b = board.MakeMove(b, 1, "O")
	b = board.MakeMove(b, 2, "X")
	b = board.MakeMove(b, 3, "O")
	b = board.MakeMove(b, 4, "X")
	b = board.MakeMove(b, 5, "O")
	b = board.MakeMove(b, 6, "O")
	b = board.MakeMove(b, 7, "X")
	b = board.MakeMove(b, 8, "O")
	player, is_winner := board.CheckWinner(b)
	if is_winner {
		t.Errorf("Expected no winner, got winner")
	}
	if player != "" {
		t.Errorf("Expected no player to win, got %s", player)
	}
}
