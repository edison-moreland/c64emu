package main

import (
	"context"
	"embed"

	"github.com/edison-moreland/c64emu/cpu"
	"github.com/edison-moreland/c64emu/memory"
)

const (
	debugCPU    = true
	debugMemory = false
	debugBanks  = true
	debugStack  = false
)

//go:embed roms/*.bin
var roms embed.FS

var (
	Roms = []struct {
		Path string
		Bank memory.Bank
		Slot memory.Slot
	}{
		{"roms/kernal.901227-03.bin", memory.Bank_7, memory.Slot_1},
		{"roms/basic.901226-01.bin", memory.Bank_4, memory.Slot_1},
		{"roms/characters.901225-01.bin", memory.Bank_6, memory.Slot_1},
	}
)

func main() {
	mem := memory.New(debugBanks)

	for _, romInfo := range Roms {
		rom, err := memory.NewRomFromFile(roms, romInfo.Path)
		if err != nil {
			panic(err)
		}
		mem.AddDevice(romInfo.Bank, romInfo.Slot, rom)
		mem.Switch(romInfo.Bank, romInfo.Slot)
	}

	memoryContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	mem.Start(memoryContext)

	cpu.New(*mem.Client(debugMemory), debugCPU, debugStack).Start()
}
