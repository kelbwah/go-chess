package pieces

import c "github.com/kelbwah/go-chess/constants"

type Bishop struct {
	Type     c.PieceType
	Location *PieceLocation
	Color    c.PieceColor
}

func (b *Bishop) GetCurrentLocation() *PieceLocation {
	return b.Location
}

func (b *Bishop) SetCurrentLocation(newLocation *PieceLocation) {
	b.Location = newLocation
}

func (b *Bishop) GetType() c.PieceType {
	return b.Type
}

func (b *Bishop) SetType(newType c.PieceType) {
	b.Type = newType
}

func (b *Bishop) GetColor() c.PieceColor {
	return b.Color
}

func (b *Bishop) SetColor(color c.PieceColor) {
	b.Color = color
}
