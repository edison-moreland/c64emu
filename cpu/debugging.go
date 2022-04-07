package cpu

// Debugger code is pretty rough, you have been warned

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/chzyer/readline"
	"github.com/edison-moreland/c64emu/cpuinfo"
)

type debug struct {
	*readline.Instance
	stepcount      int
	stepping       bool
	breakpoints    map[uint16]bool
	breakpointFlag uint8
	breakpointOp   uint8
	quiet          bool
}

func newDebug() debug {
	// rl will get leaked, for now I don't care
	rl, err := readline.New("debug> ")
	if err != nil {
		panic(err)
	}
	rl.Config.AutoComplete = readline.NewPrefixCompleter(
		readline.PcItem("help"),
		readline.PcItem("run"),
		readline.PcItem("step"),
		readline.PcItem("break",
			readline.PcItem("clear"),
			readline.PcItem("list"),
			readline.PcItem("flag",
				readline.PcItem("n"),
				readline.PcItem("v"),
				readline.PcItem("b"),
				readline.PcItem("d"),
				readline.PcItem("i"),
				readline.PcItem("z"),
				readline.PcItem("c"),
			),
			readline.PcItem("op"),
			readline.PcItem("quiet"),
		),
	)
	rl.Config.HistoryFile = "/tmp/c64eemu-debug.history"
	rl.Config.DisableAutoSaveHistory = false

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
	instruction, err := cpuinfo.Decode(raw[0])
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
	if c.debug.quiet {
		return
	}

	inst := formatInstruction(c.Memory.Read(c.PC, 3))
	fmt.Println("------------+")
	fmt.Printf("instruction | %s\n", inst)
	fmt.Printf("registers   | PC: %04X S: %02X A: %02X X: %02X Y: %02X \n", c.PC, c.S, c.A, c.X, c.Y)
	fmt.Printf("status      | N: %v V: %v -: 0 B: %v D: %v I: %v Z: %v C: %v \n",
		c.flag(cpuinfo.Status_Negative),
		c.flag(cpuinfo.Status_Overflow),
		c.flag(cpuinfo.Status_BreakCommand),
		c.flag(cpuinfo.Status_Decimal),
		c.flag(cpuinfo.Status_InterruptDisable),
		c.flag(cpuinfo.Status_Zero),
		c.flag(cpuinfo.Status_Carry))

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
	// Opcode breakpoint
	opcode := c.Memory.ReadByte(c.PC)
	if c.debug.breakpointOp != 0xFF && opcode == c.debug.breakpointOp {
		c.debug.stepping = true
	}

	// Flag breakpoint
	if c.debug.breakpointFlag != 0 && c.isFlagSet(c.debug.breakpointFlag) {
		c.debug.stepping = true
	}

	// PC breakpoint
	if _, ok := c.debug.breakpoints[c.PC]; ok {
		c.debug.stepping = true
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
	case command == "help":
		fmt.Println("Available commands:")
		fmt.Println("  help                 - show this help message")
		fmt.Println("  quiet			   	- toggle quiet mode")
		fmt.Println("  run                  - run until a breakpoint is hit")
		fmt.Println("  step <count>         - step <count> instructions")
		fmt.Println("  break <addr>         - set a breakpoint at <addr>. <addr> is a 4-digit hex number")
		fmt.Println("  break list           - list all breakpoints")
		fmt.Println("  break clear          - clear breakpoint that just triggered")
		fmt.Println("  break flag [nvbdizc] - set breakpoint on flag sete")
		fmt.Println("  break op <opcode>    - set breakpoint on opcode. <opcode> is a 2-digit hex number")
		fmt.Println()
		fmt.Println("Hint: <enter> with no command will step one instruction, <ctrl-c> to stop emulation")
		goto input
	case command == "run":
		c.debug.stepping = false
	case command == "quiet":
		c.debug.quiet = !c.debug.quiet
		goto input
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
			if c.debug.breakpointFlag != 0 {
				c.debug.breakpointFlag = 0
			}
			if opcode == c.debug.breakpointOp {
				c.debug.breakpointOp = 0xFF
			}

		case "list":
			for addr := range c.debug.breakpoints {
				fmt.Printf("$%04X\n", addr)
			}

		case "flag":
			if len(args) == 2 {
				fmt.Println("Flag not specified")
				goto input
			}
			switch args[2] {
			case "n":
				c.debug.breakpointFlag = cpuinfo.Status_Negative
			case "v":
				c.debug.breakpointFlag = cpuinfo.Status_Overflow
			case "b":
				c.debug.breakpointFlag = cpuinfo.Status_BreakCommand
			case "d":
				c.debug.breakpointFlag = cpuinfo.Status_Decimal
			case "i":
				c.debug.breakpointFlag = cpuinfo.Status_InterruptDisable
			case "z":
				c.debug.breakpointFlag = cpuinfo.Status_Zero
			case "c":
				c.debug.breakpointFlag = cpuinfo.Status_Carry
			default:
				fmt.Println("Invalid flag")
				goto input
			}

		case "op":
			if len(args) == 2 {
				fmt.Println("Opcode not specified")
				goto input
			}
			opcode, err := hex.DecodeString(args[2])
			if err != nil {
				fmt.Println("Invalid opcode")
				goto input
			}

			c.debug.breakpointOp = opcode[0]
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
	case command == "":
		return
	default:
		fmt.Println("Unknown command")
		goto input
	}
}
