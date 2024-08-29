package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	b "github.com/kelbwah/go-chess/board"
)

const (
	WIDTH  = 1280
	HEIGHT = 720
)

func main() {
	board := b.CreateBoard()
	board.InitBoardState()
	board.Print()

	rl.InitWindow(WIDTH, HEIGHT, "go-chess")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("go-chess", ((WIDTH / 2) - 100), ((HEIGHT / 2) - 100), 20, rl.LightGray)

		rl.EndDrawing()
	}
}
