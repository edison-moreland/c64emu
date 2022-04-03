package cpu

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
)

type debug struct {
	*readline.Instance
	stepcount   int
	stepping    bool
	breakpoints map[uint16]bool
}

func newDebug() debug {
	// rl will get leaked, for now I don't care
	rl, err := readline.New("debug> ")
	if err != nil {
		panic(err)
	}
	rl.Config.AutoComplete = readline.NewPrefixCompleter(
		readline.PcItem("run"),
		readline.PcItem("step"),
		readline.PcItem("break",
			readline.PcItem("clear"),
			readline.PcItem("list")),
	)

	return debug{
		Instance:    rl,
		stepcount:   0,
		stepping:    true,
		breakpoints: make(map[uint16]bool),
	}
}

func (c *CPU) debugHook() {
	c.debugInformation()
	c.debugPrompt()
}

func formatInstruction(raw []byte) string {
	instruction, err := Decode(raw[0])
	if err != nil {
		return fmt.Sprintf("Invalid opcode: %02X", raw[0])
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

	return fmt.Sprintf("%s %s (%s)",
		instruction.Mnemonic.String(),
		addressingModeStr,
		instruction.Category.String(),
	)
}

func (c *CPU) debugInformation() {
	inst := formatInstruction(c.Memory.Read(c.PC, 3))
	fmt.Println("------------+")
	fmt.Printf("instruction | %s\n", inst)
	fmt.Printf("registers   | PC: %04X S: %02X A: %02X X: %02X Y: %02X \n", c.PC, c.S, c.A, c.X, c.Y)
	fmt.Printf("status      | N: %v V: %v -: 0 B: %v D: %v I: %v Z: %v C: %v \n",
		c.flag(Status_Negative),
		c.flag(Status_Overflow),
		c.flag(Status_BreakCommand),
		c.flag(Status_Decimal),
		c.flag(Status_InterruptDisable),
		c.flag(Status_Zero),
		c.flag(Status_Carry))

	next8Bytes := c.Memory.Read(c.PC, 8)
	fmt.Printf("memory      | %04X: %02X %02X %02X %02X %02X %02X %02X %02X\n",
		c.PC,
		next8Bytes[0],
		next8Bytes[1],
		next8Bytes[2],
		next8Bytes[3],
		next8Bytes[4],
		next8Bytes[5],
		next8Bytes[6],
		next8Bytes[7])
	fmt.Printf("interrupts  | Pending: %-5v Vector: %04X\n", c.InterruptPending, c.InterruptVector)
}

func (c *CPU) debugPrompt() {
	if _, ok := c.breakpoints[c.PC]; ok {
		c.stepping = true
	}

	if !c.debug.stepping {
		return
	}

	if c.debug.stepcount > 0 {
		c.debug.stepcount--
		return
	}
	fmt.Println("------------+")

input:
	command, err := c.debug.Readline()
	if err != nil {
		if errors.Is(err, readline.ErrInterrupt) {
			fmt.Println("Stopping execution")
			c.Stop()
			return
		}
		panic(err)
	}

	switch {
	case strings.HasPrefix(command, "run"):
		c.debug.stepping = false
	case strings.HasPrefix(command, "step"):
		args := strings.Split(command, " ")
		if len(args) == 1 {
			fmt.Println("Step count not specified")
			goto input
		}

		stepcount, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Invalid step count")
			goto input
		}

		c.debug.stepcount = stepcount
	case strings.HasPrefix(command, "break"):
		args := strings.Split(command, " ")
		if len(args) == 1 {
			fmt.Println("Breakpoint not specified")
			goto input
		}

		switch args[1] {
		case "clear":
			delete(c.debug.breakpoints, c.PC)

		case "list":
			for addr := range c.debug.breakpoints {
				fmt.Printf("$%04X\n", addr)
			}

		default:
			addressBytes, err := hex.DecodeString(args[1])
			if err != nil {
				fmt.Println("Invalid breakpoint address")
				goto input
			}
			address := binary.BigEndian.Uint16(addressBytes)
			c.debug.breakpoints[address] = true
		}

		goto input
	}
}
