package cpu

import (
	"github.com/edison-moreland/c64emu/cpuinfo"
	"github.com/edison-moreland/c64emu/memory"
)

type CPU struct {
	cpuinfo.Registers
	Memory memory.Client

	InterruptPending bool
	InterruptVector  uint16

	shouldStop bool

	debugCPU   bool
	debugStack bool
	debug      debug
}

func New(memory memory.Client, debugCPU, debugStack bool) *CPU {
	return &CPU{
		Registers: cpuinfo.Registers{
			A: 0x0,
			X: 0x0,
			Y: 0x0,
			S: 0xFF,
			P: 0x00,
		},
		Memory: memory,

		InterruptPending: false,
		debugCPU:         debugCPU,
		debugStack:       debugStack,
		debug:            newDebug(),
	}
}

func (c *CPU) Start(doneChan chan<- interface{}) {
	// Make sure banks are in the right mode
	c.Memory.WriteByte(0x0001, 0xFF)
	c.Memory.WriteByte(0x0000, 0xFF)

	// Reset CPU, give control to kernal
	c.Interrupt(cpuinfo.Vector_RESET)

	go func() {
		for !c.shouldStop {
			c.handleInterrupt()

			if c.debugCPU {
				c.debugHook()
			}

			opcode := c.Memory.ReadByte(c.PC)
			instruction, err := cpuinfo.Decode(opcode)
			if err != nil {
				// Opcode not found
				panic(err)
			}

			target := c.executeAddressingMode(instruction)
			c.PC += c.executeInstruction(instruction, target)
		}

		doneChan <- nil
	}()
}

func (c *CPU) Stop() {
	c.shouldStop = true
}

func (c *CPU) Interrupt(vector uint16) {
	if (!c.isFlagSet(cpuinfo.Status_InterruptDisable)) || (vector == cpuinfo.Vector_NMI) {
		c.InterruptPending = true
		c.InterruptVector = vector
	}
}

func (c *CPU) handleInterrupt() {
	if !c.InterruptPending {
		return
	}

	if c.InterruptVector != cpuinfo.Vector_RESET {
		c.stackPushWord(c.PC)
		c.stackPushByte(c.P)
	}

	// c.setFlag(Status_InterruptDisable)
	c.PC = c.Memory.ReadWord(c.InterruptVector)
	c.InterruptPending = false
}
