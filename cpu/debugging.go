package cpu

import "fmt"

func (c *CPU) debug() {
	nextInstruction, err := Decode(c.Memory.ReadByte(c.PC))
	if err != nil {
		fmt.Printf("Invalid opcode: %02X\n", c.Memory.ReadByte(c.PC))
	} else {
		fmt.Printf("%-5s | %s\n",
			nextInstruction.Category.String(),
			nextInstruction.Mnemonic.String())
	}

	fmt.Printf("PC: %04X S: %02X A: %02X X: %02X Y: %02X \n", c.PC, c.S, c.A, c.X, c.Y)
	fmt.Printf("N: %v V: %v -: 0 B: %v D: %v I: %v Z: %v C: %v \n",
		c.flag(Status_Negative),
		c.flag(Status_Overflow),
		c.flag(Status_BreakCommand),
		c.flag(Status_Decimal),
		c.flag(Status_InterruptDisable),
		c.flag(Status_Zero),
		c.flag(Status_Carry))
	fmt.Printf("%04X: %02X %02X %02X %02X %02X %02X %02X %02X\n",
		c.PC,
		c.Memory.ReadByte(c.PC),
		c.Memory.ReadByte(c.PC+1),
		c.Memory.ReadByte(c.PC+2),
		c.Memory.ReadByte(c.PC+3),
		c.Memory.ReadByte(c.PC+4),
		c.Memory.ReadByte(c.PC+5),
		c.Memory.ReadByte(c.PC+6),
		c.Memory.ReadByte(c.PC+7))
	fmt.Printf("Interrupt Pending: %-5v Vector: %04X\n", c.InterruptPending, c.InterruptVector)
	fmt.Printf("> ")
	fmt.Scanln()
}
