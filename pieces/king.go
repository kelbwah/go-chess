package pieces

import "github.com/kelbwah/go-chess/constants"

type King struct {
	Type     constants.PieceType
	Location *PieceLocation
	Color    constants.PieceColor
}

func (k *King) GetCurrentLocation() *PieceLocation {
	return k.Location
}

func (k *King) SetCurrentLocation(newLocation *PieceLocation) {
	k.Location = newLocation
}

func (k *King) GetType() constants.PieceType {
	return k.Type
}

func (k *King) SetType(newType constants.PieceType) {
	k.Type = newType
}

func (k *King) GetColor() constants.PieceColor {
	return k.Color
}

func (k *King) SetColor(color constants.PieceColor) {
	k.Color = color
}
