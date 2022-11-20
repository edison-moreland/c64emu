package cpu

import (
	"github.com/edison-moreland/c64emu/cpuinfo"
	"github.com/edison-moreland/c64emu/memory"
	"github.com/edison-moreland/c64emu/trace"
)

type CPU struct {
	cpuinfo.Registers
	Memory memory.Client

	InterruptPending bool
	InterruptVector  uint16

	shouldStop bool

	trace       bool
	tracer      *trace.SystemTracer
	traceStack  bool
	tracerStack *trace.SystemTracer
}

func New(memory memory.Client, tracer *trace.SystemTracer, debugCPU, debugStack bool) *CPU {
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
		trace:            debugCPU,
		tracer:           tracer,
		traceStack:       debugStack,
		tracerStack:      tracer.System("stack"),
		//debug:            newDebug(),
	}
}

func (c *CPU) Start(clock <-chan bool) (chan<- Command, <-chan interface{}) {
	// Make sure banks are in the right mode
	c.Memory.WriteByte(0x0001, 0xFF)
	c.Memory.WriteByte(0x0000, 0xFF)

	// Reset CPU, give control to kernal
	c.Interrupt(cpuinfo.Vector_RESET)

	doneChan := make(chan interface{})
	commandChan := make(chan Command)

	go func() {
	topLoop:
		for {
			select {
			case <-clock:
				if c.shouldStop {
					close(doneChan)
					break topLoop
				}

				c.handleInterrupt()

				opcode := c.Memory.ReadByte(c.PC)
				instruction, err := cpuinfo.Decode(opcode)
				if err != nil {
					// Opcode not found
					panic(err)
				}

				target := c.executeAddressingMode(instruction)
				c.PC += c.executeInstruction(instruction, target)
			case icmd := <-commandChan:
				switch cmd := icmd.(type) {
				case *InterruptCommand:
					c.Interrupt(cmd.Vector)

				}

			}
		}
	}()

	return commandChan, doneChan
}

func (c *CPU) Stop() {
	c.shouldStop = true
}

func (c *CPU) Interrupt(vector uint16) {
	if c.trace {
		c.tracer.Trace("interrupt", trace.Uint16("vector", vector))
	}

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
