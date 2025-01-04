package player
import(
	"ticTacToe/board/board"
	"fmt"
)
type Player interface {
	GetMove(b * board.Board ) int
	GetSymbol() string
}

type HumanPlayer struct{
	symbol string
}

func newHumanPlayer(symbol string) * HumanPlayer{
	return &HumanPlayer{symbol: symbol}
}
func (p * HumanPlayer) GetSymbol() string {
	return p.symbol
}

func (p * HumanPlayer) GetMove(b * board.Board){
	var position int
	for{
		fmt.Println("Enter a position (0 - 8): ")
		fmt.Scan(&position)
		if board.isValidMove(*b, position){
			board.makeMove(*b, position, p.symbol)
			fmt.Println("Move added to cell %d.", position)
			break
		}else{
			fmt.Println("Invalid move. Try again.")
		}

	}
}
