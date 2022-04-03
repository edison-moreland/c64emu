package cpu

func (c *CPU) clearFlag(flag uint8) {
	c.P &^= flag
}

func (c *CPU) setFlag(flag uint8) {
	c.P |= flag
}

// func (c *CPU) toggleFlag(flag uint8) {
// 	c.P ^= flag
// }

func (c *CPU) isFlagSet(flag uint8) bool {
	return c.P&flag != 0
}

func (c *CPU) flag(flag uint8) uint8 {
	if c.P&flag != 0 {
		return 1
	}
	return 0
}
