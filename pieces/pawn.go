package pieces

import c "github.com/kelbwah/go-chess/constants"

type Pawn struct {
	Type     c.PieceType
	Location *PieceLocation
	Color    c.PieceColor
}

func (p *Pawn) SetCurrentLocation(newLocation *PieceLocation) {
	p.Location = newLocation
}

func (p *Pawn) GetCurrentLocation() *PieceLocation {
	return p.Location
}

func (p *Pawn) GetType() c.PieceType {
	return p.Type
}

func (p *Pawn) SetType(newType c.PieceType) {
	p.Type = newType
}

func (p *Pawn) GetColor() c.PieceColor {
	return p.Color
}

func (p *Pawn) SetColor(color c.PieceColor) {
	p.Color = color
}
