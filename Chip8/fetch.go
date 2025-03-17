package Chip8

import "fmt"
/**
* Fetch Opcode from momery specified at program counter (PC)
*
*/
func (c *Chip8) FetchOpcode() {
	opcode := c.Memory[c.Pc] << 8 | c.Memory[c.Pc + 1]
	fmt.Println(opcode)
}

