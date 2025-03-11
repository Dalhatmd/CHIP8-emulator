package Chip8

type Chip8 struct {
	opcode uint16
	memory [4096]uint8
	V [16]uint8
	I uint16
	pc uint16
	gfx [64 * 32]uint8
	delay_timer uint8
	sound_timer uint8
	stack [16]uint16
	sp uint16
	key [16]uint8
}

func GetChip8() *Chip8 {
	return &Chip8{}
}

