package game

import (
	b "github.com/kelbwah/go-chess/board"
)

type Game struct {
	Board *b.Board
}

func CreateGame() *Game {
	return &Game{
		Board: b.CreateBoard(),
	}
}

func (g *Game) Start() {
}
