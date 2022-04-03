package cpu

import (
	"encoding/binary"
)

func (c *CPU) stackAddress() uint16 {
	// The stack live in the second page of memory
	return binary.LittleEndian.Uint16([]byte{c.S, 0x01})
}

func (c *CPU) stackPushByte(b byte) {
	c.Memory.WriteByte(c.stackAddress(), b)
	c.S--
}

func (c *CPU) stackPushWord(w uint16) {
	c.Memory.WriteWord(c.stackAddress(), w)
	c.S -= 2
}

func (c *CPU) stackPopByte() byte {
	c.S++
	return c.Memory.ReadByte(c.stackAddress())
}

func (c *CPU) stackPopWord() uint16 {
	c.S += 2
	return c.Memory.ReadWord(c.stackAddress())
}
