package cpu

import (
	"github.com/edison-moreland/c64emu/memory"
)

//go:generate go run ./generate/ -package=cpu -outfile=./cpuinfo.go

type CPU struct {
	Registers
	Memory memory.Client

	InterruptPending bool
	InterruptVector  uint16

	shouldStop bool

	debugMode bool
}

func New(memory memory.Client, debug bool) *CPU {
	return &CPU{
		Registers: Registers{
			A: 0x0,
			X: 0x0,
			Y: 0x0,
			S: 0xFF,
			P: 0x00,
		},
		Memory: memory,

		InterruptPending: false,
		debugMode:        debug,
	}
}

func (c *CPU) Start() {
	c.Interrupt(Vector_RESET)
	for !c.shouldStop {
		c.handleInterrupt()

		if c.debugMode {
			c.debug()
		}

		opcode := c.Memory.ReadByte(c.PC)
		instruction, err := Decode(opcode)
		if err != nil {
			// Opcode not found
			panic(err)
		}

		target := c.executeAddressingMode(instruction)
		c.PC += c.executeInstruction(instruction, target)
	}
}

func (c *CPU) Stop() {
	c.shouldStop = true
}

func (c *CPU) Interrupt(vector uint16) {
	if (!c.isFlagSet(Status_InterruptDisable)) || (vector == Vector_NMI) {
		c.InterruptPending = true
		c.InterruptVector = vector
	}
}

func (c *CPU) handleInterrupt() {
	if !c.InterruptPending {
		return
	}

	if c.InterruptVector != Vector_RESET {
		c.stackPushWord(c.PC)
		c.stackPushByte(c.P)
	}

	c.setFlag(Status_InterruptDisable)
	c.PC = c.Memory.ReadWord(c.InterruptVector)
	c.InterruptPending = false
}
