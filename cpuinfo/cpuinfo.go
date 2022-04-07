// Code generated on 2022-04-06T19:47:31-07:00; DO NOT EDIT.
// MOS 6502 (1975)

package cpuinfo

import "errors"

type Registers struct {
	A  uint8  // Accumulator
	X  uint8  // X Index Register
	Y  uint8  // Y Index Register
	S  uint8  // Stack Pointer
	P  uint8  // Processor Status
	PC uint16 // Program Counter
}
type Register int

const (
	Register_A  Register = iota
	Register_X  Register = iota
	Register_Y  Register = iota
	Register_S  Register = iota
	Register_P  Register = iota
	Register_PC Register = iota
)

var ErrRegisterNotFound = errors.New("Register not found")

func RegisterFromString(s string) (Register, error) {
	switch s {
	case "A":
		return Register_A, nil
	case "X":
		return Register_X, nil
	case "Y":
		return Register_Y, nil
	case "S":
		return Register_S, nil
	case "P":
		return Register_P, nil
	case "PC":
		return Register_PC, nil
	}
	return 0, ErrRegisterNotFound
}

const (
	Vector_NMI   uint16 = uint16(0xfffa)
	Vector_RESET uint16 = uint16(0xfffc)
	Vector_IRQ   uint16 = uint16(0xfffe)
)
const (
	Status_Negative         uint8 = 1 << 7 // N
	Status_Overflow         uint8 = 1 << 6 // V
	Status_BreakCommand     uint8 = 1 << 4 // B
	Status_Decimal          uint8 = 1 << 3 // D
	Status_InterruptDisable uint8 = 1 << 2 // I
	Status_Zero             uint8 = 1 << 1 // Z
	Status_Carry            uint8 = 1 << 0 // C
)

var ErrStatusNotFound = errors.New("Status not found")

func Status(s string) (uint8, error) {
	switch s {
	case "N":
		return Status_Negative, nil
	case "V":
		return Status_Overflow, nil
	case "B":
		return Status_BreakCommand, nil
	case "D":
		return Status_Decimal, nil
	case "I":
		return Status_InterruptDisable, nil
	case "Z":
		return Status_Zero, nil
	case "C":
		return Status_Carry, nil
	}
	return 0, ErrStatusNotFound
}

type Mnemonic int

