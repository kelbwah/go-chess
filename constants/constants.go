package constants

import "image/color"

type Runtime string

const (
	SERVER Runtime = "server"
	GAME           = "game"
)

// Various piece and game types
type (
	PieceColor string
	PieceType  string
	GameType   string
)

// Piece color constants
const (
	WHITE PieceColor = "white"
	BLACK            = "black"
)

// Piece type constants
const (
	PAWN   PieceType = "pawn"
	KNIGHT           = "knight"
	BISHOP           = "bishop"
	ROOK             = "rook"
	QUEEN            = "queen"
	KING             = "king"
)

// Game type constants
const (
	BULLET GameType = "bullet"
	BLITZ  GameType = "blitz"
	RAPID  GameType = "standard"
)

// Map to hold times for different game types in milliseconds
var GameToTime = map[GameType]float64{
	BULLET: 60000,
	BLITZ:  180000,
	RAPID:  600000,
}

// Raylib window constants
const (
	WIDTH      = 800
	HEIGHT     = 800
	BOARD_SIZE = 8
	SQUARE_X   = WIDTH / 8
	SQUARE_Y   = HEIGHT / 8
)

// Raylib color vars
var (
	LIGHT_COLOR = color.RGBA{240, 236, 212, 255}
	DARK_COLOR  = color.RGBA{120, 148, 84, 255}
)
