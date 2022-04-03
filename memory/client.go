package memory

import "encoding/binary"

type Client struct {
	request RequestChannel
}

func (c *Client) ReadByte(a uint16) byte {
	response := make(ResponseChannel)
	defer close(response)

	c.request <- Request{
		Type:     RequestType_ReadByte,
		Address:  a,
		Response: response,
	}

	return (<-response)[0]
}

func (c *Client) ReadWord(a uint16) uint16 {
	response := make(ResponseChannel)
	defer close(response)

	c.request <- Request{
		Type:     RequestType_ReadWord,
		Address:  a,
		Response: response,
	}

	wordBytes := <-response

	word := binary.LittleEndian.Uint16(wordBytes[:])

	return word
}

func (c *Client) WriteByte(a uint16, b byte) {
	c.request <- Request{
		Type:    RequestType_WriteByte,
		Address: a,
		Data:    [2]byte{b},
	}
}

func (c *Client) WriteWord(a uint16, w uint16) {
	wordBytes := [2]byte{}
	binary.LittleEndian.PutUint16(wordBytes[:], w)

	c.request <- Request{
		Type:    RequestType_WriteWord,
		Address: a,
		Data:    wordBytes,
	}
}
