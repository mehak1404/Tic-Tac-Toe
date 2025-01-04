package test

import (
	"testing"
	"tictactoe/board"
	"tictactoe/game"
)

func TestNewGame(t *testing.T) {
	g := game.NewGame(1)
	if g == nil {
		t.Errorf("Expected game to be created, got nil")
	}

}

func TestSwitchPlayer(t *testing.T) {
	g := game.NewGame(1)
	g.SwitchPlayer()
	if g.Current_player == g.Player1 {
		t.Errorf("Expected player 2, got player 1")
	}
}

func TestIsGameOver(t *testing.T) {
	g := game.NewGame(1)
	if g.IsGameOver() {
		t.Errorf("Expected game to not be over, got over")
	}
	// Winning row
	g.Board = board.MakeMove(g.Board, 0, "X")
	g.Board = board.MakeMove(g.Board, 1, "X")
	g.Board = board.MakeMove(g.Board, 2, "X")
	if !g.IsGameOver() {
		t.Errorf("Expected game to be over, got not over")
	}

	// Draw
	g.Board = board.InitializeBoard()
	g.Board = board.MakeMove(g.Board, 0, "X")
	g.Board = board.MakeMove(g.Board, 1, "O")
	g.Board = board.MakeMove(g.Board, 2, "X")
	g.Board = board.MakeMove(g.Board, 3, "O")
	g.Board = board.MakeMove(g.Board, 4, "X")
	g.Board = board.MakeMove(g.Board, 5, "O")
	g.Board = board.MakeMove(g.Board, 6, "O")
	g.Board = board.MakeMove(g.Board, 7, "X")
	g.Board = board.MakeMove(g.Board, 8, "O")
	if !g.IsGameOver() {
		t.Errorf("Expected game to be over, got not over")
	}
}
