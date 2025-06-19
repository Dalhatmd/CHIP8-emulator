package Chip8 

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
)


func (c *Chip8) DrawDisplay() *Chip8 {

	chip8Height := len(c.Gfx)
	chip8Width := len(c.Gfx[0])

	for y := 0; y < chip8Height; y++ {
		for x:= 0; x < chip8Width; x++ {
			if c.Gfx[y][x] {
				drawX := int32(x) * int32(c.Scale)
				drawY := int32(y) * int32(c.Scale)

				r1.DrawRectangle(drawX, drawY, int32(c.Scale), int32(c.Scale), r1.White)
			}
		}
	}
	return c
}

func (c *Chip8) ClearDisplay() {
	for y:= range c.Gfx {
		for x := range c.Gfx[y] {
			c.Gfx[y][x] = false
		}
	}
}
