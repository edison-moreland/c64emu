package memory

import (
	"context"
	"fmt"
)

const (
	bank3Start uint16 = 0x8000
	bank3End   uint16 = 0x9FFF

	bank4Start uint16 = 0xA000
	bank4End   uint16 = 0xBFFF

	bank6Start uint16 = 0xD000
	bank6End   uint16 = 0xDFFF

	bank7Start uint16 = 0xE000
	bank7End   uint16 = 0xFFFF
)

type Bank int

const (
	Bank_None Bank = iota - 1
	// Banks 1 and 2 don't have any slots, other than ram.
	Bank_3 // 1 slot, cartridge low
	Bank_4 // 2 slots, basic & cartridge high
	// Same with bank 5
	Bank_6 // 2 slots, character & io
	Bank_7 // 2 slots, kernel & cartridge high
)

func (b Bank) String() string {
	switch b {
	case Bank_3:
		return "3"
	case Bank_4:
		return "4"
	case Bank_6:
		return "6"
	case Bank_7:
		return "7"
	default:
		return "X"
	}
}

type Slot int

const (
	Slot_RAM Slot = iota - 1
	Slot_1
	Slot_2
)

func (s Slot) String() string {
	switch s {
	case Slot_RAM:
		return "RAM"
	case Slot_1:
		return "001"
	case Slot_2:
		return "002"
	default:
		return "XXX"
	}
}

type Banks struct {
	_banks [4]struct {
		Slots  [2]MemoryDevice
		Mapped Slot
	}

	debug bool
}

func newBanks(debug bool) Banks {
	return Banks{
		debug: debug,
	}
}

func (b *Banks) AddDevice(bank Bank, slot Slot, device MemoryDevice) {
	if b.debug {
		fmt.Printf("banks --- AddDevice bank=%s slot=%s device=%T\n", bank, slot, device)
	}

	if bank == Bank_3 && slot == Slot_2 {
		panic("Bank 3 Slot 2 cannot be filled")
	}

	b._banks[bank].Slots[slot] = device
}

func (b *Banks) Switch(bank Bank, slot Slot) {
	if b.debug {
		fmt.Printf("banks --- Switch bank=%s, slot=%s\n", bank, slot)
	}

	if bank == Bank_3 && slot == Slot_2 {
		panic("Bank 3 Slot 2 cannot be mapped")
	}

	b._banks[bank].Mapped = slot

	if b.debug {
		// Sanity check
		if slot != Slot_RAM {
			if b._banks[bank].Slots[slot] == nil {
				panic(fmt.Sprintf("No device mapped to bank %s slot %s", bank, slot))
			}
		}
	}
}

func selectBank(address uint16) (Bank, uint16) {
	switch {
	case address >= bank3Start && address <= bank3End:
		return Bank_3, bank3Start
	case address >= bank4Start && address <= bank4End:
		return Bank_4, bank4Start
	case address >= bank6Start && address <= bank6End:
		return Bank_6, bank6Start
	case address >= bank7Start && address <= bank7End:
		return Bank_7, bank7Start
	default:
		return Bank_None, 0
	}
}

func (b *Banks) redirectRequest(req Request) (redirected bool) {
	bank, startAddress := selectBank(req.Address)
	if bank == Bank_None {
		return false
	}

	slot := b._banks[bank].Mapped
	if slot == Slot_RAM {
		return false
	}

	req.Address -= startAddress
	b._banks[bank].Slots[slot].Request() <- req

	return false
}

func (b *Banks) startDevices(ctx context.Context) {
	for _, bank := range b._banks {
		for _, slot := range bank.Slots {
			if slot != nil {
				slot.Start(ctx)
			}
		}
	}
}
