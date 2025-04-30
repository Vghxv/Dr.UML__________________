package utils

type Color struct {
	R, G, B uint8
}

func (c Color) ToHex() int {
	return int(c.R)<<16 | int(c.G)<<8 | int(c.B)
}
