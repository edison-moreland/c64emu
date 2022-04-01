package main

import (
	"bufio"
	"embed"
	"encoding/binary"
	"encoding/hex"
	"io/fs"
	"regexp"
	"strconv"
	"strings"
)

const fieldSeparator = "|"

//go:embed cpu_6502.txt
var f embed.FS

type CPUInfo struct {
	Manufacturer string
	Name         string
	Year         string
	ID           string
	Description  string
}

type CPURegister struct {
	Name        string
	Size        int
	Description string
}

type CPUFlag struct {
	Bit         int
	Identifier  string
	Description string
}

type CPUMnemonic struct {
	Short       string
	Description string
}

type CPUOperation struct {
	Mnemonic  string
	Category  string
	Flags     []int
	Operation string
}

type CPUAddressingMode struct {
	SyntaxNovel string
	Size        int
	Syntax      string
	Description string
}

type CPUOpcode struct {
	Op             byte
	Illegal        bool
	Mnemonic       string
	AddressingMode string
}

type CPUVector struct {
	Name    string
	Address uint16
}

type CPU6502 struct {
	Info            CPUInfo
	Registers       []CPURegister
	Flags           []CPUFlag
	Mnemonics       []CPUMnemonic
	Operations      []CPUOperation
	AddressingModes []CPUAddressingMode
	Opcodes         []CPUOpcode
	Vectors         []CPUVector
}

func parseCpuDescription() CPU6502 {
	infoFile, err := f.Open("cpu_6502.txt")
	if err != nil {
		panic(err)
	}
	defer infoFile.Close()

	sections := parseFirstPass(infoFile)

	cpu := CPU6502{
		Info:            parseInfo(sections["info"]),
		Registers:       parseRegisters(sections["registers"]),
		Flags:           parseFlags(sections["flags"]),
		Mnemonics:       parseMnemos(sections["mnemos"]),
		Operations:      parseOperations(sections["operations"]),
		AddressingModes: parseAddressingModes(sections["addmodes"]),
		Opcodes:         parseOpcodes(sections["opcodes"]),
		Vectors:         parseVectors(sections["vectors"]),
	}

	return cpu
}

// parseFirstPass buckets the config by section.
func parseFirstPass(file fs.File) map[string][]string {
	sections := make(map[string][]string)
	currentSection := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}
		if line[0] == '#' {
			continue
		}
		if line[0] == '[' {
			// Sections start with [<sectionName>]
			currentSection = line[1 : len(line)-1]
			continue
		}

		// remove any comments from the middle of the line
		stripped := strings.TrimRight(strings.SplitN(line, "#", 2)[0], " ")

		// Next, any group of two or more spaces needs to be turned into a separator.
		var re = regexp.MustCompile(`(?m)([^\s])( {2,})`)
		seperated := re.ReplaceAllString(stripped, "$1"+fieldSeparator)

		sections[currentSection] = append(sections[currentSection], seperated)
	}
	return sections
}

// parseSectionToMap parses a section into a map of key/value pairs.
func parseSectionToMap(section []string) map[string]string {
	sectionMap := make(map[string]string)
	for _, line := range section {
		// Split on the first whitespace.
		split := strings.SplitN(line, fieldSeparator, 2)
		key := strings.TrimSpace(split[0])
		value := strings.TrimSpace(split[1])

		sectionMap[key] = value
	}
	return sectionMap
}

func parseInfo(section []string) CPUInfo {
	infoMap := parseSectionToMap(section)
	info := CPUInfo{
		Manufacturer: infoMap["manufacturer"],
		Name:         infoMap["name"],
		Year:         infoMap["year"],
		ID:           infoMap["id"],
		Description:  infoMap["description"],
	}
	return info
}

func parseRegisters(section []string) []CPURegister {
	registers := []CPURegister{}
	for _, line := range section {
		// Split on the first whitespace.
		split := strings.SplitN(line, fieldSeparator, 3)
		name := strings.TrimSpace(split[0])
		desc := strings.TrimSpace(split[2])
		size, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			panic(err)
		}

		registers = append(registers, CPURegister{
			Name:        name,
			Size:        size,
			Description: desc,
		})
	}
	return registers
}

