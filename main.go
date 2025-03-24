package main

import (
	"CHIP8/Chip8"
	r1 "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

/**
* main function handling all Chip8 functionality
*/

func main() {
	chip8 := Chip8.GetChip8()
	chip8.Initialise()
//	chip8.FetchOpcode()
	screenWidth := 64 * int(chip8.PixelWidth)
	screenHeight := 32 * int(chip8.PixelHeight)
	
	fmt.Println(chip8)
	r1.InitWindow(int32(screenWidth * chip8.Scale), int32(screenHeight * chip8.Scale), "Md's Chip8 Emulator")
	r1.SetTargetFPS(60)

	for !r1.WindowShouldClose() {
		r1.BeginDrawing()
		r1.ClearBackground(r1.Black)
		chip8.DrawDisplay()

		r1.EndDrawing()
	}
	r1.CloseWindow()

}
