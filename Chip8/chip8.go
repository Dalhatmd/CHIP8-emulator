package Chip8

const (
	programStartAddress = 0x200
	fontMemoryAddress   = 0x050
	charPixelWidth      = 4
	charPixelHeight     = 5
)

type Chip8 struct {
	Opcode      uint16       //current instructioon being executed
	Memory      [4096]byte   //array representating Memory
	V           [16]uint8    //flag for certain operations
	I           uint16       //register for storing sprites or data
	Pc          uint16       //points to the current instruction in memory
	Gfx         [32][64]bool //2d array repr monochrome display. each value either 0 or 1
	DelayTimer  uint8        //8-bit timer that decrements at 60hz
	SoundTimer  uint8        //8-bit timer that plays a sound when non-zero
	Stack       [16]uint16   //array used to store return addresses for subroutines
	Sp          uint16       //stack pointer; points to top of stack
	Key         [16]bool    //16 bit values representing the chip8 keypad
	PixelWidth  uint16       //width of a single pixel. Value used in drawing
	PixelHeight uint16       //height of a single pixel. Value used in drawing
	Scale       int          //Scaling factor for rendering display
}

// New initializes a new Chip8 instance with default values.
// The program counter starts at 0x200, and the font set is loaded into memory starting at 0x050.
// The pixel dimensions are set to the default character size of 4x5 pixels.
func New() *Chip8 {
	c := &Chip8{
		Pc:          programStartAddress, // Program counter starts at 0x200
		PixelWidth:  charPixelWidth,      // Default pixel PixelWidth
		PixelHeight: charPixelHeight,     // Default pixel PixelHeight
		Scale:       2,                   // Default scale factor for display
	}

	for i, v := range fontset {
		c.Memory[fontMemoryAddress+i] = v
	}

	return c
}

var fontset = []byte{
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
	0xF0, 0x80, 0xF0, 0x80, 0x80, // F
}

func (c *Chip8) LoadRom(rom []byte) {
	for i, b := range rom {
		if i+programStartAddress < len(c.Memory) {
			c.Memory[i+programStartAddress] = b
		}
	}
}

// update timers
func (c *Chip8) UpdateTimers() {
	if c.DelayTimer > 0 {
		c.DelayTimer--
	}
	if c.SoundTimer > 0 {
		c.SoundTimer--
		// Beep sound
	}
}
