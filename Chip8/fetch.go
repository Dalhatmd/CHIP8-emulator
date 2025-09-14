package Chip8

import (
	"fmt"
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
	case 0x0000: //CLS
		switch opcode & 0x00ff {
		case 0x00E0:
			fmt.Println("Clear Display called")
			c.ClearDisplay()
		case 0x00EE:
			c.Sp--
			c.Pc = c.Stack[c.Sp]

		}
	case 0x1000: // JP addr
		address := opcode & 0x0FFF
		fmt.Printf("Jumping to address: 0x%X\n", address)
		c.Pc = address

	case 0x2000: // CALL addr
		c.Stack[c.Sp] = c.Pc	
		c.Sp++
		address := opcode & 0x0FFF
		c.Pc = address

	case 0x3000: // SE Vx, byte
		x := (opcode & 0x0F00) >> 8
		value := byte(opcode & 0x00ff)
		if c.V[x] == value {
			c.Pc += 2
		}

	case 0x6000: // LD Vx, byte
		x := (opcode & 0x0F00) >> 8
		nn := opcode & 0x00ff
		c.V[x] += uint8(nn)

	case 0x7000: // ADD Vx, byte
		x := (opcode & 0x0F00) >> 8
		nn := opcode & 0x00FF
		c.V[x] += uint8(nn)

	case 0xA000: // LD I, addr
		c.I = opcode & 0x0FFF

	case 0xD000:
		x := (opcode & 0x0F00) >> 8
		y := (opcode & 0x00F0) >> 4
		n := opcode & 0x000F
		fmt.Printf("Drawing sprite v[%x]=%d, V[%X]=[%d], height=%d\n", x, c.V[x], y, c.V[y], n)
	}
}
