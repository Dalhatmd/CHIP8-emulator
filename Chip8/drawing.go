package Chip8 

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
)


func (c *Chip8) DrawDisplay() *Chip8 {

	chip8Height := len(c.Gfx)
	chip8Width := len(c.Gfx[0])

	pixelSize := int(c.PixelSize)

	for y := 0; y < chip8Height; y++ {
		for x:= 0; x < chip8Width; x++ {
			if c.Gfx[y][x] != 0 {
				r1.DrawRectangle(int32(x*pixelSize), int32(y*pixelSize), int32(pixelSize), int32(pixelSize), r1.White)
			}
		}
	}
	return c
}
