package utils

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	// b "github.com/kelbwah/go-chess/board"
	c "github.com/kelbwah/go-chess/constants"
)

func InitDrawLoop() {
	rl.InitWindow(c.WIDTH, c.HEIGHT, "go-chess")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)

		DrawCheckerBoard()

		rl.EndDrawing()
	}
}
