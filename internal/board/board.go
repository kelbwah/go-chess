package board

import (
	"fmt"

	c "github.com/kelbwah/go-chess/constants"
	p "github.com/kelbwah/go-chess/internal/pieces"
)

type Board struct {
	State [8][8]p.Piece
}

func CreateBoard() *Board {
	board := &Board{}
	board.InitBoardState()

	return board
}

func (b *Board) InitBoardState() {
	b.State = [8][8]p.Piece{}

	// Filling up pawns
	for i := 0; i < 8; i++ {
		pawnTopLocation := &p.PieceLocation{Row: 1, Col: i}
		pawnBottomLocation := &p.PieceLocation{Row: 6, Col: i}

		b.State[1][i] = &p.Pawn{
			Type:     c.PAWN,
			Color:    c.BLACK,
			Location: pawnTopLocation,
		}
		b.State[6][i] = &p.Pawn{
			Type:     c.PAWN,
			Color:    c.WHITE,
			Location: pawnBottomLocation,
		}
	}

	// Filling rest of the board
	piecesInOrder := [8]func() p.Piece{
		func() p.Piece { return &p.Rook{Type: c.ROOK} },
		func() p.Piece { return &p.Knight{Type: c.KNIGHT} },
		func() p.Piece { return &p.Bishop{Type: c.BISHOP} },
		func() p.Piece { return &p.Queen{Type: c.QUEEN} },
		func() p.Piece { return &p.King{Type: c.KING} },
		func() p.Piece { return &p.Bishop{Type: c.BISHOP} },
		func() p.Piece { return &p.Knight{Type: c.KNIGHT} },
		func() p.Piece { return &p.Rook{Type: c.ROOK} },
	}
	for i := 0; i < 8; i++ {
		topLocation := &p.PieceLocation{Row: 0, Col: i}
		bottomLocation := &p.PieceLocation{Row: 7, Col: i}

		b.State[0][i] = piecesInOrder[i]()
		b.State[0][i].SetCurrentLocation(topLocation)
		b.State[0][i].SetColor(c.BLACK)

		b.State[7][i] = piecesInOrder[i]()
		b.State[7][i].SetCurrentLocation(bottomLocation)
		b.State[7][i].SetColor(c.WHITE)
	}
}

func (b *Board) HasWinner() bool {
	return false
}

func (b *Board) Move(*p.Piece, *p.PieceLocation, *p.PieceLocation) bool {
	return false
}

func (b *Board) ResetBoard() {
	b.State = [8][8]p.Piece{}
}

func (b *Board) Print() {
	for i := 0; i < 8; i++ {
		fmt.Print("| ")
		for j := 0; j < 8; j++ {
			if b.State[i][j] != nil {
				fmt.Printf("%v:%v:%v", b.State[i][j].GetType(), b.State[i][j].GetColor(), &b.State[i][j])
			} else {
				fmt.Print("<nil>, ")
			}
		}
		fmt.Print(" |\n")
	}
}
