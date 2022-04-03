package main

import (
	"context"
	"embed"

	"github.com/edison-moreland/c64emu/cpu"
	"github.com/edison-moreland/c64emu/memory"
)

//go:embed roms/*.bin
var roms embed.FS

const (
	KernalRom    = "roms/kernal.901227-03.bin"
	BasicRom     = "roms/basic.901226-01.bin"
	CharacterRom = "roms/characters.901225-01.bin"
)

func main() {
	mem := memory.New()

	kernalRom, err := memory.NewRomFromFile(roms, KernalRom)
	if err != nil {
		panic(err)
	}
	mem.AddDevice(0xE000, 0xFFFF, kernalRom)

	basicRom, err := memory.NewRomFromFile(roms, BasicRom)
	if err != nil {
		panic(err)
	}
	mem.AddDevice(0xA000, 0xBFFF, basicRom)

	characterRom, err := memory.NewRomFromFile(roms, CharacterRom)
	if err != nil {
		panic(err)
	}
	mem.AddDevice(0xD000, 0xDFFF, characterRom)

	memoryContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	mem.Start(memoryContext)

	cpu.New(*mem.Client(), true).Start()
}
