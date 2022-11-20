package memory

import (
	"context"
)

type RequestType int

const (
	RequestType_Read RequestType = iota
	RequestType_Write
)

func (t RequestType) String() string {
	switch t {
	case RequestType_Read:
		return "Read"
	case RequestType_Write:
		return "Write"
	default:
		return "Unknown"
	}
}

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

type Sinkhole struct {
	request RequestChannel
}

func NewSinkhole() *Sinkhole {
	return &Sinkhole{
		request: make(RequestChannel),
	}
}

func (r *Sinkhole) Request() RequestChannel {
	return r.request
}

func (r *Sinkhole) Start(ctx context.Context) {
	go r.start(ctx)
}

func (r *Sinkhole) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-r.request:
			switch request.Type {
			case RequestType_Read:
				request.Response <- 0xff
			}
		}
	}
}
