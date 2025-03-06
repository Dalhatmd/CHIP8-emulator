package drawing

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
	"CHIP8/config"
)


var display [config.Chip8Width][config.Chip8Height]bool

func DrawDisplay() {
	for y := 0; y < config.Chip8Height; y++ {
		for x:= 0; x < config.Chip8Width; x++ {
			if display[x][y] {
				r1.DrawRectangle(int32(x*config.PixelSize), int32(y*config.PixelSize), config.PixelSize, config.PixelSize, r1.White)
			}
		}
	}
}
