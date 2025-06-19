package Chip8
import "fmt"

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
	}
}

