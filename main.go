package main

import (
	"CHIP8/Chip8"
	"fmt"
	r1 "github.com/gen2brain/raylib-go/raylib"
)

/**
* main function handling all Chip8 functionality
*/

func main() {
	chip8 := Chip8.GetChip8()

	screenWidth := 64 * int(chip8.PixelSize)
	screenHeight := 32 * int(chip8.PixelSize)

	r1.InitWindow(int32(screenWidth), int32(screenHeight), "Md's Chip8 Emulator")
	r1.SetTargetFPS(60)

	for !r1.WindowShouldClose() {
		r1.BeginDrawing()
		r1.ClearBackground(r1.Black)
		chip8.DrawDisplay()

		r1.EndDrawing()
	}
	r1.CloseWindow()

}