const (
	Mnemonic_BranchifCarryClear      Mnemonic = iota // BCC
	Mnemonic_DecrementYRegister      Mnemonic = iota // DEY
	Mnemonic_NoOperation             Mnemonic = iota // NOP
	Mnemonic_PullProcessorStatus     Mnemonic = iota // PLP
	Mnemonic_SetInterruptDisable     Mnemonic = iota // SEI
	Mnemonic_BranchifPlus            Mnemonic = iota // BPL
	Mnemonic_ForceInterrupt          Mnemonic = iota // BRK
	Mnemonic_ClearInterruptDisable   Mnemonic = iota // CLI
	Mnemonic_LogicalShiftRight       Mnemonic = iota // LSR
	Mnemonic_TransferXtoStackPointer Mnemonic = iota // TXS
	Mnemonic_AddwithCarry            Mnemonic = iota // ADC
	Mnemonic_LogicalOR               Mnemonic = iota // ORA
	Mnemonic_PullAccumulator         Mnemonic = iota // PLA
	Mnemonic_ReturnfromSubroutine    Mnemonic = iota // RTS
	Mnemonic_BranchifEqual           Mnemonic = iota // BEQ
	Mnemonic_ClearCarryFlag          Mnemonic = iota // CLC
	Mnemonic_LoadXRegister           Mnemonic = iota // LDX
	Mnemonic_PushProcessorStatus     Mnemonic = iota // PHP
	Mnemonic_SubtractwithCarry       Mnemonic = iota // SBC
	Mnemonic_BranchifNotEqual        Mnemonic = iota // BNE
	Mnemonic_ClearDecimalMode        Mnemonic = iota // CLD
	Mnemonic_JumptoSubroutine        Mnemonic = iota // JSR
	Mnemonic_BranchifMinus           Mnemonic = iota // BMI
	Mnemonic_DecrementXRegister      Mnemonic = iota // DEX
	Mnemonic_IncrementXRegister      Mnemonic = iota // INX
	Mnemonic_StoreXRegister          Mnemonic = iota // STX
	Mnemonic_BranchifOverflowSet     Mnemonic = iota // BVS
	Mnemonic_StoreYRegister          Mnemonic = iota // STY
	Mnemonic_TransferYtoAccumulator  Mnemonic = iota // TYA
	Mnemonic_CompareXRegister        Mnemonic = iota // CPX
	Mnemonic_IncrementYRegister      Mnemonic = iota // INY
	Mnemonic_Jump                    Mnemonic = iota // JMP
	Mnemonic_TransferAccumulatortoX  Mnemonic = iota // TAX
	Mnemonic_TransferStackPointertoX Mnemonic = iota // TSX
	Mnemonic_CompareYRegister        Mnemonic = iota // CPY
	Mnemonic_IncrementMemory         Mnemonic = iota // INC
	Mnemonic_LoadAccumulator         Mnemonic = iota // LDA
	Mnemonic_RotateRight             Mnemonic = iota // ROR
	Mnemonic_SetDecimalFlag          Mnemonic = iota // SED
	Mnemonic_TransferXtoAccumulator  Mnemonic = iota // TXA
	Mnemonic_ArithmeticShiftLeft     Mnemonic = iota // ASL
	Mnemonic_Compare                 Mnemonic = iota // CMP
	Mnemonic_DecrementMemory         Mnemonic = iota // DEC
	Mnemonic_PushAccumulator         Mnemonic = iota // PHA
	Mnemonic_RotateLeft              Mnemonic = iota // ROL
	Mnemonic_SetCarryFlag            Mnemonic = iota // SEC
	Mnemonic_LogicalAND              Mnemonic = iota // AND
	Mnemonic_BitTest                 Mnemonic = iota // BIT
	Mnemonic_ClearOverflowFlag       Mnemonic = iota // CLV
	Mnemonic_ExclusiveOR             Mnemonic = iota // EOR
	Mnemonic_ReturnfromInterrupt     Mnemonic = iota // RTI
	Mnemonic_StoreAccumulator        Mnemonic = iota // STA
	Mnemonic_TransferAccumulatortoY  Mnemonic = iota // TAY
	Mnemonic_BranchifCarrySet        Mnemonic = iota // BCS
	Mnemonic_BranchifOverflowClear   Mnemonic = iota // BVC
	Mnemonic_LoadYRegister           Mnemonic = iota // LDY
)

