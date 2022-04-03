package cpu

import (
	"encoding/binary"
	"fmt"
)

func (c *CPU) stackAddress() uint16 {
	// The stack live in the second page of memory
	addreess := binary.LittleEndian.Uint16([]byte{c.S, 0x01})

	if c.debugMode {
		fmt.Printf("stack  --- address $%04X\n", addreess)
	}

	return addreess
}

func (c *CPU) stackPushByte(b byte) {
	if c.debugMode {
		fmt.Printf("stack  --- pushByte $%02X\n", b)
	}

	c.Memory.WriteByte(c.stackAddress(), b)
	c.S--
}

func (c *CPU) stackPopByte() byte {
	c.S++
	value := c.Memory.ReadByte(c.stackAddress())

	if c.debugMode {
		fmt.Printf("stack  --- popByte  $%02X\n", value)
	}

	return value
}

func (c *CPU) stackPushWord(w uint16) {
	if c.debugMode {
		fmt.Printf("stack  --- pushWord $%04X\n", w)
	}

	wordBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(wordBytes, w)

	c.stackPushByte(wordBytes[0])
	c.stackPushByte(wordBytes[1])
}

func (c *CPU) stackPopWord() uint16 {
	wordBytes := make([]byte, 2)
	wordBytes[1] = c.stackPopByte()
	wordBytes[0] = c.stackPopByte()

	value := binary.LittleEndian.Uint16(wordBytes)

	if c.debugMode {
		fmt.Printf("stack  --- popWord  $%04X\n", value)
	}

	return value
}