func parseFlags(section []string) []CPUFlag {
	flags := []CPUFlag{}
	for _, line := range section {
		// Split on the first whitespace.
		split := strings.SplitN(line, fieldSeparator, 3)
		bit, err := strconv.Atoi(strings.TrimSpace(split[0]))
		if err != nil {
			panic(err)
		}
		identifier := strings.TrimSpace(split[1])
		desc := strings.TrimSpace(split[2])

		flags = append(flags, CPUFlag{
			Bit:         bit,
			Identifier:  identifier,
			Description: desc,
		})
	}
	return flags
}

func parseMnemos(section []string) []CPUMnemonic {
	mnemonics := []CPUMnemonic{}

	mnemoMap := parseSectionToMap(section)
	for key, value := range mnemoMap {
		mnemonics = append(mnemonics, CPUMnemonic{
			Short:       key,
			Description: value,
		})
	}

	return mnemonics
}

func parseOperations(section []string) []CPUOperation {
	operations := []CPUOperation{}
	for _, line := range section {
		split := strings.SplitN(line, fieldSeparator, 4)
		mnemo := strings.TrimSpace(split[0])
		type_ := strings.TrimSpace(split[1])
		operation := strings.TrimSpace(split[3])

		// fmt.Println(split)

		flags := strings.TrimSpace(split[2])
		affectedFlags := []int{}
		affectedFlag := 7
		for _, flag := range flags {
			if flag != '-' {
				affectedFlags = append(affectedFlags, affectedFlag)
			}
			affectedFlag--
		}

		operations = append(operations, CPUOperation{
			Mnemonic:  mnemo,
			Category:  type_,
			Flags:     affectedFlags,
			Operation: operation,
		})
	}
	return operations
}

func parseAddressingModes(section []string) []CPUAddressingMode {
	addressingModes := []CPUAddressingMode{}
	for _, line := range section {
		split := strings.SplitN(line, fieldSeparator, 4)
		novelSyntax := strings.TrimSpace(split[0])
		syntax := strings.TrimSpace(split[2])
		description := strings.TrimSpace(split[3])
		size, err := strconv.Atoi(strings.TrimSpace(split[1]))
		if err != nil {
			panic(err)
		}

		addressingModes = append(addressingModes, CPUAddressingMode{
			SyntaxNovel: novelSyntax,
			Size:        size,
			Syntax:      syntax,
			Description: description,
		})
	}
	return addressingModes
}

func parseOpcodes(section []string) []CPUOpcode {
	opcodes := []CPUOpcode{}
	for _, line := range section {
		split := strings.Split(line, fieldSeparator)
		op, err := hex.DecodeString(strings.TrimSpace(split[0]))
		if err != nil {
			panic(err)
		}

		illegal := false
		mnemonic := strings.TrimSpace(split[1])
		if strings.HasPrefix(mnemonic, "*") {
			illegal = true
			mnemonic = strings.TrimPrefix(mnemonic, "*")
		}

		addressingMode := "-" // Implied
		if len(split) > 2 {
			addressingMode = strings.TrimSpace(split[2])
		}

		opcodes = append(opcodes, CPUOpcode{
			Op:             op[0],
			Illegal:        illegal,
			Mnemonic:       mnemonic,
			AddressingMode: addressingMode,
		})
	}
	return opcodes
}

func parseVectors(section []string) []CPUVector {
	vectors := []CPUVector{}

	vectorMap := parseSectionToMap(section)
	for addressString, name := range vectorMap {
		addressBytes, err := hex.DecodeString(addressString)
		if err != nil {
			panic(err)
		}

		address := binary.BigEndian.Uint16(addressBytes)

		vectors = append(vectors, CPUVector{
			Address: address,
			Name:    name,
		})
	}

	return vectors
}
