package memory

import "context"

type StopChannel chan struct{}

type MemoryDevice interface {
	Request() RequestChannel
	Start(context.Context)
}
