package memory

import (
	"context"
	"io/fs"
	"io/ioutil"
)

type Rom struct {
	rom     []byte
	request RequestChannel
}

func NewRomFromFile(f fs.FS, path string) (*Rom, error) {
	romFile, err := f.Open(path)
	if err != nil {
		return nil, err
	}
	defer romFile.Close()

	romBytes, err := ioutil.ReadAll(romFile)
	if err != nil {
		return nil, err
	}

	return &Rom{
		rom:     romBytes,
		request: make(RequestChannel),
	}, nil
}

func (r *Rom) Request() RequestChannel {
	return r.request
}

func (r *Rom) Start(ctx context.Context) {
	go r.start(ctx)
}

func (r *Rom) start(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case request := <-r.request:
			switch request.Type {
			case RequestType_Read:
				request.Response <- r.rom[request.Address : request.Address+request.Size]
			}
		}
	}
}
