package player

import c "github.com/kelbwah/go-chess/constants"

type Player struct {
	Username      string
	Color         c.PieceColor
	IsPlayerReady bool
}

func (p *Player) GetUsername() string {
	return p.Username
}

func (p *Player) SetUsername(newUsername string) {
	p.Username = newUsername
}

func (p *Player) GetColor() c.PieceColor {
	return p.Color
}

func (p *Player) SetColor(newColor c.PieceColor) {
	p.Color = newColor
}

func (p *Player) IsReady() bool {
	return p.IsPlayerReady
}

func (p *Player) SetIsReady(isReady bool) {
	p.IsPlayerReady = isReady
}
