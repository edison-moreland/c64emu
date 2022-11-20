package main

import (
	"context"
	"embed"
	"fmt"
	"time"

	"github.com/edison-moreland/c64emu/clock"
	"github.com/edison-moreland/c64emu/cpu"
	"github.com/edison-moreland/c64emu/memory"
	"github.com/edison-moreland/c64emu/trace"
)

const (
	debugCPU    = true
	debugMemory = false
	debugBanks  = false
	debugStack  = true
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
	mem.AddDevice(memory.Bank_6, memory.Slot_2, memory.NewSinkhole())

	memoryContext, cancel := context.WithCancel(context.Background())
	defer cancel()

	mem.Start(memoryContext)
	c := clock.New(time.Second / 4)

	t := trace.NewTracer()
	defer t.Close()

	go func() {
		for trace := range t.Out() {
			fmt.Println(trace.String())
		}
	}()

	emu := cpu.New(*mem.Client(debugMemory), t.System("cpu"), debugCPU, debugStack)

	_, doneChan := emu.Start(c.C())

	c.Start()
	defer c.Stop()

	<-doneChan
}
