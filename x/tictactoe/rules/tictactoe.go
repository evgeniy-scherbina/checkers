package rules

import "encoding/json"

const (
	XSymbol         = 'X'
	OSymbol         = 'O'
	EmptyCellSymbol = '.'

	N = 3
)

type Game struct {
	Board    [][]byte `json:"board"`
	NextTurn byte     `json:"next_turn"`
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
		Board:    board,
		NextTurn: XSymbol,
	}
}

func (g *Game) Serialize() ([]byte, error) {
	return json.Marshal(g)
}

func (g *Game) MustSerialize() []byte {
	data, err := g.Serialize()
	if err != nil {
		panic(err)
	}

	return data
}

func Deserialize(data []byte) (*Game, error) {
	var g Game
	if err := json.Unmarshal(data, &g); err != nil {
		return nil, err
	}

	return &g, nil
}

//func (g *Game) NextTurn() byte {
//	return g.nextTurn
//}

func (g *Game) PlayMove(i, j int) {
	// TODO: implement
}

func (g *Game) Winner() {
	// TODO: implement
}
