package main

import (
	"context"
	"embed"

	"github.com/edison-moreland/c64emu/cpu"
	"github.com/edison-moreland/c64emu/memory"
)

//go:embed roms/*.bin
var roms embed.FS

var (
	Roms = []struct {
		Path       string
		Start, End uint16
	}{
		{"roms/kernal.901227-03.bin", 0xE000, 0xFFFF},
		{"roms/basic.901226-01.bin", 0xA000, 0xBFFF},
		{"roms/characters.901225-01.bin", 0xD000, 0xDFFF},
	}
)

func main() {
	mem := memory.New()

	for _, romInfo := range Roms {
		rom, err := memory.NewRomFromFile(roms, romInfo.Path)
		if err != nil {
			panic(err)
		}
		mem.AddDevice(romInfo.Start, romInfo.End, rom)
	}

	memoryContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	mem.Start(memoryContext)

	debug := true
	cpu.New(*mem.Client(false), debug).Start()
}
