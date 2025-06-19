package Chip8 

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
)


func (c *Chip8) DrawDisplay() *Chip8 {

	for y := range c.Gfx { 
		for x:= range c.Gfx[y] {
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
