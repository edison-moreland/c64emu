package memory

import "context"

// type memoryRedirection struct {
// 	// Some devices may want to restrict request types
// 	MemoryDevice
// 	Start, End uint16
// 	Request    RequestChannel
// }

const ramSize = 0xFFFF + 1

type MMU struct {
	Banks

	//redirections []memoryRedirection
	request RequestChannel
	ram     [ramSize]byte
}

func New() *MMU {
	// Accuracy note: C64 ram is not all zeroed on startup.
	// https://csdb.dk/forums/?roomid=11&topicid=116800&firstpost=2
	return &MMU{
		Banks:   newBanks(),
		request: make(RequestChannel),
		ram:     [ramSize]byte{},
	}
}

// func (m *MMU) AddDevice(startAddress, endAddress uint16, device MemoryDevice) {
// 	m.redirections = append(m.redirections, memoryRedirection{
// 		MemoryDevice: device,
// 		Start:        startAddress,
// 		End:          endAddress,
// 		Request:      device.Request(),
// 	})
// }

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

// func (m *MMU) redirectRequest(request Request) (redirected bool) {
// 	for _, redirection := range m.redirections {
// 		if request.Address >= redirection.Start && request.Address <= redirection.End {
// 			request.Address -= redirection.Start
// 			redirection.Request <- request
// 			return true
// 		}
// 	}

// 	return false
// }

func (m *MMU) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-m.request:
			if m.redirectRequest(request) {
				continue
			}

			switch request.Type {
			case RequestType_Read:
				request.Response <- m.ram[request.Address : request.Address+request.Size]
			case RequestType_Write:
				copy(m.ram[request.Address:request.Address+request.Size], request.Data)
			}
		}
	}
}
