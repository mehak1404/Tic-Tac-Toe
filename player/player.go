package player

import (
	"fmt"
	"tictactoe/board"
)

type Player interface {
	GetMove(b *board.Board) int
	GetSymbol() string
}

type HumanPlayer struct {
	symbol string
}

func NewHumanPlayer(symbol string) *HumanPlayer {
	return &HumanPlayer{symbol: symbol}
}
func (p *HumanPlayer) GetSymbol() string {
	return p.symbol
}

func (p *HumanPlayer) GetMove(b *board.Board) int {
	var position int
	for {
		fmt.Println("Enter a position (0 - 8): ")
		fmt.Scan(&position)
		if board.IsValidMove(*b, position) {
			board.MakeMove(*b, position, p.symbol)
			fmt.Printf("Player Move added to cell %d. \n", position)
			break
		} else {
			fmt.Println("Invalid move. Try again.")
		}

	}
	return position
}