func (m Mnemonic) String() string {
	switch m {
	case Mnemonic_PullProcessorStatus:
		return "PLP"
	case Mnemonic_PullAccumulator:
		return "PLA"
	case Mnemonic_BranchifNotEqual:
		return "BNE"
	case Mnemonic_ClearDecimalMode:
		return "CLD"
	case Mnemonic_BranchifOverflowClear:
		return "BVC"
	case Mnemonic_DecrementYRegister:
		return "DEY"
	case Mnemonic_ReturnfromSubroutine:
		return "RTS"
	case Mnemonic_RotateRight:
		return "ROR"
	case Mnemonic_LogicalAND:
		return "AND"
	case Mnemonic_ForceInterrupt:
		return "BRK"
	case Mnemonic_RotateLeft:
		return "ROL"
	case Mnemonic_BranchifCarrySet:
		return "BCS"
	case Mnemonic_IncrementXRegister:
		return "INX"
	case Mnemonic_TransferYtoAccumulator:
		return "TYA"
	case Mnemonic_CompareYRegister:
		return "CPY"
	case Mnemonic_Compare:
		return "CMP"
	case Mnemonic_ClearOverflowFlag:
		return "CLV"
	case Mnemonic_BranchifCarryClear:
		return "BCC"
	case Mnemonic_IncrementYRegister:
		return "INY"
	case Mnemonic_ArithmeticShiftLeft:
		return "ASL"
	case Mnemonic_BitTest:
		return "BIT"
	case Mnemonic_LogicalShiftRight:
		return "LSR"
	case Mnemonic_BranchifEqual:
		return "BEQ"
	case Mnemonic_LoadXRegister:
		return "LDX"
	case Mnemonic_TransferAccumulatortoX:
		return "TAX"
	case Mnemonic_TransferAccumulatortoY:
		return "TAY"
	case Mnemonic_TransferStackPointertoX:
		return "TSX"
	case Mnemonic_StoreAccumulator:
		return "STA"
	case Mnemonic_SetInterruptDisable:
		return "SEI"
	case Mnemonic_ClearInterruptDisable:
		return "CLI"
	case Mnemonic_SubtractwithCarry:
		return "SBC"
	case Mnemonic_TransferXtoAccumulator:
		return "TXA"
	case Mnemonic_LoadYRegister:
		return "LDY"
	case Mnemonic_ClearCarryFlag:
		return "CLC"
	case Mnemonic_StoreXRegister:
		return "STX"
	case Mnemonic_Jump:
		return "JMP"
	case Mnemonic_IncrementMemory:
		return "INC"
	case Mnemonic_StoreYRegister:
		return "STY"
	case Mnemonic_CompareXRegister:
		return "CPX"
	case Mnemonic_BranchifMinus:
		return "BMI"
	case Mnemonic_DecrementMemory:
		return "DEC"
	case Mnemonic_BranchifPlus:
		return "BPL"
	case Mnemonic_PushProcessorStatus:
		return "PHP"
	case Mnemonic_SetDecimalFlag:
		return "SED"
	case Mnemonic_AddwithCarry:
		return "ADC"
	case Mnemonic_LoadAccumulator:
		return "LDA"
	case Mnemonic_ExclusiveOR:
		return "EOR"
	case Mnemonic_LogicalOR:
		return "ORA"
	case Mnemonic_PushAccumulator:
		return "PHA"
	case Mnemonic_ReturnfromInterrupt:
		return "RTI"
	case Mnemonic_NoOperation:
		return "NOP"
	case Mnemonic_TransferXtoStackPointer:
		return "TXS"
	case Mnemonic_DecrementXRegister:
		return "DEX"
	case Mnemonic_JumptoSubroutine:
		return "JSR"
	case Mnemonic_BranchifOverflowSet:
		return "BVS"
	case Mnemonic_SetCarryFlag:
		return "SEC"
	}
	return "???"
}

type MnemonicCategory int

const (
	MnemonicCategory_Arithmetic  MnemonicCategory = iota // arith
	MnemonicCategory_Logic       MnemonicCategory = iota // logic
	MnemonicCategory_Flags       MnemonicCategory = iota // flags
	MnemonicCategory_Kill        MnemonicCategory = iota // kil
	MnemonicCategory_Shift       MnemonicCategory = iota // shift
	MnemonicCategory_Branch      MnemonicCategory = iota // bra
	MnemonicCategory_Control     MnemonicCategory = iota // ctrl
	MnemonicCategory_Increment   MnemonicCategory = iota // inc
	MnemonicCategory_Load        MnemonicCategory = iota // load
	MnemonicCategory_NoOperation MnemonicCategory = iota // nop
	MnemonicCategory_Stack       MnemonicCategory = iota // stack
	MnemonicCategory_Transfer    MnemonicCategory = iota // trans
)

func (c MnemonicCategory) String() string {
	switch c {
	case MnemonicCategory_NoOperation:
		return "nop"
	case MnemonicCategory_Stack:
		return "stack"
	case MnemonicCategory_Transfer:
		return "trans"
	case MnemonicCategory_Shift:
		return "shift"
	case MnemonicCategory_Branch:
		return "bra"
	case MnemonicCategory_Control:
		return "ctrl"
	case MnemonicCategory_Increment:
		return "inc"
	case MnemonicCategory_Load:
		return "load"
	case MnemonicCategory_Arithmetic:
		return "arith"
	case MnemonicCategory_Logic:
		return "logic"
	case MnemonicCategory_Flags:
		return "flags"
	case MnemonicCategory_Kill:
		return "kil"
	}
	return "???"
}

type AddressingMode int

const (
	AddressingMode_Implied                  AddressingMode = iota // -
	AddressingMode_Accumulator              AddressingMode = iota // A
	AddressingMode_Immediate                AddressingMode = iota // #d8
	AddressingMode_ZeroPage                 AddressingMode = iota // a8
	AddressingMode_XIndexedZeroPage         AddressingMode = iota // a8,X
	AddressingMode_YIndexedZeroPage         AddressingMode = iota // a8,Y
	AddressingMode_XIndexedZeroPageIndirect AddressingMode = iota // (a8,X)
	AddressingMode_ZeroPageIndirectYIndexed AddressingMode = iota // (a8),Y
	AddressingMode_Absolute                 AddressingMode = iota // a16
	AddressingMode_XIndexedAbsolute         AddressingMode = iota // a16,X
	AddressingMode_YIndexedAbsolute         AddressingMode = iota // a16,Y
	AddressingMode_AbsoluteIndirect         AddressingMode = iota // (a16)
	AddressingMode_Relative                 AddressingMode = iota // r8
)

