package board;

type Board struct{
	cells [3][3] string

}

func initializeBoard() Board{
	var board Board;
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.cells[i][j] = ".";
		}
	}
	return board;
}

func displayBoard(board Board){
	for i:= 0; i < 3; i++ {
		for j:=0; j < 3; j++ {
			print(board.cells[i][j]);
			if j < 2 {
				print(" | ");
			}
		}
		print("\n");
		if i < 2 {
			print("----------\n");
		}
		
	}
} 

func isValidMove(board Board, position int) bool{
	row := position / 3;
	col := position % 3;

	if row < 0 || row > 2 || col < 0 || col > 2{
		return false
	}
	return board.cells[row][col] == ".";
}

func makeMove(board Board, position int, player string) Board{
	row := position / 3;
	col := position % 3;

	board.cells[row][col] = player;
	return board;
}

func checkWinner(board Board) (string, bool){
	for i:= 0; i < 3; i++{
		if board.cells[i][0] != "." && board.cells[i][0] == board.cells[i][1] && board.cells[i][1] == board.cells[i][2]{
			return board.cells[i][0], true
		}
		if board.cells[0][i] != "." && board.cells[0][i] == board.cells[1][i] && board.cells[1][i] == board.cells[2][i]{
			return board.cells[0][i], true
		}

		if board.cells[0][0] != "." && board.cells[0][0] == board.cells[1][1] && board.cells[1][1] == board.cells[2][2]{
			return board.cells[0][0], true
		}
		if board.cells[0][2] != "." && board.cells[0][2] == board.cells[1][1] && board.cells[1][1] == board.cells[2][0]{
			return board.cells[0][2], true
		}	
	}
	return "", false
}

func isBoardFull(board Board) bool{
	for i:= 0; i < 3; i++{
		for j:= 0; j < 3; j++{
			if board.cells[i][j] == "."{
				return false
			}
		}
	}
	return true
}

func removeMove(board Board, position int) Board{
	row := position / 3;
	col := position % 3;

	board.cells[row][col] = ".";
	return board;
}
