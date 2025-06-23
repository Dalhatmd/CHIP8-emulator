package Chip8

import (
	"fmt"

	"github.com/go-playground/locales/nn"
)

/**
* Fetch Opcode from momery specified at program counter (PC)
*
 */
func (c *Chip8) FetchOpcode() uint16 {
	opcode := uint16(c.Memory[c.Pc])<<8 | uint16(c.Memory[c.Pc+1])
	c.Pc += 2
	return opcode
}

func (c *Chip8) ExecuteOpcode(opcode uint16) {
	switch opcode & 0xF000 {
	case 0x0000:
		switch opcode & 0x00ff {
		case 0x00E0:
			fmt.Println("Clear Display called")
			c.ClearDisplay()
		}
	case 0x1000:
		address := opcode & 0x0FFF
		fmt.Printf("Jumping to address: 0x%X\n", address)
		c.Pc = address
	case 0x00EE:
		c.Pc = c.Stack[c.Sp]
		c.Sp--
	case 0x2000:
		c.Stack[c.Sp] = c.Pc	
		c.Sp++
		address := opcode & 0x0FFF
		c.Pc = address
	case 0x3000:
		x := (opcode & 0x0F00) >> 8
		value := opcode & 0x00ff
		if c.V[x] == value {
			c.Pc += 2
		}
	case 0x6000:
		x := (opcode & 0x0F00) >> 8
		nn = opcode & 0x00ff
		v[x] = nn
	case 0x7000:
		x := (opcode & 0x0F00) >> 8
		nn := opcode & 0x00FF
		c.V[x] += nn
	case 0xA000:
		i = opcode & 0x0FFF
		c.I = i
	case 0xD000:
		x := c.V[(opcode & 0x0F00) >> 8]
		y := c.V[(opcode & 0x00F0) >> 4]
		height := opcode & 0x000F

		c.V[0xF] = 0


