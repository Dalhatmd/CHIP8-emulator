package main

import (
	"CHIP8/Chip8"
	"fmt"
	"os"

	r1 "github.com/gen2brain/raylib-go/raylib"
)

/**
* main function handling all Chip8 functionality
 */

func main() {
	chip8 := Chip8.New()

	// Load a ROM file
	var romPath string
	if len(os.Args) > 1 {
		romPath = os.Args[1]
	} else {
		romPath = "roms/PONG"
	}

	rom, err := os.ReadFile(romPath)
	if err != nil {
		fmt.Println("An error ooccured while reading rom", err)
		os.Exit(1)
	}
	chip8.LoadRom(rom)
	fmt.Println("Rom Loaded successfully! Size:", len(rom), "bytes")

	screenWidth := 64 * int(chip8.PixelWidth)
	screenHeight := 32 * int(chip8.PixelHeight)

	fmt.Println(chip8)
	r1.InitWindow(int32(screenWidth*chip8.Scale), int32(screenHeight*chip8.Scale), "Md's Chip8 Emulator")
	r1.SetTargetFPS(60)

	for !r1.WindowShouldClose() {
		opcode := chip8.FetchOpcode()
		chip8.ExecuteOpcode(opcode)

		chip8.UpdateTimers()

		r1.BeginDrawing()
		r1.ClearBackground(r1.Black)
		chip8.DrawDisplay()
		r1.EndDrawing()
	}
	r1.CloseWindow()

}
