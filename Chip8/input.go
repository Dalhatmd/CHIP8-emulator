package Chip8

import (
	"fmt"
	r1 "github.com/gen2brain/raylib-go/raylib"
)

var keyMap = map[int32]byte{
	r1.KeyOne:   0x01,
	r1.KeyTwo:   0x02,
	r1.KeyThree: 0x03,
	r1.KeyFour:  0x0C,
	r1.KeyQ:     0x04,
	r1.KeyW:     0x05,
	r1.KeyE:     0x06,
	r1.KeyR:     0x0D,
	r1.KeyA:     0x07,
	r1.KeyS:     0x08,
	r1.KeyD:     0x09,
	r1.KeyF:     0x0E,
	r1.KeyZ:     0x0A,
	r1.KeyX:     0x00,
	r1.KeyC:     0x0B,
	r1.KeyV:     0x0F,
}

func (c *Chip8) HandleInput() {
	for i := range c.Key {
		c.Key[i] = false
	}
	

	for rlKey, chipKey := range keyMap {
		if r1.IsKeyPressed(rlKey) {
			fmt.Println(chipKey, "pressed")
		}
		c.Key[chipKey] = r1.IsKeyDown(rlKey)
	}
}
