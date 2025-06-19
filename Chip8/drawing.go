package Chip8 

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
)


func (c *Chip8) DrawDisplay() *Chip8 {

	chip8Height := len(c.Gfx)
	chip8Width := len(c.Gfx[0])

	for y := 0; y < chip8Height/4; y++ {
		for x:= 0; x < chip8Width/4; x++ {
			if c.Gfx[y][x] {
				drawX := int32(x)
				drawY := int32(y)

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
