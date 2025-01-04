package board

type Board struct {
	Cells [3][3]string
}

func InitializeBoard() Board {
	var board Board
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board.Cells[i][j] = "."
		}
	}
	return board
}

func DisplayBoard(board Board) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(board.Cells[i][j])
			if j < 2 {
				print(" | ")
			}
		}
		print("\n")
		if i < 2 {
			print("----------\n")
		}

	}
}

func IsValidMove(board Board, position int) bool {
	row := position / 3
	col := position % 3

	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	return board.Cells[row][col] == "."
}

func MakeMove(board Board, position int, player string) Board {
	row := position / 3
	col := position % 3

	board.Cells[row][col] = player
	return board
}

func CheckWinner(board Board) (string, bool) {
	for i := 0; i < 3; i++ {
		if board.Cells[i][0] != "." && board.Cells[i][0] == board.Cells[i][1] && board.Cells[i][1] == board.Cells[i][2] {
			return board.Cells[i][0], true
		}
		if board.Cells[0][i] != "." && board.Cells[0][i] == board.Cells[1][i] && board.Cells[1][i] == board.Cells[2][i] {
			return board.Cells[0][i], true
		}

		if board.Cells[0][0] != "." && board.Cells[0][0] == board.Cells[1][1] && board.Cells[1][1] == board.Cells[2][2] {
			return board.Cells[0][0], true
		}
		if board.Cells[0][2] != "." && board.Cells[0][2] == board.Cells[1][1] && board.Cells[1][1] == board.Cells[2][0] {
			return board.Cells[0][2], true
		}
	}
	return "", false
}

func IsBoardFull(board Board) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board.Cells[i][j] == "." {
				return false
			}
		}
	}
	return true
}

func RemoveMove(board Board, position int) Board {
	row := position / 3
	col := position % 3

	board.Cells[row][col] = "."
	return board
}
