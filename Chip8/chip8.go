package Chip8

type Chip8 struct {
	Opcode uint16
	Memory [4096]byte
	V [16]uint8
	I uint16
	Pc uint16
	Gfx [32][64]bool
	Delay_timer uint8
	Sound_timer uint8
	Stack [16]uint16
	Sp uint16
	Key [16]uint8
	PixelWidth uint16
	PixelHeight uint16
	Scale int
}

func GetChip8() *Chip8 {
	return &Chip8{}
}

/**
* Initialises Chip 8
* 1: Copies fontdata to memory
* 2: set program counter to start of program space (ox200)
*/
func (c *Chip8) Initialise () {

	fontset :=[]byte {
	0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
	0x20, 0x60, 0x20, 0x20, 0x70, // 1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
	0x90, 0x90, 0xF0, 0x10, 0x10, // 4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
	0xF0, 0x10, 0x20, 0x40, 0x40, // 7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
	0xF0, 0x90, 0xF0, 0x90, 0x90, // A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
	0xF0, 0x80, 0x80, 0x80, 0xF0, // C
	0xE0, 0x90, 0x90, 0x90, 0xE0, // D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
	0xF0, 0x80, 0xF0, 0x80, 0x80,  // F
}

	for i, v := range fontset {
		c.Memory[0x050+i] = v
	}
	c.Pc = 0x200
	c.PixelWidth = 4
	c.PixelHeight = 5
	c.Scale = 4
}
