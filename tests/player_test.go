package test;

import (
	"testing"
	"tictactoe/board"
	"tictactoe/player"
)

func TestNewHumanPlayer(t *testing.T) {
	p := player.NewHumanPlayer("X")
	if p.GetSymbol() != "X" {
		t.Errorf("Expected symbol to be X, got %s", p.GetSymbol())
	}
}
func TestNewAIPlayer(t *testing.T) {
	p := player.NewAIPlayer("O")
	if p.GetSymbol() != "O" {
		t.Errorf("Expected symbol to be O, got %s", p.GetSymbol())
	}
}

func TestGetMove(t *testing.T) {
	p := player.NewHumanPlayer("X")
	b := board.InitializeBoard()
	position := p.GetMove(&b)
	if position < 0 || position > 8 {
		t.Errorf("Expected position between 0 and 8, got %d", position)
	}
}

func TestAIWinningMove(t *testing.T) {
	p := player.NewAIPlayer("O")
	b := board.InitializeBoard()
	b = board.MakeMove(b, 0, "O")
	b = board.MakeMove(b, 1, "O")
	position := p.GetMove(&b)
	if position != 2 {
		t.Errorf("Expected position 2, got %d", position)
	}
}

func TestAIBlockingMove(t *testing.T) {
	p := player.NewAIPlayer("O")
	b := board.InitializeBoard()
	b = board.MakeMove(b, 0, "X")
	b = board.MakeMove(b, 1, "X")
	position := p.GetMove(&b)
	if position != 2 {
		t.Errorf("Expected position 2, got %d", position)
	}
}

func TestAIStrategyMove(t *testing.T) {
	p := player.NewAIPlayer("O")
	b := board.InitializeBoard()
	b = board.MakeMove(b, 0, "X")
	b = board.MakeMove(b, 1, "O")
	b = board.MakeMove(b, 2, "X")
	b = board.MakeMove(b, 3, "O")
	b = board.MakeMove(b, 4, "X")
	position := p.GetMove(&b)
	if position != 6 {
		t.Errorf("Expected position 6, got %d", position)
	}
}

