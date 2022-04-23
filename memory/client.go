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
		Address:  a,
		Response: response,
	}

	out := <-response

	if c.debugMode {
		fmt.Printf("memory --- readByte $%04X: $%02X \n", a, out)
	}

	return out
}

func (c *Client) ReadWord(a uint16) uint16 {
	response := make(ResponseChannel)
	defer close(response)
	var responseBytes [2]byte

	c.request <- Request{
		Type:     RequestType_Read,
		Address:  a,
		Response: response,
	}
	responseBytes[0] = <-response

	c.request <- Request{
		Type:     RequestType_Read,
		Address:  a + 1,
		Response: response,
	}
	responseBytes[1] = <-response

	word := binary.LittleEndian.Uint16(responseBytes[:])

	if c.debugMode {
		fmt.Printf("memory --- readWord $%04X: $%04X \n", a, word)
	}

	return word
}

func (c *Client) Read(address, size uint16) []byte {
	response := make(ResponseChannel)
	defer close(response)
	responseBytes := make([]byte, size)

	for i := uint16(0); i < size; i++ {
		c.request <- Request{
			Type:     RequestType_Read,
			Address:  address + i,
			Response: response,
		}
		responseBytes[i] = <-response
	}

	if c.debugMode {
		fmt.Printf("memory --- read $%04X: %X \n", address, responseBytes)
	}

	return responseBytes
}

func (c *Client) WriteByte(a uint16, b byte) {
	if c.debugMode {
		fmt.Printf("memory --- writeByte $%04X: $%02X \n", a, b)
	}

	c.request <- Request{
		Type:    RequestType_Write,
		Address: a,
		Data:    b,
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
		Address: a,
		Data:    wordBytes[0],
	}

	c.request <- Request{
		Type:    RequestType_Write,
		Address: a + 1,
		Data:    wordBytes[1],
	}
}

func (c *Client) Write(address uint16, data []byte) {
	if c.debugMode {
		fmt.Printf("memory --- write $%04X: %X \n", address, data)
	}

	for i := uint16(0); i < uint16(len(data)); i++ {
		c.request <- Request{
			Type:    RequestType_Write,
			Address: address + i,
			Data:    data[i],
		}
	}
}
