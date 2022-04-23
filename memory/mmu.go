package memory

import "context"

const ramSize = 0xFFFF + 1

type MMU struct {
	Banks

	request RequestChannel
	ram     [ramSize]byte
}

func New(debugBanks bool) *MMU {
	// Accuracy note: C64 ram is not all zeroed on startup.
	// https://csdb.dk/forums/?roomid=11&topicid=116800&firstpost=2
	return &MMU{
		Banks:   newBanks(debugBanks),
		request: make(RequestChannel),
		ram:     [ramSize]byte{},
	}
}

func (m *MMU) Client(debugMode bool) *Client {
	return &Client{
		request:   m.request,
		debugMode: debugMode,
	}
}

func (m *MMU) Start(ctx context.Context) {
	m.startDevices(ctx)
	go m.start(ctx)
}

func (m *MMU) Request() RequestChannel {
	return m.request
}

func (m *MMU) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-m.request:
			// Attempt to redirect to the appropriate bank
			if m.Banks.redirectRequest(request) {
				continue
			}

			// Fallthrough to RAM if it wasnt redirected
			switch request.Type {
			case RequestType_Read:
				request.Response <- m.ram[request.Address]
			case RequestType_Write:
				m.ram[request.Address] = request.Data
			}

			// If cpu port changes, some banks probably need to be swapped
			if request.Address <= 0x0001 {
				m.updateBanks()
			}
		}
	}
}

const (
	CONTROL_LINE_LORAM  = 1 << 0
	CONTROL_LINE_HIRAM  = 1 << 1
	CONTROL_LINE_CHAREN = 1 << 2
)

func (m *MMU) updateBanks() {
	// control lines (0x0001), control direction (0x0000, 1=out, 0=in)
	direction := m.ram[0x0000]
	control := m.ram[0x0001]

	// These control lines control whether these banks are mapped
	// LORAM: basic rom
	// HIRAM: kernal rom
	// CHAREN: character rom

	// Accuracy note:
	//   The actual mapping is a little more complicated than this.
	//   When cartridges are added, this will have to be revisited.
	// More info: https://www.c64-wiki.com/wiki/Bank_Switching

	if direction&CONTROL_LINE_LORAM != 0 {
		// Control is set to output
		if control&CONTROL_LINE_LORAM != 0 {
			// Control line is high
			m.Banks.Switch(Bank_4, Slot_1)
		} else {
			// Control line is low
			m.Banks.Switch(Bank_4, Slot_RAM)
		}
	}

	if direction&CONTROL_LINE_HIRAM != 0 {
		if control&CONTROL_LINE_HIRAM != 0 {
			m.Banks.Switch(Bank_7, Slot_1)
		} else {
			m.Banks.Switch(Bank_7, Slot_RAM)
		}
	}

	if direction&CONTROL_LINE_CHAREN != 0 {
		if control&CONTROL_LINE_CHAREN != 0 {
			m.Banks.Switch(Bank_6, Slot_2)
		} else {
			m.Banks.Switch(Bank_6, Slot_1)
		}
	}

}
