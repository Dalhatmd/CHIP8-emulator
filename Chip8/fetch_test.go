package Chip8
import (
	"testing"
)

func testFetchOpcode(t *testing.T) {
	chip := New()

	chip.Memory[0x200] = 0x12
	chip.Memory[0x201] = 0x34

	opcode := chip.FetchOpcode()
	if opcode != 0x1234 {
		t.Errorf("Expected opcode 0x1234, got 0x%04X", opcode)
	}
}

func TestExecuteOpcode(t *testing.T) {
	chip := &Chip8{
		Pc: 0x200,
	}

	// Test jump instrucion (0x1NNN)
	opcode := uint16(0x1234)
	chip.ExecuteOpcode(opcode)

	if chip.Pc != 0x0234 {
		t.Errorf("Expected opcode 0x234, got 0x%X", chip.Pc)
	}
}
