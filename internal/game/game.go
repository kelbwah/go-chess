package game

import (
	c "github.com/kelbwah/go-chess/constants"
	b "github.com/kelbwah/go-chess/internal/board"
	p "github.com/kelbwah/go-chess/internal/player"
)

type Game struct {
	Board       *b.Board
	WhitePlayer *p.Player
	BlackPlayer *p.Player
	WhiteTime   float64
	BlackTime   float64
	IsPaused    bool
	IsEnded     bool
	CurrentTurn c.PieceColor
	CanStart    bool
	GameType    c.GameType
}

func CreateGame(gameType c.GameType, whitePlayer *p.Player, blackPlayer *p.Player) *Game {
	return &Game{
		Board:       b.CreateBoard(),
		WhiteTime:   c.GameToTime[gameType],
		WhitePlayer: whitePlayer,
		BlackTime:   c.GameToTime[gameType],
		BlackPlayer: blackPlayer,
		IsPaused:    true,
		IsEnded:     false,
		CanStart:    false,
		CurrentTurn: c.WHITE,
	}
}

func (g *Game) Start() {
}

func (g *Game) Reset() {
}
