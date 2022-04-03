package memory

import "encoding/binary"

type Client struct {
	request RequestChannel
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

	return (<-response)[0]
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

	return <-response
}

func (c *Client) WriteByte(a uint16, b byte) {
	c.request <- Request{
		Type:    RequestType_Write,
		Size:    1,
		Address: a,
		Data:    []byte{b},
	}
}

func (c *Client) WriteWord(a uint16, w uint16) {
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
	c.request <- Request{
		Type:    RequestType_Write,
		Size:    uint16(len(data)),
		Address: address,
		Data:    data,
	}
}