func (a AddressingMode) String() string {
	switch a {
	case AddressingMode_Implied:
		return "-"
	case AddressingMode_Accumulator:
		return "A"
	case AddressingMode_XIndexedZeroPage:
		return "a8,X"
	case AddressingMode_YIndexedZeroPage:
		return "a8,Y"
	case AddressingMode_XIndexedZeroPageIndirect:
		return "(a8,X)"
	case AddressingMode_Absolute:
		return "a16"
	case AddressingMode_YIndexedAbsolute:
		return "a16,Y"
	case AddressingMode_AbsoluteIndirect:
		return "(a16)"
	case AddressingMode_Relative:
		return "r8"
	case AddressingMode_Immediate:
		return "#d8"
	case AddressingMode_ZeroPage:
		return "a8"
	case AddressingMode_ZeroPageIndirectYIndexed:
		return "(a8),Y"
	case AddressingMode_XIndexedAbsolute:
		return "a16,X"
	}
	return "???"
}

type Instruction struct {
	Mnemonic       Mnemonic
	AddressingMode AddressingMode
	Category       MnemonicCategory
	AffectedFlags  uint8
	Size           uint16
	Opcode         byte
}

var ErrInstructionNotFound = errors.New("Instruction not found")

func Decode(opcode byte) (Instruction, error) {
	switch opcode {
	case uint8(0x0):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_InterruptDisable,
			Category:       MnemonicCategory_Control,
			Mnemonic:       Mnemonic_ForceInterrupt,
			Opcode:         uint8(0x0),
			Size:           1,
		}, nil
	case uint8(0x1):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x1),
			Size:           2,
		}, nil
	case uint8(0x5):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x5),
			Size:           2,
		}, nil
	case uint8(0x6):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_ArithmeticShiftLeft,
			Opcode:         uint8(0x6),
			Size:           2,
		}, nil
	case uint8(0x8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			Category:       MnemonicCategory_Stack,
			Mnemonic:       Mnemonic_PushProcessorStatus,
			Opcode:         uint8(0x8),
			Size:           1,
		}, nil
	case uint8(0x9):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x9),
			Size:           2,
		}, nil
	case uint8(0xa):
		return Instruction{
			AddressingMode: AddressingMode_Accumulator,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_ArithmeticShiftLeft,
			Opcode:         uint8(0xa),
			Size:           1,
		}, nil
	case uint8(0xd):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0xd),
			Size:           3,
		}, nil
	case uint8(0xe):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_ArithmeticShiftLeft,
			Opcode:         uint8(0xe),
			Size:           3,
		}, nil
	case uint8(0x10):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifPlus,
			Opcode:         uint8(0x10),
			Size:           2,
		}, nil
	case uint8(0x11):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x11),
			Size:           2,
		}, nil
	case uint8(0x15):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x15),
			Size:           2,
		}, nil
	case uint8(0x16):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_ArithmeticShiftLeft,
			Opcode:         uint8(0x16),
			Size:           2,
		}, nil
	case uint8(0x18):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Carry,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_ClearCarryFlag,
			Opcode:         uint8(0x18),
			Size:           1,
		}, nil
	case uint8(0x19):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x19),
			Size:           3,
		}, nil
	case uint8(0x1d):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalOR,
			Opcode:         uint8(0x1d),
			Size:           3,
		}, nil
	case uint8(0x1e):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_ArithmeticShiftLeft,
			Opcode:         uint8(0x1e),
			Size:           3,
		}, nil
	case uint8(0x20):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			Category:       MnemonicCategory_Control,
			Mnemonic:       Mnemonic_JumptoSubroutine,
			Opcode:         uint8(0x20),
			Size:           3,
		}, nil
	case uint8(0x21):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x21),
			Size:           2,
		}, nil
	case uint8(0x24):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_BitTest,
			Opcode:         uint8(0x24),
			Size:           2,
		}, nil
	case uint8(0x25):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x25),
			Size:           2,
		}, nil
	case uint8(0x26):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateLeft,
			Opcode:         uint8(0x26),
			Size:           2,
		}, nil
	case uint8(0x28):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Decimal | Status_InterruptDisable | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Stack,
			Mnemonic:       Mnemonic_PullProcessorStatus,
			Opcode:         uint8(0x28),
			Size:           1,
		}, nil
	case uint8(0x29):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x29),
			Size:           2,
		}, nil
	case uint8(0x2a):
		return Instruction{
			AddressingMode: AddressingMode_Accumulator,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateLeft,
			Opcode:         uint8(0x2a),
			Size:           1,
		}, nil
	case uint8(0x2c):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_BitTest,
			Opcode:         uint8(0x2c),
			Size:           3,
		}, nil
	case uint8(0x2d):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x2d),
			Size:           3,
		}, nil
	case uint8(0x2e):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateLeft,
			Opcode:         uint8(0x2e),
			Size:           3,
		}, nil
	case uint8(0x30):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifMinus,
			Opcode:         uint8(0x30),
			Size:           2,
		}, nil
	case uint8(0x31):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x31),
			Size:           2,
		}, nil
	case uint8(0x35):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x35),
			Size:           2,
		}, nil
	case uint8(0x36):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateLeft,
			Opcode:         uint8(0x36),
			Size:           2,
		}, nil
	case uint8(0x38):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Carry,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_SetCarryFlag,
			Opcode:         uint8(0x38),
			Size:           1,
		}, nil
	case uint8(0x39):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x39),
			Size:           3,
		}, nil
	case uint8(0x3d):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_LogicalAND,
			Opcode:         uint8(0x3d),
			Size:           3,
		}, nil
	case uint8(0x3e):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateLeft,
			Opcode:         uint8(0x3e),
			Size:           3,
		}, nil
	case uint8(0x40):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Decimal | Status_InterruptDisable | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Control,
			Mnemonic:       Mnemonic_ReturnfromInterrupt,
			Opcode:         uint8(0x40),
			Size:           1,
		}, nil
	case uint8(0x41):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x41),
			Size:           2,
		}, nil
	case uint8(0x45):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x45),
			Size:           2,
		}, nil
	case uint8(0x46):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_LogicalShiftRight,
			Opcode:         uint8(0x46),
			Size:           2,
		}, nil
	case uint8(0x48):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			Category:       MnemonicCategory_Stack,
			Mnemonic:       Mnemonic_PushAccumulator,
			Opcode:         uint8(0x48),
			Size:           1,
		}, nil
	case uint8(0x49):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x49),
			Size:           2,
		}, nil
	case uint8(0x4a):
		return Instruction{
			AddressingMode: AddressingMode_Accumulator,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_LogicalShiftRight,
			Opcode:         uint8(0x4a),
			Size:           1,
		}, nil
	case uint8(0x4c):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			Category:       MnemonicCategory_Control,
			Mnemonic:       Mnemonic_Jump,
			Opcode:         uint8(0x4c),
			Size:           3,
		}, nil
	case uint8(0x4d):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x4d),
			Size:           3,
		}, nil
	case uint8(0x4e):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_LogicalShiftRight,
			Opcode:         uint8(0x4e),
			Size:           3,
		}, nil
	case uint8(0x50):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifOverflowClear,
			Opcode:         uint8(0x50),
			Size:           2,
		}, nil
	case uint8(0x51):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x51),
			Size:           2,
		}, nil
	case uint8(0x55):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x55),
			Size:           2,
		}, nil
	case uint8(0x56):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_LogicalShiftRight,
			Opcode:         uint8(0x56),
			Size:           2,
		}, nil
	case uint8(0x58):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_InterruptDisable,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_ClearInterruptDisable,
			Opcode:         uint8(0x58),
			Size:           1,
		}, nil
	case uint8(0x59):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x59),
			Size:           3,
		}, nil
	case uint8(0x5d):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Logic,
			Mnemonic:       Mnemonic_ExclusiveOR,
			Opcode:         uint8(0x5d),
			Size:           3,
		}, nil
	case uint8(0x5e):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_LogicalShiftRight,
			Opcode:         uint8(0x5e),
			Size:           3,
		}, nil
	case uint8(0x60):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			Category:       MnemonicCategory_Control,
			Mnemonic:       Mnemonic_ReturnfromSubroutine,
			Opcode:         uint8(0x60),
			Size:           1,
		}, nil
	case uint8(0x61):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x61),
			Size:           2,
		}, nil
	case uint8(0x65):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x65),
			Size:           2,
		}, nil
	case uint8(0x66):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateRight,
			Opcode:         uint8(0x66),
			Size:           2,
		}, nil
	case uint8(0x68):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Stack,
			Mnemonic:       Mnemonic_PullAccumulator,
			Opcode:         uint8(0x68),
			Size:           1,
		}, nil
	case uint8(0x69):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x69),
			Size:           2,
		}, nil
	case uint8(0x6a):
		return Instruction{
			AddressingMode: AddressingMode_Accumulator,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateRight,
			Opcode:         uint8(0x6a),
			Size:           1,
		}, nil
	case uint8(0x6c):
		return Instruction{
			AddressingMode: AddressingMode_AbsoluteIndirect,
			Category:       MnemonicCategory_Control,
			Mnemonic:       Mnemonic_Jump,
			Opcode:         uint8(0x6c),
			Size:           3,
		}, nil
	case uint8(0x6d):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x6d),
			Size:           3,
		}, nil
	case uint8(0x6e):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateRight,
			Opcode:         uint8(0x6e),
			Size:           3,
		}, nil
	case uint8(0x70):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifOverflowSet,
			Opcode:         uint8(0x70),
			Size:           2,
		}, nil
	case uint8(0x71):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x71),
			Size:           2,
		}, nil
	case uint8(0x75):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x75),
			Size:           2,
		}, nil
	case uint8(0x76):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateRight,
			Opcode:         uint8(0x76),
			Size:           2,
		}, nil
	case uint8(0x78):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_InterruptDisable,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_SetInterruptDisable,
			Opcode:         uint8(0x78),
			Size:           1,
		}, nil
	case uint8(0x79):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x79),
			Size:           3,
		}, nil
	case uint8(0x7d):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_AddwithCarry,
			Opcode:         uint8(0x7d),
			Size:           3,
		}, nil
	case uint8(0x7e):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Shift,
			Mnemonic:       Mnemonic_RotateRight,
			Opcode:         uint8(0x7e),
			Size:           3,
		}, nil
	case uint8(0x81):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x81),
			Size:           2,
		}, nil
	case uint8(0x84):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreYRegister,
			Opcode:         uint8(0x84),
			Size:           2,
		}, nil
	case uint8(0x85):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x85),
			Size:           2,
		}, nil
	case uint8(0x86):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreXRegister,
			Opcode:         uint8(0x86),
			Size:           2,
		}, nil
	case uint8(0x88):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_DecrementYRegister,
			Opcode:         uint8(0x88),
			Size:           1,
		}, nil
	case uint8(0x8a):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Transfer,
			Mnemonic:       Mnemonic_TransferXtoAccumulator,
			Opcode:         uint8(0x8a),
			Size:           1,
		}, nil
	case uint8(0x8c):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreYRegister,
			Opcode:         uint8(0x8c),
			Size:           3,
		}, nil
	case uint8(0x8d):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x8d),
			Size:           3,
		}, nil
	case uint8(0x8e):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreXRegister,
			Opcode:         uint8(0x8e),
			Size:           3,
		}, nil
	case uint8(0x90):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifCarryClear,
			Opcode:         uint8(0x90),
			Size:           2,
		}, nil
	case uint8(0x91):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x91),
			Size:           2,
		}, nil
	case uint8(0x94):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreYRegister,
			Opcode:         uint8(0x94),
			Size:           2,
		}, nil
	case uint8(0x95):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x95),
			Size:           2,
		}, nil
	case uint8(0x96):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedZeroPage,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreXRegister,
			Opcode:         uint8(0x96),
			Size:           2,
		}, nil
	case uint8(0x98):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Transfer,
			Mnemonic:       Mnemonic_TransferYtoAccumulator,
			Opcode:         uint8(0x98),
			Size:           1,
		}, nil
	case uint8(0x99):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x99),
			Size:           3,
		}, nil
	case uint8(0x9a):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			Category:       MnemonicCategory_Transfer,
			Mnemonic:       Mnemonic_TransferXtoStackPointer,
			Opcode:         uint8(0x9a),
			Size:           1,
		}, nil
	case uint8(0x9d):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_StoreAccumulator,
			Opcode:         uint8(0x9d),
			Size:           3,
		}, nil
	case uint8(0xa0):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadYRegister,
			Opcode:         uint8(0xa0),
			Size:           2,
		}, nil
	case uint8(0xa1):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xa1),
			Size:           2,
		}, nil
	case uint8(0xa2):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadXRegister,
			Opcode:         uint8(0xa2),
			Size:           2,
		}, nil
	case uint8(0xa4):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadYRegister,
			Opcode:         uint8(0xa4),
			Size:           2,
		}, nil
	case uint8(0xa5):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xa5),
			Size:           2,
		}, nil
	case uint8(0xa6):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadXRegister,
			Opcode:         uint8(0xa6),
			Size:           2,
		}, nil
	case uint8(0xa8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Transfer,
			Mnemonic:       Mnemonic_TransferAccumulatortoY,
			Opcode:         uint8(0xa8),
			Size:           1,
		}, nil
	case uint8(0xa9):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xa9),
			Size:           2,
		}, nil
	case uint8(0xaa):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Transfer,
			Mnemonic:       Mnemonic_TransferAccumulatortoX,
			Opcode:         uint8(0xaa),
			Size:           1,
		}, nil
	case uint8(0xac):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadYRegister,
			Opcode:         uint8(0xac),
			Size:           3,
		}, nil
	case uint8(0xad):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xad),
			Size:           3,
		}, nil
	case uint8(0xae):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadXRegister,
			Opcode:         uint8(0xae),
			Size:           3,
		}, nil
	case uint8(0xb0):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifCarrySet,
			Opcode:         uint8(0xb0),
			Size:           2,
		}, nil
	case uint8(0xb1):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xb1),
			Size:           2,
		}, nil
	case uint8(0xb4):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadYRegister,
			Opcode:         uint8(0xb4),
			Size:           2,
		}, nil
	case uint8(0xb5):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xb5),
			Size:           2,
		}, nil
	case uint8(0xb6):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadXRegister,
			Opcode:         uint8(0xb6),
			Size:           2,
		}, nil
	case uint8(0xb8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Overflow,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_ClearOverflowFlag,
			Opcode:         uint8(0xb8),
			Size:           1,
		}, nil
	case uint8(0xb9):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xb9),
			Size:           3,
		}, nil
	case uint8(0xba):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Transfer,
			Mnemonic:       Mnemonic_TransferStackPointertoX,
			Opcode:         uint8(0xba),
			Size:           1,
		}, nil
	case uint8(0xbc):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadYRegister,
			Opcode:         uint8(0xbc),
			Size:           3,
		}, nil
	case uint8(0xbd):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadAccumulator,
			Opcode:         uint8(0xbd),
			Size:           3,
		}, nil
	case uint8(0xbe):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Load,
			Mnemonic:       Mnemonic_LoadXRegister,
			Opcode:         uint8(0xbe),
			Size:           3,
		}, nil
	case uint8(0xc0):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_CompareYRegister,
			Opcode:         uint8(0xc0),
			Size:           2,
		}, nil
	case uint8(0xc1):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xc1),
			Size:           2,
		}, nil
	case uint8(0xc4):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_CompareYRegister,
			Opcode:         uint8(0xc4),
			Size:           2,
		}, nil
	case uint8(0xc5):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xc5),
			Size:           2,
		}, nil
	case uint8(0xc6):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_DecrementMemory,
			Opcode:         uint8(0xc6),
			Size:           2,
		}, nil
	case uint8(0xc8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_IncrementYRegister,
			Opcode:         uint8(0xc8),
			Size:           1,
		}, nil
	case uint8(0xc9):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xc9),
			Size:           2,
		}, nil
	case uint8(0xca):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_DecrementXRegister,
			Opcode:         uint8(0xca),
			Size:           1,
		}, nil
	case uint8(0xcc):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_CompareYRegister,
			Opcode:         uint8(0xcc),
			Size:           3,
		}, nil
	case uint8(0xcd):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xcd),
			Size:           3,
		}, nil
	case uint8(0xce):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_DecrementMemory,
			Opcode:         uint8(0xce),
			Size:           3,
		}, nil
	case uint8(0xd0):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifNotEqual,
			Opcode:         uint8(0xd0),
			Size:           2,
		}, nil
	case uint8(0xd1):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xd1),
			Size:           2,
		}, nil
	case uint8(0xd5):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xd5),
			Size:           2,
		}, nil
	case uint8(0xd6):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_DecrementMemory,
			Opcode:         uint8(0xd6),
			Size:           2,
		}, nil
	case uint8(0xd8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Decimal,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_ClearDecimalMode,
			Opcode:         uint8(0xd8),
			Size:           1,
		}, nil
	case uint8(0xd9):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xd9),
			Size:           3,
		}, nil
	case uint8(0xdd):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_Compare,
			Opcode:         uint8(0xdd),
			Size:           3,
		}, nil
	case uint8(0xde):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_DecrementMemory,
			Opcode:         uint8(0xde),
			Size:           3,
		}, nil
	case uint8(0xe0):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_CompareXRegister,
			Opcode:         uint8(0xe0),
			Size:           2,
		}, nil
	case uint8(0xe1):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPageIndirect,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xe1),
			Size:           2,
		}, nil
	case uint8(0xe4):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_CompareXRegister,
			Opcode:         uint8(0xe4),
			Size:           2,
		}, nil
	case uint8(0xe5):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xe5),
			Size:           2,
		}, nil
	case uint8(0xe6):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_IncrementMemory,
			Opcode:         uint8(0xe6),
			Size:           2,
		}, nil
	case uint8(0xe8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_IncrementXRegister,
			Opcode:         uint8(0xe8),
			Size:           1,
		}, nil
	case uint8(0xe9):
		return Instruction{
			AddressingMode: AddressingMode_Immediate,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xe9),
			Size:           2,
		}, nil
	case uint8(0xea):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			Category:       MnemonicCategory_NoOperation,
			Mnemonic:       Mnemonic_NoOperation,
			Opcode:         uint8(0xea),
			Size:           1,
		}, nil
	case uint8(0xec):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_CompareXRegister,
			Opcode:         uint8(0xec),
			Size:           3,
		}, nil
	case uint8(0xed):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xed),
			Size:           3,
		}, nil
	case uint8(0xee):
		return Instruction{
			AddressingMode: AddressingMode_Absolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_IncrementMemory,
			Opcode:         uint8(0xee),
			Size:           3,
		}, nil
	case uint8(0xf0):
		return Instruction{
			AddressingMode: AddressingMode_Relative,
			Category:       MnemonicCategory_Branch,
			Mnemonic:       Mnemonic_BranchifEqual,
			Opcode:         uint8(0xf0),
			Size:           2,
		}, nil
	case uint8(0xf1):
		return Instruction{
			AddressingMode: AddressingMode_ZeroPageIndirectYIndexed,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xf1),
			Size:           2,
		}, nil
	case uint8(0xf5):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xf5),
			Size:           2,
		}, nil
	case uint8(0xf6):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedZeroPage,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_IncrementMemory,
			Opcode:         uint8(0xf6),
			Size:           2,
		}, nil
	case uint8(0xf8):
		return Instruction{
			AddressingMode: AddressingMode_Implied,
			AffectedFlags:  Status_Decimal,
			Category:       MnemonicCategory_Flags,
			Mnemonic:       Mnemonic_SetDecimalFlag,
			Opcode:         uint8(0xf8),
			Size:           1,
		}, nil
	case uint8(0xf9):
		return Instruction{
			AddressingMode: AddressingMode_YIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xf9),
			Size:           3,
		}, nil
	case uint8(0xfd):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Overflow | Status_Zero | Status_Carry,
			Category:       MnemonicCategory_Arithmetic,
			Mnemonic:       Mnemonic_SubtractwithCarry,
			Opcode:         uint8(0xfd),
			Size:           3,
		}, nil
	case uint8(0xfe):
		return Instruction{
			AddressingMode: AddressingMode_XIndexedAbsolute,
			AffectedFlags:  Status_Negative | Status_Zero,
			Category:       MnemonicCategory_Increment,
			Mnemonic:       Mnemonic_IncrementMemory,
			Opcode:         uint8(0xfe),
			Size:           3,
		}, nil
	}
	return Instruction{}, ErrInstructionNotFound
}
