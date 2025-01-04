package main

import (
	"tictactoe/game"
	"fmt"
)
func main(){

	fmt.Println("Welcome to Tic Tac Toe")
	fmt.Println("Choose game mode: ")
	fmt.Println("1. Player vs AI")
	fmt.Println("2. Player vs Player")
	var game_mode int
	fmt.Scan(&game_mode)
	g := game.NewGame(game_mode)
	g.Start()

	fmt.Println("Would you like to play again? (y/n)")
	var play_again string
	fmt.Scan(&play_again)
	if play_again == "y" || play_again == "Y"{
		main()
	}else{
		fmt.Println("Thanks for playing!")
	}
	  
}