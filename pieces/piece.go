package pieces

import c "github.com/kelbwah/go-chess/constants"

type PieceLocation struct {
	Row int
	Col int
}

type Piece interface {
	SetCurrentLocation(newLocation *PieceLocation)
	GetCurrentLocation() *PieceLocation
	GetType() c.PieceType
	SetType(c.PieceType)
	GetColor() c.PieceColor
	SetColor(color c.PieceColor)
}
