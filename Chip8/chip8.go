package Chip8

type Chip8 struct {
	Opcode uint16
	Memory [4096]uint8
	V [16]uint8
	I uint16
	Pc uint16
	Gfx [32][64]uint8
	Delay_timer uint8
	Sound_timer uint8
	Stack [16]uint16
	Sp uint16
	Key [16]uint8
	PixelSize uint16
}

func GetChip8() *Chip8 {
	return &Chip8{PixelSize: 10}
}

