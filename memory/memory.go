package memory

import "context"

type Address uint16

type memoryRedirection struct {
	// Some devices may want to restrict request types
	MemoryDevice
	Start, End Address
	Request    RequestChannel
}

const ramSize = 0xFFFF + 1

// Note: This does not handle bank switching yet, but i've got ideas
type Memory struct {
	redirections []memoryRedirection
	request      RequestChannel
	ram          [ramSize]byte
}

func New() *Memory {
	// Accuracy note: C64 ram is not all zeroed on startup.
	// https://csdb.dk/forums/?roomid=11&topicid=116800&firstpost=2
	return &Memory{
		request: make(RequestChannel),
		ram:     [ramSize]byte{},
	}
}

func (m *Memory) AddDevice(startAddress, endAddress Address, device MemoryDevice) {
	m.redirections = append(m.redirections, memoryRedirection{
		MemoryDevice: device,
		Start:        startAddress,
		End:          endAddress,
		Request:      device.Request(),
	})
}

func (m *Memory) Client() *Client {
	return &Client{
		request: m.request,
	}
}

func (m *Memory) Start(ctx context.Context) error {
	for _, redirection := range m.redirections {
		redirection.MemoryDevice.Start(ctx)
	}
	go m.start(ctx)
	return nil
}

func (m *Memory) Request() RequestChannel {
	return m.request
}

func (m *Memory) redirectRequest(request Request) (redirected bool) {
	for _, redirection := range m.redirections {
		if request.Address >= redirection.Start && request.Address <= redirection.End {
			redirection.Request <- request
			return true
		}
	}

	return false
}

func (m *Memory) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-m.request:
			if m.redirectRequest(request) {
				continue
			}

			switch request.Type {
			case RequestType_ReadByte:
				request.Response <- [2]byte{m.ram[request.Address], 0}
			case RequestType_ReadWord:
				request.Response <- [2]byte{m.ram[request.Address], m.ram[request.Address+1]}
			case RequestType_WriteByte:
				m.ram[request.Address] = request.Data[0]
			case RequestType_WriteWord:
				m.ram[request.Address] = request.Data[0]
				m.ram[request.Address+1] = request.Data[1]
			}
		}
	}
}
