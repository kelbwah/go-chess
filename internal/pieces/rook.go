package pieces

import c "github.com/kelbwah/go-chess/constants"

type Rook struct {
	Type     c.PieceType
	Location *PieceLocation
	Color    c.PieceColor
}

func (r *Rook) GetCurrentLocation() *PieceLocation {
	return r.Location
}

func (r *Rook) SetCurrentLocation(newLocation *PieceLocation) {
	r.Location = newLocation
}

func (r *Rook) GetType() c.PieceType {
	return r.Type
}

func (r *Rook) SetType(newType c.PieceType) {
	r.Type = newType
}

func (r *Rook) GetColor() c.PieceColor {
	return r.Color
}

func (r *Rook) SetColor(color c.PieceColor) {
	r.Color = color
}

func (r *Rook) IsValidMove(newLocation PieceLocation) {
	panic("implement these thangs")
}
