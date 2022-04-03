package cpu

import "fmt"

func (c *CPU) setZeroNegative(value uint8) {
	if value == 0 {
		c.setFlag(Status_Zero)
	} else {
		c.clearFlag(Status_Zero)
	}

	if value&(1<<7) != 0 {
		c.setFlag(Status_Negative)
	} else {
		c.clearFlag(Status_Negative)
	}
}

func (c *CPU) setOverflow(value uint8, result uint8) {
	if (value & (1 << 7)) != (result & (1 << 7)) {
		c.setFlag(Status_Overflow)
	} else {
		c.clearFlag(Status_Overflow)
	}
}

func (c *CPU) executeInstruction(inst Instruction, target uint16) (pcIncrement uint16) {
	pcIncrement = inst.Size

	switch inst.Mnemonic {
	// Arithmetic
	case Mnemonic_AddwithCarry: // ADC
		// NOTE: This does not consider decimal mode
		original := c.A
		value := uint16(original) + uint16(c.Memory.ReadByte(target)) + uint16(c.flag(Status_Carry))
		c.A = uint8(value)

		if value > 0xFF {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(c.A)
		c.setOverflow(original, c.A)

	case Mnemonic_Compare: // CMP
		value := c.Memory.ReadByte(target)
		result := c.A - value

		if value >= c.A {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(result)

	case Mnemonic_CompareXRegister: // CPX
		value := c.Memory.ReadByte(target)
		result := c.X - value

		if value >= c.X {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(result)

	case Mnemonic_CompareYRegister: // CPY
		value := c.Memory.ReadByte(target)
		result := c.Y - value

		if value >= c.Y {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(result)

	case Mnemonic_SubtractwithCarry: // SBC
		// NOTE: This does not consider decimal mode
		fetched := uint16(c.Memory.ReadByte(target)) ^ 0x00FF
		original := c.A
		value := uint16(original) + fetched + uint16(c.flag(Status_Carry))
		c.A = uint8(value & 0x00FF)

		if value&0xFF00 != 0 {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(c.A)
		c.setOverflow(original, c.A)

	// Branch
	case Mnemonic_BranchifCarryClear: // BCC
		if !c.isFlagSet(Status_Carry) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifCarrySet: // BCS
		if c.isFlagSet(Status_Carry) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifNotEqual: // BNE
		if !c.isFlagSet(Status_Zero) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifEqual: // BEQ
		if c.isFlagSet(Status_Zero) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifPlus: // BPL
		if !c.isFlagSet(Status_Negative) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifMinus: // BMI
		if c.isFlagSet(Status_Negative) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifOverflowClear: // BVC
		if !c.isFlagSet(Status_Overflow) {
			c.PC = target
			pcIncrement = 0
		}

	case Mnemonic_BranchifOverflowSet: // BVS
		if c.isFlagSet(Status_Overflow) {
			c.PC = target
			pcIncrement = 0
		}

	// Control
	case Mnemonic_ForceInterrupt: // BRK
		c.P |= Status_BreakCommand
		c.Interrupt(Vector_IRQ)

	case Mnemonic_Jump: // JMP
		c.PC = target
		pcIncrement = 0

	case Mnemonic_JumptoSubroutine: // JSR
		c.stackPushWord(c.PC + inst.Size)
		c.PC = target
		pcIncrement = 0

	case Mnemonic_ReturnfromInterrupt: // RTI
		c.P = c.stackPopByte()
		c.P &^= Status_BreakCommand

		c.PC = c.stackPopWord()
		pcIncrement = 0

	case Mnemonic_ReturnfromSubroutine: // RTS
		c.PC = c.stackPopWord()
		pcIncrement = 0

	// Flags
	case Mnemonic_ClearCarryFlag: // CLC
		c.clearFlag(Status_Carry)

	case Mnemonic_ClearDecimalMode: // CLD
		c.clearFlag(Status_Decimal)

	case Mnemonic_ClearInterruptDisable: // CLI
		c.clearFlag(Status_InterruptDisable)

	case Mnemonic_ClearOverflowFlag: // CLV
		c.clearFlag(Status_Overflow)

	case Mnemonic_SetCarryFlag: // SEC
		c.setFlag(Status_Carry)

	case Mnemonic_SetDecimalFlag: // SED
		c.setFlag(Status_Decimal)

	case Mnemonic_SetInterruptDisable: // SEI
		c.setFlag(Status_InterruptDisable)

	// Increment/Decrement
	case Mnemonic_DecrementMemory: // DEC
		value := c.Memory.ReadByte(target)
		value--
		c.Memory.WriteByte(target, value)
		c.setZeroNegative(value)

	case Mnemonic_DecrementXRegister: // DEX
		c.X--
		c.setZeroNegative(c.X)

	case Mnemonic_DecrementYRegister: // DEY
		c.Y--
		c.setZeroNegative(c.Y)

	case Mnemonic_IncrementMemory: // INC
		value := c.Memory.ReadByte(target)
		value++
		c.Memory.WriteByte(target, value)
		c.setZeroNegative(value)

	case Mnemonic_IncrementXRegister: // INX
		c.X++
		c.setZeroNegative(c.X)

	case Mnemonic_IncrementYRegister: // INY
		c.Y++
		c.setZeroNegative(c.Y)

	// Load
	case Mnemonic_LoadAccumulator: // LDA
		c.A = c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	case Mnemonic_LoadXRegister: // LDX
		c.X = c.Memory.ReadByte(target)
		c.setZeroNegative(c.X)

	case Mnemonic_LoadYRegister: // LDY
		c.Y = c.Memory.ReadByte(target)
		c.setZeroNegative(c.Y)

	case Mnemonic_StoreAccumulator: // STA
		c.Memory.WriteByte(target, c.A)

	case Mnemonic_StoreXRegister: // STX
		c.Memory.WriteByte(target, c.X)

	case Mnemonic_StoreYRegister: // STY
		c.Memory.WriteByte(target, c.Y)

	// Logic
	case Mnemonic_LogicalAND: // AND
		c.A &= c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	case Mnemonic_BitTest: // BIT
		value := c.A & c.Memory.ReadByte(target)

		if value&(1<<6) != 0 {
			c.setFlag(Status_Overflow)
		} else {
			c.clearFlag(Status_Overflow)
		}

		c.setZeroNegative(value)

	case Mnemonic_ExclusiveOR: // EOR
		c.A ^= c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	case Mnemonic_LogicalOR: // ORA
		c.A |= c.Memory.ReadByte(target)
		c.setZeroNegative(c.A)

	// No Operation
	case Mnemonic_NoOperation: // NOP
		break // Do nothing

	// Shift
	case Mnemonic_ArithmeticShiftLeft: // ASL
		var orig, value uint8
		if inst.AddressingMode == AddressingMode_Accumulator {
			orig = c.A
			value = c.A << 1
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = orig << 1
			c.Memory.WriteByte(target, value)
		}

		if orig&(1<<7) != 0 {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(value)

	case Mnemonic_LogicalShiftRight: // LSR
		var orig, value uint8
		if inst.AddressingMode == AddressingMode_Accumulator {
			orig = c.A
			value = c.A >> 1
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = orig >> 1
			c.Memory.WriteByte(target, value)
		}

		if orig&(1) != 0 {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(value)

	case Mnemonic_RotateLeft: // ROL
		incommingBit := uint8(0)
		if c.isFlagSet(Status_Carry) {
			incommingBit = 1
		}

		var orig, value uint8
		if inst.AddressingMode == AddressingMode_Accumulator {
			orig = c.A
			value = (c.A << 1) | incommingBit
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = (orig << 1) | incommingBit
			c.Memory.WriteByte(target, value)
		}

		if orig&(1<<7) != 0 {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(value)

	case Mnemonic_RotateRight: // ROR
		incommingBit := uint8(0)
		if c.isFlagSet(Status_Carry) {
			incommingBit = (1 << 7)
		}

		var orig, value uint8
		if inst.AddressingMode == AddressingMode_Accumulator {
			orig = c.A
			value = (c.A >> 1) | incommingBit
			c.A = value
		} else {
			orig := c.Memory.ReadByte(target)
			value = (orig >> 1) | incommingBit
			c.Memory.WriteByte(target, value)
		}

		if orig&(1) != 0 {
			c.setFlag(Status_Carry)
		} else {
			c.clearFlag(Status_Carry)
		}

		c.setZeroNegative(value)

	// Stack
	case Mnemonic_PushAccumulator: // PHA
		c.stackPushByte(c.A)

	case Mnemonic_PushProcessorStatus: // PHP
		c.stackPushByte(c.P)

	case Mnemonic_PullAccumulator: // PLA
		c.A = c.stackPopByte()
		c.setZeroNegative(c.A)

	case Mnemonic_PullProcessorStatus: // PLP
		c.P = c.stackPopByte()

	// Transfer
	case Mnemonic_TransferAccumulatortoX: // TAX
		c.X = c.A
		c.setZeroNegative(c.X)

	case Mnemonic_TransferAccumulatortoY: // TAY
		c.Y = c.A
		c.setZeroNegative(c.Y)

	case Mnemonic_TransferStackPointertoX: // TSX
		c.X = c.S
		c.setZeroNegative(c.X)

	case Mnemonic_TransferXtoAccumulator: // TXA
		c.A = c.X
		c.setZeroNegative(c.A)

	case Mnemonic_TransferXtoStackPointer: // TXS
		c.S = c.X
		c.setZeroNegative(c.S)

	case Mnemonic_TransferYtoAccumulator: // TYA
		c.A = c.Y
		c.setZeroNegative(c.A)

	default:
		panic(fmt.Sprintf("Instruction not implemented %s", inst.Mnemonic.String()))
	}

	return pcIncrement
}
