package rules

const (
	XSymbol         = 'X'
	OSymbol         = 'O'
	EmptyCellSymbol = '.'

	N = 3
)

type Game struct {
	board          [][]byte
	nextTurnPlayer byte
}

func NewGame() *Game {
	board := make([][]byte, N)
	for i := 0; i < N; i++ {
		board[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			board[i][j] = EmptyCellSymbol
		}
	}

	return &Game{
		board:          board,
		nextTurnPlayer: XSymbol,
	}
}

func (g *Game) PlayMove(i, j int) {
	// TODO: implement
}

func (g *Game) Winner() {
	// TODO: implement
}
