package memory

type RequestType int

const (
	RequestType_ReadByte RequestType = iota
	RequestType_ReadWord
	RequestType_WriteByte
	RequestType_WriteWord
)

type Request struct {
	Type     RequestType
	Address  uint16
	Data     [2]byte
	Response ResponseChannel
}

type RequestChannel chan Request
type ResponseChannel chan [2]byte
