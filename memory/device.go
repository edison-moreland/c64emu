package memory

import "context"

type RequestType int

const (
	RequestType_Read RequestType = iota
	RequestType_Write
)

type Request struct {
	Type     RequestType
	Address  uint16
	Data     byte
	Response ResponseChannel
}

type RequestChannel chan Request
type ResponseChannel chan byte

type MemoryDevice interface {
	Request() RequestChannel
	Start(context.Context)
}
