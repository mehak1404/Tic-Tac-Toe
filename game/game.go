package game

import (
	"fmt"
	"tictactoe/board"
	"tictactoe/player"
)

type Game struct {
	Board          board.Board
	Player1        player.Player
	Player2        player.Player
	Current_player player.Player
}

func NewGame(game_mode int) *Game {

	b := board.InitializeBoard()
	var Player1, Player2 player.Player

	Player1 = player.NewHumanPlayer("X")

	if game_mode == 1 {
		Player2 = player.NewAIPlayer("O")
	} else {
		Player2 = player.NewHumanPlayer("O")
	}

	return &Game{
		Board:          b,
		Player1:        Player1,
		Player2:        Player2,
		Current_player: Player1,
	}
}

// start

func (g *Game) Start() {
	for !g.IsGameOver() {
		board.DisplayBoard(g.Board)
		position := g.Current_player.GetMove(&g.Board)
		g.Board = board.MakeMove(g.Board, position, g.Current_player.GetSymbol())
		g.SwitchPlayer()
	}
}

// switch player
func (g *Game) SwitchPlayer() {
	if g.Current_player == g.Player1 {
		g.Current_player = g.Player2
	} else {
		g.Current_player = g.Player1
	}
}

// end game
func (g *Game) IsGameOver() bool {

	player, is_winner := board.CheckWinner(g.Board)
	if is_winner {
		board.DisplayBoard(g.Board)
		fmt.Printf("Player %s wins!\n", player)
		return true
	}
	is_board_full := board.IsBoardFull(g.Board)
	if is_board_full {
		board.DisplayBoard(g.Board)
		fmt.Println("Game is a draw!")
		return true
	}
	return false

}
