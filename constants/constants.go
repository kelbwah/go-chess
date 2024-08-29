package constants

type (
	PieceColor string
	PieceType  string
)

const (
	WHITE PieceColor = "white"
	BLACK            = "black"
)

const (
	PAWN   PieceType = "pawn"
	KNIGHT           = "knight"
	BISHOP           = "bishop"
	ROOK             = "rook"
	QUEEN            = "queen"
	KING             = "king"
)
