package Chip8
import "fmt"

/**
* Fetch Opcode from momery specified at program counter (PC)
*
 */
func (c *Chip8) FetchOpcode() byte {
	opcode := c.Memory[c.Pc]<<8 | c.Memory[c.Pc+1]
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
	}
}

