package Chip8

import (
	r1 "github.com/gen2brain/raylib-go/raylib"
)

func (c *Chip8) DrawDisplay() *Chip8 {
	for y := range c.Gfx {
		for x := range c.Gfx[y] {
			if c.Gfx[y][x] {
				drawX := int32(x) * int32(c.PixelWidth) * int32(c.Scale)
				drawY := int32(y) * int32(c.PixelHeight) * int32(c.Scale)
				width := int32(c.PixelWidth) * int32(c.Scale)
				height := int32(c.PixelHeight) * int32(c.Scale)
				r1.DrawRectangle(drawX, drawY, width, height, r1.White)
			}
		}
	}
	return c
}

func (c *Chip8) ClearDisplay() {
	for y := range c.Gfx {
		for x := range c.Gfx[y] {
			c.Gfx[y][x] = false
		}
	}
}

