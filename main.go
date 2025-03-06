package main
import (
	"CHIP8/drawing"
	"CHIP8/config"
	r1 "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	r1.InitWindow(config.ScreenWidth, config.ScreenHeight, "CHIP8")
	defer r1.CloseWindow()

	r1.SetTargetFPS(60)

	for !r1.WindowShouldClose() {
		r1.BeginDrawing()
		r1.ClearBackground(r1.Black)

		drawing.DrawDisplay()
		r1.EndDrawing()
	}
}
