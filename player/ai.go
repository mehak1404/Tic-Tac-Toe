package player

import (
	"ticTacToe/board/board"
	"fmt"
)

type AiPlayer struct {
	symbol string
}

func newAiPlayer(symbol string) *AiPlayer {
	return &AiPlayer{symbol: symbol}
}

func (p *AiPlayer) GetSymbol() string {
	return p.symbol
}

func (p *AiPlayer) GetMove(b *board.Board) {

	// checking for the winning move
	for position := 0; position < 9; position++ {
		if board.isValidMove(*b, position) {
			board.makeMove(*b, position, p.symbol)
			if board.checkWinner(*b) {
				fmt.Println("Move added to cell %d.", position)
				return
			} else {
				board.makeMove(*b, position, ".")
			}
		}
	}

	// check for blocking move, here we will check
	// for opponent wining move and block it
	opponent_symbol := "X"
	if p.symbol == "O" {
		opponent_symbol = "X"
	} else {
		opponent_symbol = "O"
	}
	for position := 0; position < 9; position++ {
		if board.isValidMove(*b, position) {
			board.makeMove(*b, position, opponent_symbol)
			if board.checkWinner(*b) {
				board.makeMove(*b, position, p.symbol)
				fmt.Println("Move added to cell %d.", position)
				return
			} else {
				board.makeMove(*b, position, ".")
			}
		}
	}

	// at last , we will make a random move
	for position := 0; position < 9; position++ {
		if board.isValidMove(*b, position) {
			board.makeMove(*b, position, p.symbol)
			fmt.Println("Move added to cell %d.", position)
			return
		}
	}
	return
}
