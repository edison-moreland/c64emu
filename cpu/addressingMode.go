package cpu

import (
	"encoding/binary"

	"github.com/edison-moreland/c64emu/cpuinfo"
)

func zeroPageAddress(zpAddress byte) uint16 {
	return binary.LittleEndian.Uint16([]byte{zpAddress, 0x00})
}

func (c *CPU) executeAddressingMode(inst cpuinfo.Instruction) (target uint16) {
	switch inst.AddressingMode {
	case cpuinfo.AddressingMode_Accumulator:
	case cpuinfo.AddressingMode_Implied:
		target = 0x0000 // Implied doesn't need inst.AddressingMode target

	case cpuinfo.AddressingMode_Absolute:
		target = c.Memory.ReadWord(c.PC + 1)

	case cpuinfo.AddressingMode_AbsoluteIndirect:
		indirect := c.Memory.ReadWord(c.PC + 1)
		target = c.Memory.ReadWord(indirect)

	case cpuinfo.AddressingMode_Immediate:
		target = c.PC + 1

	case cpuinfo.AddressingMode_Relative:
		offset := int8(c.Memory.ReadByte(c.PC + 1))
		target = uint16(int32(c.PC+inst.Size) + int32(offset))

	case cpuinfo.AddressingMode_XIndexedAbsolute:
		target = c.Memory.ReadWord(c.PC+1) + uint16(c.X)

	case cpuinfo.AddressingMode_XIndexedZeroPage:
		target = zeroPageAddress(c.Memory.ReadByte(c.PC+1) + c.X)

	case cpuinfo.AddressingMode_XIndexedZeroPageIndirect:
		indirect := zeroPageAddress(c.Memory.ReadByte(c.PC+1) + c.X)
		target = c.Memory.ReadWord(indirect)

	case cpuinfo.AddressingMode_YIndexedAbsolute:
		target = c.Memory.ReadWord(c.PC+1) + uint16(c.Y)

	case cpuinfo.AddressingMode_YIndexedZeroPage:
		target = zeroPageAddress(c.Memory.ReadByte(c.PC+1) + c.Y)

	case cpuinfo.AddressingMode_ZeroPage:
		target = zeroPageAddress(c.Memory.ReadByte(c.PC + 1))

	case cpuinfo.AddressingMode_ZeroPageIndirectYIndexed:
		indirect := zeroPageAddress(c.Memory.ReadByte(c.PC + 1))
		target = c.Memory.ReadWord(indirect) + uint16(c.Y)
	}

	return target
}
