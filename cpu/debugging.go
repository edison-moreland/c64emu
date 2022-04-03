package cpu

import (
	"encoding/binary"
	"fmt"
	"strings"
)

func printInstruction(raw []byte) {
	instruction, err := Decode(raw[0])
	if err != nil {
		fmt.Printf("Invalid opcode: %02X\n", raw[0])
		return
	}

	addressingModeStr := instruction.AddressingMode.String()
	if instruction.Size == 2 {
		var bytestr = fmt.Sprintf("$%02X", raw[1])
		addressingModeStr = strings.ReplaceAll(addressingModeStr, "d8", bytestr)
		addressingModeStr = strings.ReplaceAll(addressingModeStr, "a8", bytestr)
		addressingModeStr = strings.ReplaceAll(addressingModeStr, "r8", bytestr)
	}
	if instruction.Size == 3 {
		var word = binary.LittleEndian.Uint16(raw[1:])
		addressingModeStr = strings.ReplaceAll(addressingModeStr, "a16", fmt.Sprintf("$%04X", word))
	}

	fmt.Printf("%-5s | %s %s\n",
		instruction.Category.String(),
		instruction.Mnemonic.String(),
		addressingModeStr,
	)

}

func (c *CPU) debug() {
	printInstruction(c.Memory.Read(c.PC, 3))
	fmt.Printf("PC: %04X S: %02X A: %02X X: %02X Y: %02X \n", c.PC, c.S, c.A, c.X, c.Y)
	fmt.Printf("N: %v V: %v -: 0 B: %v D: %v I: %v Z: %v C: %v \n",
		c.flag(Status_Negative),
		c.flag(Status_Overflow),
		c.flag(Status_BreakCommand),
		c.flag(Status_Decimal),
		c.flag(Status_InterruptDisable),
		c.flag(Status_Zero),
		c.flag(Status_Carry))

	next8Bytes := c.Memory.Read(c.PC, 8)
	fmt.Printf("%04X: %02X %02X %02X %02X %02X %02X %02X %02X\n",
		c.PC,
		next8Bytes[0],
		next8Bytes[1],
		next8Bytes[2],
		next8Bytes[3],
		next8Bytes[4],
		next8Bytes[5],
		next8Bytes[6],
		next8Bytes[7])
	fmt.Printf("Interrupt Pending: %-5v Vector: %04X\n", c.InterruptPending, c.InterruptVector)
	fmt.Printf("> ")
	fmt.Scanln()
}
