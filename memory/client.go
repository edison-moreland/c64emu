package memory

import (
	"encoding/binary"
	"fmt"
)

type Client struct {
	request   RequestChannel
	debugMode bool
}

func (c *Client) ReadByte(a uint16) byte {
	response := make(ResponseChannel)
	defer close(response)

	c.request <- Request{
		Type:     RequestType_Read,
		Size:     1,
		Address:  a,
		Response: response,
	}

	out := (<-response)[0]

	if c.debugMode {
		fmt.Printf("memory --- readByte $%04X: $%02X \n", a, out)
	}

	return out
}

func (c *Client) ReadWord(a uint16) uint16 {
	response := make(ResponseChannel)
	defer close(response)

	c.request <- Request{
		Type:     RequestType_Read,
		Size:     2,
		Address:  a,
		Response: response,
	}

	wordBytes := <-response

	word := binary.LittleEndian.Uint16(wordBytes[:])

	if c.debugMode {
		fmt.Printf("memory --- readWord $%04X: $%04X \n", a, word)
	}

	return word
}

func (c *Client) Read(address, size uint16) []byte {
	response := make(ResponseChannel)
	defer close(response)

	c.request <- Request{
		Type:     RequestType_Read,
		Size:     size,
		Address:  address,
		Response: response,
	}

	out := <-response

	if c.debugMode {
		fmt.Printf("memory --- read $%04X: %X \n", address, out)
	}

	return out
}

func (c *Client) WriteByte(a uint16, b byte) {
	if c.debugMode {
		fmt.Printf("memory --- writeByte $%04X: $%02X \n", a, b)
	}

	c.request <- Request{
		Type:    RequestType_Write,
		Size:    1,
		Address: a,
		Data:    []byte{b},
	}
}

func (c *Client) WriteWord(a uint16, w uint16) {
	if c.debugMode {
		fmt.Printf("memory --- writeWord $%04X: $%04X \n", a, w)
	}

	wordBytes := make([]byte, 2)
	binary.LittleEndian.PutUint16(wordBytes, w)

	c.request <- Request{
		Type:    RequestType_Write,
		Size:    2,
		Address: a,
		Data:    wordBytes,
	}
}

func (c *Client) Write(address uint16, data []byte) {
	if c.debugMode {
		fmt.Printf("memory --- write $%04X: %X \n", address, data)
	}

	c.request <- Request{
		Type:    RequestType_Write,
		Size:    uint16(len(data)),
		Address: address,
		Data:    data,
	}
}
