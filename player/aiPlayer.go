package player

import (
	"fmt"
	"tictactoe/board"
)

type AIPlayer struct {
	symbol string
}

func NewAIPlayer(symbol string) *AIPlayer {
	return &AIPlayer{symbol: symbol}
}

func (p *AIPlayer) GetSymbol() string {
	return p.symbol
}

func (p *AIPlayer) GetMove(b *board.Board) int {

	tempBoard := *b
	// checking for the winning move
	ai_symbol := p.GetSymbol()
	for position := 0; position < 9; position++ {
		if board.IsValidMove(*b, position) {
			tempBoard = board.MakeMove(tempBoard, position, ai_symbol)
			_, is_winner := board.CheckWinner(tempBoard)
			tempBoard = board.MakeMove(tempBoard, position, ".")
			if is_winner {
				board.MakeMove(*b, position, ai_symbol)
				fmt.Printf("AI made a winning move to cell %d.\n", position)
				return position
			}
		}
	}

	// check for blocking move, here we will check
	// for opponent wining move and block it
	opponent_symbol := "X"
	if ai_symbol == "O" {
		opponent_symbol = "X"
	} else {
		opponent_symbol = "O"
	}
	for position := 0; position < 9; position++ {
		if board.IsValidMove(*b, position) {
			tempBoard = board.MakeMove(tempBoard, position, opponent_symbol)
			_, is_winner := board.CheckWinner(tempBoard)
			tempBoard = board.MakeMove(tempBoard, position, ".")
			if is_winner {
				board.MakeMove(*b, position, ai_symbol)
				fmt.Printf("Ai made a blocking move to cell %d \n", position)
				return position
			}
		}
	}

	// at last , we will make a strategic move
	preferredMoves := []int{4, 0, 2, 6, 8, 1, 3, 5, 7}
	for _, position := range preferredMoves {
		if board.IsValidMove(*b, position) {
			board.MakeMove(*b, position, p.symbol)
			fmt.Printf("AI choose strategic position %d\n", position)
			return position
		}
	}

	return -1
}
