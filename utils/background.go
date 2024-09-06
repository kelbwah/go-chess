package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	c "github.com/kelbwah/go-chess/constants"
)

func DrawCheckerBoard() {
	for row := 0; row < c.BOARD_SIZE; row++ {
		for col := 0; col < c.BOARD_SIZE; col++ {
			x := col * c.SQUARE_X
			y := row * c.SQUARE_Y

			if ((row + col) % 2) != 0 {
				rl.DrawRectangle(int32(x), int32(y), int32(c.SQUARE_X), int32(c.SQUARE_Y), c.DARK_COLOR)
			} else {
				rl.DrawRectangle(int32(x), int32(y), int32(c.SQUARE_X), int32(c.SQUARE_Y), c.LIGHT_COLOR)
			}
		}
	}
}
