package cpu

import (
	"encoding/binary"

	"github.com/edison-moreland/c64emu/trace"
)

func (c *CPU) stackAddress() uint16 {
	// The stack live in the second page of memory
	addreess := binary.LittleEndian.Uint16([]byte{c.S, 0x01})

	if c.traceStack {
		c.tracerStack.Trace("address", trace.Uint16("address", addreess))
	}

	return addreess
}

func (c *CPU) stackPushByte(b byte) {
	if c.traceStack {
		c.tracerStack.Trace("pushByte", trace.Byte("byte", b))
	}

	c.Memory.WriteByte(c.stackAddress(), b)
	c.S--
}

func (c *CPU) stackPopByte() byte {
	c.S++
	value := c.Memory.ReadByte(c.stackAddress())

	if c.traceStack {
		c.tracerStack.Trace("popByte", trace.Byte("byte", value))
	}

	return value
}

func (c *CPU) stackPushWord(w uint16) {
	if c.traceStack {
		c.tracerStack.Trace("pushWord", trace.Uint16("word", w))
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

	if c.traceStack {
		c.tracerStack.Trace("popWord", trace.Uint16("word", value))
	}

	return value
}
