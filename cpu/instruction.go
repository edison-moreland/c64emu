package cpu

import (
	"fmt"

	"github.com/edison-moreland/c64emu/cpuinfo"
)

func (c *CPU) setZeroNegative(value uint8) {
	if value == 0 {
		c.setFlag(cpuinfo.Status_Zero)
	} else {
		c.clearFlag(cpuinfo.Status_Zero)
	}

	if value&(1<<7) != 0 {
		c.setFlag(cpuinfo.Status_Negative)
	} else {
		c.clearFlag(cpuinfo.Status_Negative)
	}
}

func (c *CPU) setOverflow(value uint8, result uint8) {
	if (value & (1 << 7)) != (result & (1 << 7)) {
		c.setFlag(cpuinfo.Status_Overflow)
	} else {
		c.clearFlag(cpuinfo.Status_Overflow)
	}
}

func (c *CPU) executeInstruction(inst cpuinfo.Instruction, target uint16) (pcIncrement uint16) {
	pcIncrement = inst.Size

	switch inst.Mnemonic {
	// Arithmetic
	case cpuinfo.Mnemonic_AddwithCarry: // ADC
		// NOTE: This does not consider decimal mode
		original := c.A
		value := uint16(original) + uint16(c.Memory.ReadByte(target)) + uint16(c.flag(cpuinfo.Status_Carry))
		c.A = uint8(value)

		if value > 0xFF {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(c.A)
		c.setOverflow(original, c.A)

	case cpuinfo.Mnemonic_Compare: // CMP
		value := c.Memory.ReadByte(target)
		result := c.A - value

		if value >= c.A {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(result)

	case cpuinfo.Mnemonic_CompareXRegister: // CPX
		value := c.Memory.ReadByte(target)
		result := c.X - value

		if value >= c.X {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(result)

	case cpuinfo.Mnemonic_CompareYRegister: // CPY
		value := c.Memory.ReadByte(target)
		result := c.Y - value

		if value >= c.Y {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(result)

	case cpuinfo.Mnemonic_SubtractwithCarry: // SBC
		// NOTE: This does not consider decimal mode
		fetched := uint16(c.Memory.ReadByte(target)) ^ 0x00FF
		original := c.A
		value := uint16(original) + fetched + uint16(c.flag(cpuinfo.Status_Carry))
		c.A = uint8(value & 0x00FF)

		if value&0xFF00 != 0 {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(c.A)
		c.setOverflow(original, c.A)

	// Branch
	case cpuinfo.Mnemonic_BranchifCarryClear: // BCC
		if !c.isFlagSet(cpuinfo.Status_Carry) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifCarrySet: // BCS
		if c.isFlagSet(cpuinfo.Status_Carry) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifNotEqual: // BNE
		if !c.isFlagSet(cpuinfo.Status_Zero) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifEqual: // BEQ
		if c.isFlagSet(cpuinfo.Status_Zero) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifPlus: // BPL
		if !c.isFlagSet(cpuinfo.Status_Negative) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifMinus: // BMI
		if c.isFlagSet(cpuinfo.Status_Negative) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifOverflowClear: // BVC
		if !c.isFlagSet(cpuinfo.Status_Overflow) {
			c.PC = target
			pcIncrement = 0
		}

	case cpuinfo.Mnemonic_BranchifOverflowSet: // BVS
		if c.isFlagSet(cpuinfo.Status_Overflow) {
			c.PC = target
			pcIncrement = 0
		}

	// Control
	case cpuinfo.Mnemonic_ForceInterrupt: // BRK
		c.P |= cpuinfo.Status_BreakCommand
		c.Interrupt(cpuinfo.Vector_IRQ)

	case cpuinfo.Mnemonic_Jump: // JMP
		c.PC = target
		pcIncrement = 0

	case cpuinfo.Mnemonic_JumptoSubroutine: // JSR
		c.stackPushWord(c.PC + inst.Size)
		c.PC = target
		pcIncrement = 0

	case cpuinfo.Mnemonic_ReturnfromInterrupt: // RTI
		c.P = c.stackPopByte()
		c.clearFlag(cpuinfo.Status_BreakCommand)
		// c.clearFlag(cpuinfo.Status_InterruptDisable)

		c.PC = c.stackPopWord()
		pcIncrement = 0

	case cpuinfo.Mnemonic_ReturnfromSubroutine: // RTS
		c.PC = c.stackPopWord()
		pcIncrement = 0

	// Flags
	case cpuinfo.Mnemonic_ClearCarryFlag: // CLC
		c.clearFlag(cpuinfo.Status_Carry)

	case cpuinfo.Mnemonic_ClearDecimalMode: // CLD
		c.clearFlag(cpuinfo.Status_Decimal)

	case cpuinfo.Mnemonic_ClearInterruptDisable: // CLI
		c.clearFlag(cpuinfo.Status_InterruptDisable)

	case cpuinfo.Mnemonic_ClearOverflowFlag: // CLV
		c.clearFlag(cpuinfo.Status_Overflow)

	case cpuinfo.Mnemonic_SetCarryFlag: // SEC
		c.setFlag(cpuinfo.Status_Carry)

	case cpuinfo.Mnemonic_SetDecimalFlag: // SED
		c.setFlag(cpuinfo.Status_Decimal)

	case cpuinfo.Mnemonic_SetInterruptDisable: // SEI
		c.setFlag(cpuinfo.Status_InterruptDisable)

	// Increment/Decrement
	case cpuinfo.Mnemonic_DecrementMemory: // DEC
		value := c.Memory.ReadByte(target)
		value--
		c.Memory.WriteByte(target, value)
		c.setZeroNegative(value)

	case cpuinfo.Mnemonic_DecrementXRegister: // DEX
		c.X--
		c.setZeroNegative(c.X)

	case cpuinfo.Mnemonic_DecrementYRegister: // DEY
		c.Y--
		c.setZeroNegative(c.Y)

	case cpuinfo.Mnemonic_IncrementMemory: // INC
		value := c.Memory.ReadByte(target)
		value++
		c.Memory.WriteByte(target, value)
		c.setZeroNegative(value)

	case cpuinfo.Mnemonic_IncrementXRegister: // INX
		c.X++
		c.setZeroNegative(c.X)

	case cpuinfo.Mnemonic_IncrementYRegister: // INY
		c.Y++
		c.setZeroNegative(c.Y)

	// Load
	case cpuinfo.Mnemonic_LoadAccumulator: // LDA
		c.A = c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	case cpuinfo.Mnemonic_LoadXRegister: // LDX
		c.X = c.Memory.ReadByte(target)
		c.setZeroNegative(c.X)

	case cpuinfo.Mnemonic_LoadYRegister: // LDY
		c.Y = c.Memory.ReadByte(target)
		c.setZeroNegative(c.Y)

	case cpuinfo.Mnemonic_StoreAccumulator: // STA
		c.Memory.WriteByte(target, c.A)

	case cpuinfo.Mnemonic_StoreXRegister: // STX
		c.Memory.WriteByte(target, c.X)

	case cpuinfo.Mnemonic_StoreYRegister: // STY
		c.Memory.WriteByte(target, c.Y)

	// Logic
	case cpuinfo.Mnemonic_LogicalAND: // AND
		c.A &= c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	case cpuinfo.Mnemonic_BitTest: // BIT
		value := c.A & c.Memory.ReadByte(target)

		if value&(1<<6) != 0 {
			c.setFlag(cpuinfo.Status_Overflow)
		} else {
			c.clearFlag(cpuinfo.Status_Overflow)
		}

		c.setZeroNegative(value)

	case cpuinfo.Mnemonic_ExclusiveOR: // EOR
		c.A ^= c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	case cpuinfo.Mnemonic_LogicalOR: // ORA
		c.A |= c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	// No Operation
	case cpuinfo.Mnemonic_NoOperation: // NOP
		break // Do nothing

	// Shift
	case cpuinfo.Mnemonic_ArithmeticShiftLeft: // ASL
		var orig, value uint8
		if inst.AddressingMode == cpuinfo.AddressingMode_Accumulator {
			orig = c.A
			value = c.A << 1
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = orig << 1
			c.Memory.WriteByte(target, value)
		}

		if orig&(1<<7) != 0 {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(value)

	case cpuinfo.Mnemonic_LogicalShiftRight: // LSR
		var orig, value uint8
		if inst.AddressingMode == cpuinfo.AddressingMode_Accumulator {
			orig = c.A
			value = c.A >> 1
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = orig >> 1
			c.Memory.WriteByte(target, value)
		}

		if orig&(1) != 0 {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(value)

	case cpuinfo.Mnemonic_RotateLeft: // ROL
		incommingBit := uint8(0)
		if c.isFlagSet(cpuinfo.Status_Carry) {
			incommingBit = 1
		}

		var orig, value uint8
		if inst.AddressingMode == cpuinfo.AddressingMode_Accumulator {
			orig = c.A
			value = (c.A << 1) | incommingBit
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = (orig << 1) | incommingBit
			c.Memory.WriteByte(target, value)
		}

		if orig&(1<<7) != 0 {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(value)

	case cpuinfo.Mnemonic_RotateRight: // ROR
		incommingBit := uint8(0)
		if c.isFlagSet(cpuinfo.Status_Carry) {
			incommingBit = (1 << 7)
		}

		var orig, value uint8
		if inst.AddressingMode == cpuinfo.AddressingMode_Accumulator {
			orig = c.A
			value = (c.A >> 1) | incommingBit
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = (orig >> 1) | incommingBit
			c.Memory.WriteByte(target, value)
		}

		if orig&(1) != 0 {
			c.setFlag(cpuinfo.Status_Carry)
		} else {
			c.clearFlag(cpuinfo.Status_Carry)
		}

		c.setZeroNegative(value)

	// Stack
	case cpuinfo.Mnemonic_PushAccumulator: // PHA
		c.stackPushByte(c.A)

	case cpuinfo.Mnemonic_PushProcessorStatus: // PHP
		c.stackPushByte(c.P)

	case cpuinfo.Mnemonic_PullAccumulator: // PLA
		c.A = c.stackPopByte()
		c.setZeroNegative(c.A)

	case cpuinfo.Mnemonic_PullProcessorStatus: // PLP
		c.P = c.stackPopByte()

	// Transfer
	case cpuinfo.Mnemonic_TransferAccumulatortoX: // TAX
		c.X = c.A
		c.setZeroNegative(c.X)

	case cpuinfo.Mnemonic_TransferAccumulatortoY: // TAY
		c.Y = c.A
		c.setZeroNegative(c.Y)

	case cpuinfo.Mnemonic_TransferStackPointertoX: // TSX
		c.X = c.S
		c.setZeroNegative(c.X)

	case cpuinfo.Mnemonic_TransferXtoAccumulator: // TXA
		c.A = c.X
		c.setZeroNegative(c.A)

	case cpuinfo.Mnemonic_TransferXtoStackPointer: // TXS
		c.S = c.X
		c.setZeroNegative(c.S)

	case cpuinfo.Mnemonic_TransferYtoAccumulator: // TYA
		c.A = c.Y
		c.setZeroNegative(c.A)

	default:
		panic(fmt.Sprintf("Instruction not implemented %s", inst.Mnemonic.String()))
	}

	return pcIncrement
}
