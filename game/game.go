package game

import (
	"fmt"
	"tictactoe/board"
	"tictactoe/player"
)
type Game struct{
	board board.Board
	player1 player.Player
	player2 player.Player
	current_player player.Player
}

func NewGame(game_mode int) *Game{

	b := board.InitializeBoard()
	var player1, player2 player.Player

	player1 = player.NewHumanPlayer("X")

	if game_mode == 1{
		player2 = player.NewAiPlayer("O")
	}else{
		player2 = player.NewHumanPlayer("O")
	}

	return &Game{
		board: b,
		player1: player1,
		player2: player2,
		current_player: player1,
	}
}


// start

func (g * Game) Start(){
	for ! g.IsGameOver(){
		board.DisplayBoard(g.board)
		position := g.current_player.GetMove(&g.board)
		g.board = board.MakeMove(g.board, position, g.current_player.GetSymbol())
		g.SwitchPlayer()
	}
}
// switch player
func (g * Game) SwitchPlayer(){
	if g.current_player == g.player1{
		g.current_player = g.player2
	}else{
		g.current_player = g.player1
	}
}
// end game
func (g * Game) IsGameOver() bool{
	
	player, is_winner := board.CheckWinner(g.board)
	if is_winner {
		board.DisplayBoard(g.board)
		fmt.Printf("Player %s wins!\n", player)
		return true
	}
	is_board_full := board.IsBoardFull(g.board)
	if is_board_full {
		board.DisplayBoard(g.board)
		fmt.Println("Game is a draw!")
		return true
	}
	return false
	
}