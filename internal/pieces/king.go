package pieces

import c "github.com/kelbwah/go-chess/constants"

type King struct {
	Type     c.PieceType
	Location *PieceLocation
	Color    c.PieceColor
}

func (k *King) GetCurrentLocation() *PieceLocation {
	return k.Location
}

func (k *King) SetCurrentLocation(newLocation *PieceLocation) {
	k.Location = newLocation
}

func (k *King) GetType() c.PieceType {
	return k.Type
}

func (k *King) SetType(newType c.PieceType) {
	k.Type = newType
}

func (k *King) GetColor() c.PieceColor {
	return k.Color
}

func (k *King) SetColor(color c.PieceColor) {
	k.Color = color
}

func (k *King) IsValidMove(newLocation PieceLocation) {
	panic("implement these thangs")
}
