package main

import (
	"flag"
	"fmt"
	"strings"
	"time"
	"unicode"

	"github.com/dave/jennifer/jen"
)

var (
	packageName string
	outfile     string
)

func init() {
	flag.StringVar(&packageName, "package", "", "The name of the package to generate")
	flag.StringVar(&outfile, "outfile", "", "The name of the file to write to")
	flag.Parse()
}

func stripIllegalIdentifierCharacters(str string) string {
	return strings.Map(func(r rune) rune {
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) && r != '_' {
			return -1
		}
		return r
	}, str)
}

func main() {
	cpu := parseCpuDescription()
	instructionCategoryByMnemonic := map[string]string{}
	affectedFlagsByMnemonic := map[string][]int{}
	for _, operation := range cpu.Operations {
		instructionCategoryByMnemonic[operation.Mnemonic] = operation.Category
		affectedFlagsByMnemonic[operation.Mnemonic] = operation.Flags
	}

	instructionSizeByAddressingMode := map[string]int{}
	for _, addressingMode := range cpu.AddressingModes {
		instructionSizeByAddressingMode[addressingMode.SyntaxNovel] = addressingMode.Size
	}

	f := jen.NewFile(packageName)
	f.HeaderComment(fmt.Sprintf("Code generated on %s; DO NOT EDIT.", time.Now().Format(time.RFC3339)))
	f.HeaderComment(fmt.Sprintf("%s %s (%s)", cpu.Info.Manufacturer, cpu.Info.Name, cpu.Info.Year))

	// Registers
	f.Type().Id("Registers").StructFunc(func(g *jen.Group) {
		for _, reg := range cpu.Registers {
			g.Id(reg.Name).Do(func(s *jen.Statement) {
				switch reg.Size {
				case 8:
					s.Uint8()
				case 16:
					s.Uint16()
				}
			}).Comment(reg.Description)
		}
	})

	// Vectors
	vectorType := "Vector"
	vectorPrefix := "Vector_"
	f.Type().Id(vectorType).Uint16()
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, vector := range cpu.Vectors {
			g.Id(vectorPrefix + vector.Name).Id(vectorType).Op("=").Lit(vector.Address)
		}
	})

	// Status Flags
	flagType := "StatusFlag"
	flagPrefix := "Status_"
	flagIdentifiersByBit := map[int]string{}
	f.Type().Id(flagType).Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, flag := range cpu.Flags {
			if flag.Identifier == "-" {
				continue
			}
			flagIdentifiersByBit[flag.Bit] = flagPrefix + flag.Identifier
			g.Id(flagPrefix + flag.Identifier).Id(flagType).Op("=").Lit(1).Op("<<").Lit(flag.Bit).Comment(flag.Description)
		}
	})

	// Mnemonics
	mnemonicType := "Mnemonic"
	mnemonicPrefix := "Mnemonic_"
	mnemonicIdentifiers := map[string]string{} // Used to map short to generated identifier
	f.Type().Id(mnemonicType).Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, mnemonic := range cpu.Mnemonics {
			mnemonicID := mnemonicPrefix + stripIllegalIdentifierCharacters(mnemonic.Description)
			mnemonicIdentifiers[mnemonic.Short] = mnemonicID

			g.Id(mnemonicID).Id(mnemonicType).Op("=").Iota().Comment(mnemonic.Short)
		}
	})

	// mnemonic to string
	f.Func().Params(jen.Id("m").Id(mnemonicType)).Id("String").Params().String().Block(
		jen.Switch(jen.Id("m")).BlockFunc(func(s *jen.Group) {
			for short, identifier := range mnemonicIdentifiers {
				s.Case(jen.Id(identifier)).Block(jen.Return(jen.Lit(short)))
			}
		}),
	)

	// Mnemonic Category
	mnemonicCategoryType := "MnemonicCategory"
	mnemonicCategoryPrefix := "MnemonicCategory_"
	mnemonicCategoryIdentifiers := map[string]string{
		"arith": mnemonicCategoryPrefix + "Arithmetic",
		"logic": mnemonicCategoryPrefix + "Logic",
		"shift": mnemonicCategoryPrefix + "Shift",
		"bra":   mnemonicCategoryPrefix + "Branch",
		"ctrl":  mnemonicCategoryPrefix + "Control",
		"flags": mnemonicCategoryPrefix + "Flags",
		"inc":   mnemonicCategoryPrefix + "Increment",
		"load":  mnemonicCategoryPrefix + "Load",
		"nop":   mnemonicCategoryPrefix + "NoOperation",
		"stack": mnemonicCategoryPrefix + "Stack",
		"trans": mnemonicCategoryPrefix + "Transfer",
		"kil":   mnemonicCategoryPrefix + "Kill", // Technically only used for an illegal instruction
	}
	f.Type().Id(mnemonicCategoryType).Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		for short, identifier := range mnemonicCategoryIdentifiers {
			g.Id(identifier).Id(mnemonicCategoryType).Op("=").Iota().Comment(short)
		}
	})

	// mnemonic category to string
	f.Func().Params(jen.Id("c").Id(mnemonicCategoryType)).Id("String").Params().String().Block(
		jen.Switch(jen.Id("c")).BlockFunc(func(s *jen.Group) {
			for short, identifier := range mnemonicCategoryIdentifiers {
				s.Case(jen.Id(identifier)).Block(jen.Return(jen.Lit(short)))
			}
		}),
	)

	// Addressing Modes
	addressingModeType := "AddressingMode"
	addressingModePrefix := "AddressingMode_"
	addressingModeIdentifiers := map[string]string{}
	f.Type().Id(addressingModeType).Int()
	f.Const().DefsFunc(func(g *jen.Group) {
		for _, addressingMode := range cpu.AddressingModes {
			addressingModeID := addressingModePrefix + stripIllegalIdentifierCharacters(addressingMode.Description)
			addressingModeIdentifiers[addressingMode.SyntaxNovel] = addressingModeID

			g.Id(addressingModeID).Id(addressingModeType).Op("=").Iota().Comment(addressingMode.SyntaxNovel)
		}
	})

	// addressing mode to string
	f.Func().Params(jen.Id("a").Id(addressingModeType)).Id("String").Params().String().Block(
		jen.Switch(jen.Id("a")).BlockFunc(func(s *jen.Group) {
			for syntax, identifier := range addressingModeIdentifiers {
				s.Case(jen.Id(identifier)).Block(jen.Return(jen.Lit(syntax)))
			}
		}),
	)

	// Instruction
	instructionType := "Instruction"
	f.Type().Id(instructionType).StructFunc(func(g *jen.Group) {
		g.Id("Mnemonic").Id(mnemonicType)
		g.Id("AddressingMode").Id(addressingModeType)
		g.Id("Category").Id(mnemonicCategoryType)
		g.Id("AffectedFlags").Id(flagType)
		g.Id("Size").Int()
		g.Id("Opcode").Byte()
	})

	// decode instruction
	f.Func().Id("Decode").Params(jen.Id("opcode").Byte()).Params(jen.Id(instructionType), jen.Error()).Block(
		jen.Switch(jen.Id("opcode")).BlockFunc(func(s *jen.Group) {
			for _, opcode := range cpu.Opcodes {
				if opcode.Illegal {
					continue
				}

				mnemonicIdentifier, ok := mnemonicIdentifiers[opcode.Mnemonic]
				if !ok {
					panic(fmt.Sprintf("mnemonic %s not found", opcode.Mnemonic))
				}

				addressingModeIdentifier, ok := addressingModeIdentifiers[opcode.AddressingMode]
				if !ok {
					panic(fmt.Sprintf("addressing mode %s not found", opcode.AddressingMode))
				}

				categoryShort, ok := instructionCategoryByMnemonic[opcode.Mnemonic]
				if !ok {
					panic(fmt.Sprintf("category for mnemonic %s not found", opcode.Mnemonic))
				}

				mnemonicCategoryIdentifier, ok := mnemonicCategoryIdentifiers[categoryShort]
				if !ok {
					panic(fmt.Sprintf("mnemonic category %s not found", opcode.Mnemonic))
				}

				s.Case(jen.Lit(opcode.Op)).Block(
					jen.Return(jen.Id(instructionType).Values(jen.Dict{
						jen.Id("Mnemonic"):       jen.Id(mnemonicIdentifier),
						jen.Id("AddressingMode"): jen.Id(addressingModeIdentifier),
						jen.Id("Category"):       jen.Id(mnemonicCategoryIdentifier),
						jen.Id("AffectedFlags"): jen.Do(func(s *jen.Statement) {
							affectedFlags := affectedFlagsByMnemonic[opcode.Mnemonic]
							for i, flag := range affectedFlags {
								if i == len(affectedFlags)-1 {
									s.Id(flagIdentifiersByBit[flag])
								} else {
									s.Id(flagIdentifiersByBit[flag]).Op("|")
								}
							}
						}),
						jen.Id("Size"):   jen.Lit(instructionSizeByAddressingMode[opcode.AddressingMode]),
						jen.Id("Opcode"): jen.Lit(opcode.Op),
					})),
				)
			}
		}),
	)

	err := f.Save(outfile)
	if err != nil {
		panic(err)
	}
}
