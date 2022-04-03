package memory

type RequestType int

const (
	RequestType_Read RequestType = iota
	RequestType_Write
)

type Request struct {
	Type     RequestType
	Address  uint16
	Data     []byte
	Size     uint16
	Response ResponseChannel
}

type RequestChannel chan Request
type ResponseChannel chan []byte
