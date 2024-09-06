package pieces

import c "github.com/kelbwah/go-chess/constants"

type Knight struct {
	Type     c.PieceType
	Location *PieceLocation
	Color    c.PieceColor
}

func (k *Knight) GetCurrentLocation() *PieceLocation {
	return k.Location
}

func (k *Knight) SetCurrentLocation(newLocation *PieceLocation) {
	k.Location = newLocation
}

func (k *Knight) GetType() c.PieceType {
	return k.Type
}

func (k *Knight) SetType(newType c.PieceType) {
	k.Type = newType
}

func (k *Knight) GetColor() c.PieceColor {
	return k.Color
}

func (k *Knight) SetColor(color c.PieceColor) {
	k.Color = color
}

func (k *Knight) IsValidMove(newLocation PieceLocation) {
	panic("implement these thangs")
}
