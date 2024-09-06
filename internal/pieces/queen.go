package pieces

import c "github.com/kelbwah/go-chess/constants"

type Queen struct {
	Type     c.PieceType
	Location *PieceLocation
	Color    c.PieceColor
}

func (q *Queen) GetCurrentLocation() *PieceLocation {
	return q.Location
}

func (q *Queen) SetCurrentLocation(newLocation *PieceLocation) {
	q.Location = newLocation
}

func (q *Queen) GetType() c.PieceType {
	return q.Type
}

func (q *Queen) SetType(newType c.PieceType) {
	q.Type = newType
}

func (q *Queen) GetColor() c.PieceColor {
	return q.Color
}

func (q *Queen) SetColor(color c.PieceColor) {
	q.Color = color
}

func (q *Queen) IsValidMove(newLocation PieceLocation) {
	panic("implement these thangs")
}
