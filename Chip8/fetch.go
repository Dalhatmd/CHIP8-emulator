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
//		fmt.Printf("Jumping to address: 0x%X\n", address)
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
		c.V[x] = uint8(nn)

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
		// reset collision flag
		c.V[0xF] = 0
		// Get sprite coordinates
		xPos := c.V[x]
		yPos := c.V[y]

		// Draw sprite line by line
		for row := uint16(0); row < n; row++ {
			spriteByte := c.Memory[c.I+row]
			for col := uint16(0); col < 8; col++ {
				if (spriteByte & (0x80 >> col)) != 0 {
					screenX := (xPos + uint8(col)) % 64
					screenY := (yPos + uint8(row)) % 32

					// Check for collision and draw pixel using XOR
					if c.Gfx[screenY][screenX] { // If pixel was already ON
						c.V[0xF] = 1 // Set collision flag
					}
					c.Gfx[screenY][screenX] = !c.Gfx[screenY][screenX] // XOR operation
				}
			}
		}

	case 0xE000: // EX9E, EXA1
		x := (opcode & 0x0F00) >> 8
		switch opcode & 0x00FF {
		case 0x009E: // EX9E - SKP Vx
			if c.Key[c.V[x]] {
				c.Pc += 2
			}
		case 0x00A1: // EXA1 - SKNP Vx
			if !c.Key[c.V[x]] {
				c.Pc += 2
			}
		}
	case 0xF000:
		x := (opcode & 0x0F00) >> 8
		switch opcode & 0x00FF {
		case 0x0015: // FX15 - LD DT, Vx
			c.DelayTimer = c.V[x]
		case 0x0018: // FX18 - LD ST, Vx
			c.SoundTimer = c.V[x]
		case 0x001E: // FX1E - ADD I, Vx
			c.I += uint16(c.V[x])
		case 0x0029: // FX29 - LD F, Vx
			c.I = uint16(0x50 + (c.V[x] * 5)) // Each font character is 5 bytes long, starting at memory location 0x50
		case 0x0033: // FX33 - LD B, Vx
			value := c.V[x]
			c.Memory[c.I] = value / 100
			c.Memory[c.I+1] = (value / 10) % 10
			c.Memory[c.I+2] = value % 10
		case 0x0055: // FX55 - LD [I], Vx
			for i := uint16(0); i <= x; i++ {
				c.Memory[c.I+i] = c.V[i]
			}
		case 0x0065: // FX65 - LD Vx, [I]
			for i := uint16(0); i <= x; i++ {
				c.V[i] = c.Memory[c.I+i]
			}
		}
	default:
		fmt.Printf("Unknown opcode: 0x%X\n", opcode)
	}
}
