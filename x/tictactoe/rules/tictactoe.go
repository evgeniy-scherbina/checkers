package rules

const (
	XSymbol         = 'X'
	OSymbol         = 'O'
	EmptyCellSymbol = '.'

	N = 3
)

type Game struct {
	board    [][]byte
	nextTurn byte

	xPlayer string
	yPlayer string
}

func NewGame(xPlayer, yPlayer string) *Game {
	board := make([][]byte, N)
	for i := 0; i < N; i++ {
		board[i] = make([]byte, N)
		for j := 0; j < N; j++ {
			board[i][j] = EmptyCellSymbol
		}
	}

	return &Game{
		board:    board,
		nextTurn: XSymbol,
		xPlayer:  xPlayer,
		yPlayer:  yPlayer,
	}
}

func (g *Game) PlayMove(i, j int) {
	// TODO: implement
}

func (g *Game) Winner() {
	// TODO: implement
}
