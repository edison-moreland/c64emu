package memory

import "context"

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
	Bank_None Bank = iota
	// Banks 1 and 2 don't have any slots, other than ram.
	Bank_3 // 1 slot, cartridge low
	Bank_4 // 2 slots, basic & cartridge high
	// Same with bank 5
	Bank_6 // 2 slots, character & io
	Bank_7 // 2 slots, kernel & cartridge high
)

type Slot int

const (
	Slot_RAM Slot = iota - 1
	Slot_1
	Slot_2
)

type Banks struct {
	// This gets embeded in MMU
	bank3Slots      [1]MemoryDevice
	bank3MappedSlot Slot
	bank4Slots      [2]MemoryDevice
	bank4MappedSlot Slot
	bank6Slots      [2]MemoryDevice
	bank6MappedSlot Slot
	bank7Slots      [2]MemoryDevice
	bank7MappedSlot Slot
}

func newBanks() Banks {
	return Banks{
		bank3Slots:      [1]MemoryDevice{},
		bank4Slots:      [2]MemoryDevice{},
		bank6Slots:      [2]MemoryDevice{},
		bank7Slots:      [2]MemoryDevice{},
		bank3MappedSlot: Slot_RAM,
		bank4MappedSlot: Slot_RAM,
		bank6MappedSlot: Slot_RAM,
		bank7MappedSlot: Slot_RAM,
	}
}

func (b *Banks) AddDevice(bank Bank, slot Slot, device MemoryDevice) {
	switch bank {
	case Bank_3:
		if slot == Slot_2 {
			panic("Bank 3 Slot 2 cannot be filled")
		}
		b.bank3Slots[slot] = device
	case Bank_4:
		b.bank4Slots[slot] = device
	case Bank_6:
		b.bank6Slots[slot] = device
	case Bank_7:
		b.bank7Slots[slot] = device
	}
}

func (b *Banks) Switch(bank Bank, slot Slot) {
	switch bank {
	case Bank_3:
		if slot == Slot_2 {
			panic("Bank 3 Slot 2 cannot be mapped")
		}
		b.bank3MappedSlot = slot
	case Bank_4:
		b.bank4MappedSlot = slot
	case Bank_6:
		b.bank6MappedSlot = slot
	case Bank_7:
		b.bank7MappedSlot = slot
	}
}

func (b *Banks) redirectRequest(req Request) (redirected bool) {
	switch {
	case req.Address >= bank3Start && req.Address <= bank3End:
		if b.bank3MappedSlot == Slot_RAM {
			return false
		}
		req.Address -= bank3Start
		b.bank3Slots[b.bank3MappedSlot].Request() <- req
		return true
	case req.Address >= bank4Start && req.Address <= bank4End:
		if b.bank4MappedSlot == Slot_RAM {
			return false
		}
		req.Address -= bank4Start
		b.bank4Slots[b.bank4MappedSlot].Request() <- req
		return true
	case req.Address >= bank6Start && req.Address <= bank6End:
		if b.bank6MappedSlot == Slot_RAM {
			return false
		}
		req.Address -= bank6Start
		b.bank6Slots[b.bank6MappedSlot].Request() <- req
		return true
	case req.Address >= bank7Start && req.Address <= bank7End:
		if b.bank7MappedSlot == Slot_RAM {
			return false
		}
		req.Address -= bank7Start
		b.bank7Slots[b.bank7MappedSlot].Request() <- req
		return true
	}

	return false
}

func (b *Banks) startDevices(ctx context.Context) {
	for _, slot := range b.bank3Slots {
		if slot != nil {
			slot.Start(ctx)
		}
	}
	for _, slot := range b.bank4Slots {
		if slot != nil {
			slot.Start(ctx)
		}
	}
	for _, slot := range b.bank6Slots {
		if slot != nil {
			slot.Start(ctx)
		}
	}
	for _, slot := range b.bank7Slots {
		if slot != nil {
			slot.Start(ctx)
		}
	}
}
