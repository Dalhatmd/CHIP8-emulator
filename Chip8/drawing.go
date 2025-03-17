package Chip8 

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
)


func (c *Chip8) DrawDisplay() *Chip8 {

	chip8Height := len(c.Gfx)
	chip8Width := len(c.Gfx[0])

	pixelWidth := int(c.PixelWidth)
	pixelHeight := int(c.PixelHeight)

	for y := 0; y < chip8Height; y++ {
		for x:= 0; x < chip8Width; x++ {
			if c.Gfx[y][x] {
				r1.DrawRectangle(int32(x * c.Scale), int32(y * c.Scale), int32(pixelWidth), int32(pixelHeight), r1.White)
			}
		}
	}
	return c
}
